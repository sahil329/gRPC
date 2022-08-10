package main

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "Blog/proto/pb"
)

var addr = "127.0.0.1:9080"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to dial connection ", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// id := createBlog(c)
	// fmt.Println("BLog: ", readBlog(c, id), "\n")
	// readBlog(c, id)
	// updateBlog(c, "62f2bc148720c72478cbed8b") //id)
	listBlogs(c)
	deleteBlog(c, "62f2bc148720c72478cbed8b")
	listBlogs(c)
}
