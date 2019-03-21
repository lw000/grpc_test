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
	address = "localhost:50051"
	//address = "192.168.1.201:50051"
)

var (
	count int64
	conn  *grpc.ClientConn
)

func TestGreeter() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	ctx := context.Background()
	greet := helloworld.NewGreeterClient(conn)
	for i := 0; i < 100; i++ {
		begin := time.Now()
		reply, er := greet.SayHello(ctx, &helloworld.HelloRequest{Name: "world"})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		atomic.AddInt64(&count, 1)
		log.Printf("[%d] [%v] Greeting: %s", count, time.Now().Sub(begin), reply.Message)
	}
}

func TestMathService() {
	ctx := context.Background()
	math := helloworld.NewMathServiceClient(conn)

	{
		reply, er := math.Add(ctx, &helloworld.AddRequest{A: 10, B: 30})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		atomic.AddInt64(&count, 1)
		log.Printf("Add: [%d] [%v]", count, reply.C)
	}

	{
		reply, er := math.Sub(ctx, &helloworld.SubRequest{A: 10, B: 100})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		atomic.AddInt64(&count, 1)
		log.Printf("Sub: [%d] [%v]", count, reply.C)
	}
}

func main() {
	var er error
	conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
	}

	defer func() {
		er = conn.Close()
	}()

	go TestGreeter()

	go TestMathService()

	select {}
}
