package main

import (
	"context"
	"fmt"

	pb "Calculator/proto/pb"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Println("Sum function invoked with ", in)
	return &pb.SumResponse{Result: in.A + in.B}, nil
}
