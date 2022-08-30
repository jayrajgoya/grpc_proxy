package main

import (
	"context"
	"log"

	"github.com/grpc/geometry"
	"github.com/grpc/thinLizzy"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8092", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	s := thinLizzy.NewCallEndpointClient(conn)

	sides := &geometry.Sides{Length: 10, Breadth: 20}

	// nums := &geometry.Nums{Num1: 10, Num2: 20}

	payload := &thinLizzy.SubmitBody{
		Endpoint:    "dummy",
		Api:         "dummy",
		AreaPayload: sides,
	}

	resp1, err := s.JobSubmit(context.Background(), payload)

	if err != nil {
		log.Fatalf("error when calling JobSubmit: %v", err)
	}
	log.Printf("response from JobSumit is", resp1.JobId)
}
