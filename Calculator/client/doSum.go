package main

import (
	"context"
	"fmt"

	pb "Calculator/proto/pb"
)

func doSum(c pb.CalculatorServiceClient) {
	fmt.Println("Do-sum invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{A: 1, B: 10})

	if err != nil {
		fmt.Println("Could able to do summation ", err)
	}

	fmt.Printf("Sum : %d", res.Result)
}
