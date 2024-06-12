package main

import (
	"context"
	"fmt"
	"github.com/peifengll/Interview_collection/ninneng/ans2/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	if !valid(md["authorization"]) {
		return nil, fmt.Errorf("invalid token")
	}

	return handler(ctx, req)
}

func logInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Received request: %v", info.FullMethod)
	resp, err := handler(ctx, req)
	log.Printf("Sent response: %v, error: %v", resp, err)
	return resp, err
}

func valid(authorization []string) bool {
	// 实现你的令牌验证逻辑
	return len(authorization) > 0 && authorization[0] == "valid-token"
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(logInterceptor, authInterceptor),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
