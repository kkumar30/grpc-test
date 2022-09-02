package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"

	pb "github.com/kkumar30/grpc-test/proto"
)

func main() {
	fmt.Println("Here")
	math.Abs(-6)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen to port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000 %v", err)
	}

}
