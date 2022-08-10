package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogID) (*emptypb.Empty, error) {
	fmt.Println("# Deleting Blog in DB...")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not convert ID",
		)
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Can't delete object in MongoDB",
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.Internal,
			"Blog not found !!!",
		)
	}

	return &emptypb.Empty{}, nil
}
