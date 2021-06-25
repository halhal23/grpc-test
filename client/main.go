package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "grpctest/helloworld"

	"google.golang.org/grpc"
)

func main() {
	// connection 作成
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewGreetingClient(conn)
	defer conn.Close()

	// 引数の準備
	name := "hiroharu"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// contextの準備
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// SaiHelloの呼び出し
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not invoke: %v", err)
	}
	fmt.Printf("response: %s", response.Message)
}
