package service

import (
	"demo/grpc_test/proto/helloworld"
	"golang.org/x/net/context"
	"log"
)

type MathServer struct {
}

func (s *MathServer) Add(ctx context.Context, req *helloworld.AddRequest) (*helloworld.AddReply, error) {
	log.Println(req)

	rp := &helloworld.AddReply{
		C: req.A + req.B,
	}

	return rp, nil
}

func (s *MathServer) Sub(ctx context.Context, req *helloworld.SubRequest) (*helloworld.SubReply, error) {
	log.Println(req)

	rp := &helloworld.SubReply{
		C: req.A - req.B,
	}

	return rp, nil
}
