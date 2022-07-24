package main

import (
	"log"
	"net"

	"github.com/letschers/grpc-klever/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedCryptoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterCryptoServiceServer(s, &server{})
	log.Printf("server listening at port  %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
