package main

import (
	"demo/grpc_test/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync/atomic"
	"time"
)

const (
	//address = "localhost:50051"
	address = "192.168.1.201:50051"
)

func main() {
	conn, er := grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
	}
	defer func() {
		er = conn.Close()
	}()

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	var count int64

	go func() {
		ctx := context.Background()
		greet := helloworld.NewGreeterClient(conn)
		for i := 0; i < 10000; i++ {
			begin := time.Now()
			var reply *helloworld.HelloReply
			reply, er = greet.SayHello(ctx, &helloworld.HelloRequest{Name: "world"})
			if er != nil {
				log.Printf("did not connect:%v", er)
				continue
			}
			atomic.AddInt64(&count, 1)
			log.Printf("[%d] [%v] Greeting: %s", count, time.Now().Sub(begin), reply.Message)
		}
	}()

	go func() {
		ctx := context.Background()
		math := helloworld.NewMathServiceClient(conn)
		for i := 0; i < 10000; i++ {
			begin := time.Now()
			var reply *helloworld.AddReply
			reply, er = math.Add(ctx, &helloworld.AddRequest{A: 10, B: 10})
			if er != nil {
				log.Printf("did not connect:%v", er)
				continue
			}
			atomic.AddInt64(&count, 1)
			log.Printf("[%d] [%v] Greeting: %d", count, time.Now().Sub(begin), reply.C)
		}
	}()

	select {}
}
