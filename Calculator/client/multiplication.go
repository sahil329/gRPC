package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "Calculator/proto/pb"
)

func doMultiplication(c pb.CalculatorServiceClient) {
	fmt.Println("Do-Multiplication called !")

	stream, err := c.Multiplication(context.Background())

	if err != nil {
		fmt.Println("Error while client rpc call: ", err)
	}
	// numbers := []int64{15, 2, 12, 6, 84, 5, 33, 6}
	req := []*pb.MulRequest{
		{A: 15, B: 2},
		{A: 12, B: 6},
		{A: 84, B: 5},
		{A: 33, B: 6},
	}

	waitc := make(chan struct{})

	go func() {
		// var prev int64 = 0
		// for i := 0; i < len(numbers); i++ {
		// 	if i%2 != 0 {
		// 		stream.Send(&pb.MulRequest{
		// 			A: prev,
		// 			B: numbers[i],
		// 		})
		// 		time.Sleep(1 * time.Second)
		// 	} else {
		// 		prev = numbers[i]
		// 	}
		// }
		for _, r := range req {
			fmt.Println("Sending request: ", r)
			stream.Send(r)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("Error while receiving response ", err)
				break
			}

			fmt.Println("Received : ", res.Mult)
		}
		close(waitc)
	}()

	<-waitc

}
