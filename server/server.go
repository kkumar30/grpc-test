package main

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/kkumar30/grpc-test/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
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

func (s *ChatServer) QueryLogFiles(ctx context.Context, input *pb.QueryInput) (*pb.QueryResults, error) {
	log.Printf("Searching query : %s\n", input.GetQuery())
	lines, count := scanLogs(input.GetQuery())
	return &pb.QueryResults{LogLines: lines, Count: count}, nil
}

func scanLogs(query string) ([]string, int32) {
	file, err := os.Open("sample.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int32
	var results []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // GET the line string
		if strings.Contains(line, query) {
			fmt.Println(line)
			count += 1
			results = append(results, line)
		}
	}
	return results, count
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port 5001: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &ChatServer{})
	fmt.Printf("Server listening at %v\n", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
