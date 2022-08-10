package main

import (
	pb "Blog/proto/pb"
	"context"
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	fmt.Println("Listing blogs....")

	res, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Println("Error in retrieving blogs...")
	}

	for {
		blog, err := res.Recv()
		if err == io.EOF {
			color.Blue("Retrieved all !!!")
			break
		}

		if err != nil {
			color.Red("Error while receiving blog...")
		}

		fmt.Println(proto.MarshalTextString(blog))
		// time.Sleep(1 * time.Second)
	}
}
