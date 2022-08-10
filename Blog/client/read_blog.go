package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	fmt.Println(">> Hold a sec... Reading...")

	bid := &pb.BlogID{Id: id}

	res, err := c.ReadBlog(context.Background(), bid)
	if err != nil {
		fmt.Println("Error while retrieving blog")
		return nil
	}

	fmt.Println("Blog retrieved")
	return res
}
