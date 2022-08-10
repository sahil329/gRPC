package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"
)

func createBlog(c pb.BlogServiceClient) string {
	fmt.Println("----------- Creating blog invoked ---------------")

	blog := &pb.Blog{
		AuthorId: "Sahil",
		Title:    "Rock Go",
		Content:  "Learn gRPC with GO",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		fmt.Println("Not able to create blog")
	}

	fmt.Println("Created Blog..!!!")
	return res.Id
}
