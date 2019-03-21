package srv

import (
	"demo/grpc_test/proto"
	"golang.org/x/net/context"
	"log"
)

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Println(req)

	rp := &helloworld.HelloReply{
		Message: "Hello" + req.Name,
	}

	return rp, nil
}
