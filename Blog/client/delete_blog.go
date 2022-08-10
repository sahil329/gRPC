package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"

	"github.com/fatih/color"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("Deleting blog...")

	bid := &pb.BlogID{Id: id}

	_, err := c.DeleteBlog(context.Background(), bid)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			fmt.Printf("Error code : %s\n", e.Code())
			fmt.Printf("Error Message : %s\n", e.Message())

			if e.Code() == codes.InvalidArgument {
				color.HiRed("Blog ID not exists !")
			}
		} else {
			color.Red("non-gRPC Error while deleting Blog !!!", err)
		}
		return
	}

	color.Green("Blog was Deleted !!")
}
