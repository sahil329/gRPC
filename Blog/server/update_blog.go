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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	fmt.Println("Updating blog in DB....")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Can't covert ID",
		)
	}

	blog := &BlogItem{
		AuthorId: in.AuthorId, //"Kumar",
		Title:    in.Title,    //"Update",
		Content:  in.Content,  //"See it",
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": blog},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not able to update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Could not found ID",
		)
	}

	return &emptypb.Empty{}, nil
}
