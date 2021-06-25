package main

import (
	"context"
	"fmt"
	pb "grpctest/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "konnitiha " + in.Name}, nil
}

func main() {
	fmt.Println("server now listening")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
