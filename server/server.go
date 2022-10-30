package main

import (
	"bytes"
	"context"
	"fmt"
	pb "github.com/kkumar30/grpc-test/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
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

func (s *ChatServer) UploadFile(stream pb.ChatService_UploadFileServer) error {

	fileData := bytes.Buffer{}
	var fileName string
	log.Printf("Starting to download from client: \n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("no more data from client")
			break
		}
		if err != nil {
			log.Printf("Error while reading client stream: %v", err)
		}

		chunk := req.GetFile()
		fileName = req.GetFilename()

		_, err = fileData.Write(chunk)

		err = os.WriteFile(fileName, fileData.Bytes(), 0600)
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		fmt.Printf("File %s written", fileName)
	}
	response := &pb.UploadFileResponse{
		Filename: fileName,
		Status:   "Success",
	}

	err := stream.SendAndClose(response)
	if err != nil {
		log.Fatalf("Error while sending response to client: %v", err)
	}

	return nil
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
