package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"github.com/iron-s/chat/proto"
)

var (
	addr = flag.String("server address", ":8531", "The server address to listen on")
	cert = flag.String("cert", "", "certificate file to use")
	key  = flag.String("key", "", "key file to use")
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	NewGrpcServer(addr, cert, key)
}

type user struct {
	name        string
	password    string
	connections []chan *proto.Event
}

type grpcSrv struct {
	users map[string]*user
	um    sync.RWMutex
	br    chan *proto.Event
}

//NewGrpcServer starts server listening on addr
func NewGrpcServer(addr *string, cert, key *string) {
	var (
		l   net.Listener
		err error
	)
	opts := []grpc.ServerOption{grpc.KeepaliveParams(
		keepalive.ServerParameters{Time: 10 * time.Second, Timeout: 5 * time.Second})}
	if *cert != "" && *key != "" {
		creds, err := credentials.NewServerTLSFromFile(*cert, *key)
		if err != nil {
			log.Fatalln("error loading keypair:", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	l, err = net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}
	s := grpc.NewServer(opts...)
	srv := &grpcSrv{users: make(map[string]*user), br: make(chan *proto.Event)}
	proto.RegisterChatServer(s, srv)
	go srv.broadcast()
	// TODO add graceful shutdown
	log.Fatalln("serve failed: ", s.Serve(l))
}

// Messages is a handler for the RPC of the same name
func (s *grpcSrv) Messages(stream proto.Chat_MessagesServer) error {
	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Println("receive failed:", err)
		return err
	}

	// TODO might want to use interceptor for authentication
	// handle user join
	join, ok := req.Msg.(*proto.Message_Join)
	if !ok {
		log.Println("protocol violation, expect join message, got", req.Msg)
		return grpc.Errorf(codes.FailedPrecondition, "protocol violation expect join message")
	}

	// TODO consider buffered channel
	c := make(chan *proto.Event)

	s.um.Lock()
	u, err := s.joinUser(join.Join, c)
	joined := err == nil && len(u.connections) == 1
	s.um.Unlock()
	if err != nil {
		return err
	}

	started := make(chan struct{})
	go func() {
		close(started)
		for e := range c {
			err := stream.Send(e)
			if err != nil {
				log.Println("error sending:", err)
				return
			}
		}
	}()
	<-started

	if joined {
		s.br <- &proto.Event{Time: ptypes.TimestampNow(), Who: u.name, Msg: &proto.Event_Join{&proto.Join{}}}
	}

	for {
		req, err := stream.Recv()
		msg := req.GetMsg()
		status := status.Convert(err)
		e := &proto.Event{Time: ptypes.TimestampNow(), Who: u.name}
		switch {
		case status.Code() == codes.OK:
			// do nothing, we'll use msg as is
		case err == io.EOF || status.Code() == codes.Canceled:
			if left := s.leaveUser(u.name, c); left {
				e.Msg = &proto.Event_Leave{&proto.Leave{}}
				s.br <- e
			}
			return nil
		default:
			// unknown error
			log.Println("receive failed:", err)
			return err
		}
		switch m := msg.(type) {
		case *proto.Message_Join:
			e.Msg = &proto.Event_Join{&proto.Join{}}
		case *proto.Message_Text:
			e.Msg = &proto.Event_Text{&proto.Text{Text: trimSpace(m.Text.Text)}}
		}
		s.br <- e
	}
	return nil
}

func (s *grpcSrv) broadcast() {
	for e := range s.br {
		log.Println("broadcasting", e)
		s.um.RLock()
		for _, u := range s.users {
			for _, c := range u.connections {
				select {
				case c <- e:
				default:
					log.Println("channel full for user", u.name)
				}
			}
		}
		s.um.RUnlock()
	}
}

func (s *grpcSrv) joinUser(join *proto.Join, c chan *proto.Event) (*user, error) {
	joinName, joinPass := join.GetName(), join.GetPass()
	if joinName == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "name can't be empty")
	}
	// TODO password comparison here is not cryptographycally secure, could use crypto/subtle for comparison
	if u, ok := s.users[joinName]; ok && joinPass != u.password {
		return nil, grpc.Errorf(codes.Unauthenticated, "permission denied")
	} else if ok {
		u.connections = append(u.connections, c)
	} else {
		joinName = trimSpace(joinName)
		s.users[joinName] = &user{name: joinName, password: joinPass, connections: []chan *proto.Event{c}}
	}
	return s.users[joinName], nil
}

func (s *grpcSrv) leaveUser(user string, c chan *proto.Event) bool {
	var left bool
	s.um.Lock()
	u, ok := s.users[user]
	if ok {
		for i, con := range u.connections {
			if c == con {
				u.connections = append(u.connections[:i], u.connections[i+1:]...)
				close(c)
				break
			}
		}
		if len(u.connections) == 0 {
			delete(s.users, user)
			left = true
		}
	}
	s.um.Unlock()
	return left
}

func trimSpace(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '\n', '\r', '\t':
			return ' '
		default:
			return r
		}
	}, s)
}
