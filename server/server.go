package server

import (
	"context"
	"log"
	"net"
	"os"

	db "github.com/letschers/grpc-klever/database"
	pb "github.com/letschers/grpc-klever/proto"
	"google.golang.org/grpc"
)

type CryptoServiceServer struct {
	pb.UnimplementedCryptoServiceServer
}

func (s *CryptoServiceServer) CreateCrypto(ctx context.Context, request *pb.CreateCryptoRequest) (*pb.CreateCryptoResponse, error) {
	data, err := db.CreateCrypto(request.GetName(), request.GetVotes())
	if err != nil {
		return nil, err
	}

	response := &pb.CreateCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) GetCrypto(ctx context.Context, request *pb.GetCryptoRequest) (*pb.GetCryptoResponse, error) {
	data, err := db.GetCrypto(request.GetId())
	if err != nil {
		return nil, err
	}

	response := &pb.GetCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) DeleteCrypto(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.DeleteCryptoResponse, error) {
	data, err := db.DeleteCrypto(request.GetId())
	if err != nil {
		return nil, err
	}

	response := &pb.DeleteCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) UpdateCrypto(ctx context.Context, request *pb.UpdateCryptoRequest) (*pb.UpdateCryptoResponse, error) {
	data, err := db.UpdateCrypto(request.Crypto)
	if err != nil {
		return nil, err
	}

	response := &pb.UpdateCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) UpVoteCrypto(ctx context.Context, request *pb.UpVoteCryptoRequest) (*pb.UpVoteCryptoResponse, error) {
	data, err := db.UpVoteCrypto(request.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.UpVoteCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) DownVoteCrypto(ctx context.Context, request *pb.DownVoteCryptoRequest) (*pb.DownVoteCryptoResponse, error) {
	data, err := db.DownVoteCrypto(request.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.DownVoteCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func StartServer() {
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
