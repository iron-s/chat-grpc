# GRPC chat
Prototype chat program written using gRPC bidirectional stream.
gRPC handles all the network failures itself, and we just handle errors from gRPC.

## Build
To build just run
```
make
```
Alternatively you can build each program independently using `go build`:
```
go build ./cmd/client
go build ./cmd/server
```

## Running without TLS
Server:
```
./server
```
Client:
```
./client -name Alice
```
or protected by password
```
./client -name Bob -password secret
```
## Running wiht TLS
To enable TLS you should create certificate and key:
```
make cert
```
or
```
openssl req -newkey rsa:2048 -nodes -keyout cert.pem -x509 -days 365 -out key.pem -subj "/O=Personal/CN=localhost"
```
Then run server with certificate and key:
```
./server -cert cert.pem -key key.pem
```
And client with TLS configured:
```
./client -name Alice -secure
```
