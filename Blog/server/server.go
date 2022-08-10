package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "Blog/proto/pb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr = "127.0.0.1:9080"

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// Create connection with mongoDB database -----------------------------------------------------//
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		fmt.Println("Error while creating DataBase Instance: ", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Error while connecting DataBase : ", err)
	}

	collection = client.Database("blogdb").Collection("blog") // if db or collections not exits, then its create it

	// -------------------------------------------------------------------------------------------- //

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("Failed to listen on:", err)
	}

	fmt.Println("Listening on:", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
