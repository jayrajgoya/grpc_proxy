package main

import (
	context "context"
	"log"
	"net"

	"github.com/grpc/geometry"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) GetArea(ctx context.Context, sides *geometry.Sides) (*geometry.Area, error) {

	log.Printf("Received get area request from client")
	return &geometry.Area{
		Area: int64(sides.Length) * int64(sides.Breadth),
	}, nil
}

func (s *Server) GetSum(ctx context.Context, Nums *geometry.Nums) (*geometry.Sum, error) {

	log.Printf("Received get sum request from client")
	return &geometry.Sum{
		Sum: int64(Nums.Num1) + int64(Nums.Num2),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	geometry.RegisterGeometryServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server gRPC server over port 9000: %v", err)
	}

}
