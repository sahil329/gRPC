package main

import (
	"fmt"
	"log"
	"net"

	pb "Calculator/proto/pb"

	"google.golang.org/grpc"
)

var addr = "127.0.0.1:9080"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("Failed to listen on:", err)
	}

	fmt.Println("Listening on:", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
