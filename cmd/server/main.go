package main

import (
	"demo/grpc_test/cmd/server/service"
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
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}

	serv := grpc.NewServer()

	helloworld.RegisterGreeterServer(serv, &service.GreeterServer{})
	helloworld.RegisterMathServiceServer(serv, &service.MathServer{})

	reflection.Register(serv)

	if err = serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
