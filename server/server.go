package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/letschers/grpc-klever/proto"
	"google.golang.org/grpc"
)

type CryptoServiceServer struct {
	pb.UnimplementedCryptoServiceServer
}

func (s *CryptoServiceServer) CreateCrypto(ctx context.Context, request *pb.CreateCryptoRequest) (*pb.CreateCryptoResponse, error) {
	fmt.Printf("Request: %v", request)

	cryptoData := request.GetCrypto()

	return &pb.CreateCryptoResponse{Crypto: &pb.Crypto{
		Id:    cryptoData.GetId(),
		Name:  cryptoData.GetName(),
		Votes: cryptoData.GetVotes(),
	}}, nil
}

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCryptoServiceServer(server, &CryptoServiceServer{})
	log.Printf("Server listening at: %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
