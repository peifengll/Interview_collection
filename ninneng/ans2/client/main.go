package main

import (
	"context"
	"github.com/peifengll/Interview_collection/ninneng/ans2/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	// 调用 SayHello RPC 方法
	name := "world"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = context.WithValue(ctx, "nihao", 666)
	defer cancel()
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())

}
