package main

import (
	"demo/grpc_test/cmd/server/srv"
	"demo/grpc_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &srv.GreeterServer{})
	helloworld.RegisterMathServiceServer(s, &srv.MatchServer{})

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
