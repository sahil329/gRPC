package main

import (
	"fmt"
	"io"

	pb "Calculator/proto/pb"
)

func (s *Server) Multiplication(stream pb.CalculatorService_MultiplicationServer) error {
	fmt.Println("Multiplication invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Println("Error while multiplying ", err)
		}

		res := int64(req.A * req.B)

		err = stream.Send(&pb.MulResponse{
			Mult: res,
		})

		if err != nil {
			fmt.Println("Error sending response back ", err)
		}

		// return nil
	}
}
