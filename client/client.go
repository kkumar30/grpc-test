package main

import (
	"context"
	"fmt"
	pb "github.com/kkumar30/grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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
	fmt.Printf("Here")
	fmt.Printf("Response %s\n", resp.GetBody())
}
