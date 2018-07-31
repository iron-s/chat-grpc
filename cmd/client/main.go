package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"github.com/iron-s/chat/proto"
)

var (
	addr   = flag.String("server address", ":8531", "The server address in the format of host:port")
	name   = flag.String("name", "", "name to join to the chat")
	pass   = flag.String("pass", "", "password to join to the chat")
	secure = flag.Bool("secure", false, "enable TLS")
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	rl, err := readline.NewEx(&readline.Config{Prompt: "> ", UniqueEditLine: true})
	if err != nil {
		log.Fatalln("cannot create readline:", err)
	}
	log.SetOutput(rl.Stderr())
	defer rl.Close()

	opts := []grpc.DialOption{grpc.WithKeepaliveParams(
		keepalive.ClientParameters{Time: 10 * time.Second, Timeout: 5 * time.Second})}
	if *secure {
		opts = append(opts, grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(nil, "")))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalln("connect failed:", err)
	}
	defer conn.Close()

	client := proto.NewChatClient(conn)
	stream, err := client.Messages(context.Background())
	if err != nil {
		log.Fatalln("error opening stream:", err)
	}

	go receive(stream, rl)

	join := &proto.Message{Msg: &proto.Message_Join{&proto.Join{
		Name: *name,
		Pass: *pass,
	}}}
	err = stream.Send(join)
	if err != nil {
		log.Fatalln("error joining: ", err)
	}

	rl.SetPrompt(*name + "> ")
	for {
		text, err := rl.Readline()
		if err != nil { // Ctrl-C or Ctrl-D
			stream.CloseSend()
			break
		}
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		err = stream.Send(&proto.Message{Msg: &proto.Message_Text{&proto.Text{Text: text}}})
		if err != nil {
			log.Fatalln("error sending message:", err)
		}
	}
}

func receive(stream proto.Chat_MessagesClient, rl *readline.Instance) {
	for {
		in, err := stream.Recv()
		status := status.Convert(err)
		code := status.Code()
		switch {
		case err == io.EOF || code == codes.Canceled || code == codes.Unavailable:
			// stream closed.
			return
		case code == codes.InvalidArgument:
			log.Fatalln("error:", status.Message())
		}
		if err != nil {
			log.Fatalln("Failed to receive message:", err)
		}
		switch e := in.Msg.(type) {
		case *proto.Event_Text:
			fmt.Fprintf(rl, "%-11s %s\n", in.Who+":", e.Text.Text)
		case *proto.Event_Join:
			fmt.Fprintf(rl, "*** %s is online\n", in.Who)
		case *proto.Event_Leave:
			fmt.Fprintf(rl, "*** %s is offline\n", in.Who)
		default:
			log.Printf("Got unknown message %T, %v\n", in, in)
		}
	}
}
