package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"

	"github.com/fatih/color"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("Deleting blog...")

	bid := &pb.BlogID{Id: id}

	res, err := c.DeleteBlog(context.Background(), bid)
	if err != nil || res == nil {
		color.Red("Error while deleting Blog !!!", err)
		return
	}

	color.Green("Blog was Deleted !!")
}
