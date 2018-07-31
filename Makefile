all: client server

proto:
	protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/chat.proto

cert: cert.pem

cert.pem:
	openssl req -newkey rsa:2048 -nodes -keyout cert.pem -x509 -days 365 -out key.pem -subj "/O=Personal/CN=localhost"

server:
	go build ./cmd/server

client:
	go build ./cmd/client

.PHONY: client server proto
