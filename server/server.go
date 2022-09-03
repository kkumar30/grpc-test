package main

import (
	"context"
	"fmt"
	pb "github.com/kkumar30/grpc-test/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":5001"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServer) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	log.Printf("Received message from client: %s\n", message)
	return &pb.Message{Body: "Hello from the server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &ChatServer{})
	fmt.Printf("Server listening at %v\n", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
