package srv

import (
	"demo/grpc_test/proto"
	"golang.org/x/net/context"
	"log"
)

type MatchServer struct {
}

func (s *MatchServer) Add(ctx context.Context, req *helloworld.AddRequest) (*helloworld.AddReply, error) {
	log.Println(req)

	rp := &helloworld.AddReply{
		C: req.A + req.B,
	}

	return rp, nil
}
