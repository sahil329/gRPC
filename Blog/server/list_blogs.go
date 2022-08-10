package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	fmt.Println("Listing Blog from DB....")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Error while Listing Blogs...",
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		blog := &BlogItem{}
		err := cur.Decode(blog)

		if err != nil {
			return status.Errorf(codes.Internal, "Error while Decoding")
		}

		stream.Send(documentToBlog(blog))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, "Error in cur..")
	}

	return nil
}
