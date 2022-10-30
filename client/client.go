package main

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/kkumar30/grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"time"
)

const address = "localhost:5001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Cant hear you %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var msg = " lavanlavanlavanlavan"
	resp, err := c.SayHello(ctx, &pb.Message{Body: msg})
	fmt.Printf("Response %s\n", resp.GetBody())

	var query = "kush"
	response, err := c.QueryLogFiles(ctx, &pb.QueryInput{Query: query})
	fmt.Printf("Response lines %s\n Count: %d", response.LogLines, response.Count)
	fmt.Println("*******")

	// Upload file ****************************************************

	file, err := os.Open("../sample.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stream, err := c.UploadFile(ctx)
	if err != nil {
		log.Fatalf("Error while calling UploadFile RPC: %v", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading file: %v", err)
		}
		request := &pb.UploadFileRequest{
			Filename: "sample.log",
			File:     buffer[:n],
		}
		err = stream.Send(request)
		if err != nil {
			log.Fatalf("Error while sending data to server: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err == io.EOF {
		log.Printf("Reached EOF. Data sent successfully")
	}

	log.Printf("UploadFile %s Response: %d", res.GetFilename(), res.GetStatus())

}
