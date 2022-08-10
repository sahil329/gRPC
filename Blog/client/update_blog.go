package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("Updating blog....")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Sneh",
		Title:    "Sneh Patel",
		Content:  "Ohh ya second time updated blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		fmt.Println("Error while Updating")
	}

	fmt.Println("Blog is updated!!!")
}
