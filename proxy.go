package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc/geometry"
	"github.com/grpc/thinLizzy"
	"google.golang.org/grpc"
)

type Serve struct{}

func (s *Serve) JobSubmit(ctx context.Context, payload *thinLizzy.SubmitBody) (*thinLizzy.Response, error) {
	// go func() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	g := geometry.NewGeometryClient(conn)

	resp, err2 := g.GetArea(ctx, payload.AreaPayload)

	if err2 != nil {
		log.Fatalf("error in calling Get Area: %v", err)
	}
	log.Println("response from Get Area is", resp.Area)

	// }()
	return &thinLizzy.Response{JobId: "123"}, nil
}

func main() {

	listen, err := net.Listen("tcp", ":8092")

	if err != nil {
		log.Fatalf("unable to dial tcp %v", err)
	}
	s := Serve{}

	grpcServer := grpc.NewServer()

	thinLizzy.RegisterCallEndpointServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server gRPC server over port 8092: %v", err)
	}

}
