package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogID) (*pb.Blog, error) {
	fmt.Println("Retrieving Blog.....")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	blog := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(blog); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"No Blog Found",
		)
	}

	return documentToBlog(blog), nil
}
