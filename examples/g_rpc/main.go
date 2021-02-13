package main

import (
	"context"
	"log"
	"net"

	"g_rpc/service"

	"google.golang.org/grpc"
)

type server struct {
	service.UnimplementedGreeterServer
}

func (server) SayHello(context.Context, *service.HelloRequest) (*service.HelloReply, error) {
	res := service.HelloReply{
		Message: "Hello there!",
	}

	return &res, nil
}

func (server) SayAnotherHello(context.Context, *service.HelloRequest) (*service.HelloReply, error) {
	res := service.HelloReply{
		Message: "Hello there another!",
	}

	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
