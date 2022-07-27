package server

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/letschers/grpc-klever/database"
	pb "github.com/letschers/grpc-klever/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CryptoServiceServer struct {
	pb.UnimplementedCryptoServiceServer
}

func StartServer() {
	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCryptoServiceServer(server, &CryptoServiceServer{})
	log.Printf("Server listening at: %v", lis.Addr())
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *CryptoServiceServer) CreateCrypto(ctx context.Context, request *pb.CreateCryptoRequest) (*pb.CreateCryptoResponse, error) {
	db := database.IDatabase{}
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
	db := database.IDatabase{}
	data, err := db.GetCrypto(request.GetId())
	if err != nil {
		return nil, err
	}

	response := &pb.GetCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) GetAllCrypto(ctx context.Context, request *pb.GetAllCryptoRequest) (*pb.GetAllCryptoResponse, error) {
	db := database.IDatabase{}
	data, err := db.GetAllCrypto()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllCryptoResponse{
		Cryptos: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) DeleteCrypto(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.DeleteCryptoResponse, error) {
	db := database.IDatabase{}
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
	db := database.IDatabase{}
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
	db := database.IDatabase{}
	data, err := db.UpVoteCrypto(request.Crypto.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.UpVoteCryptoResponse{
		Crypto: data,
	}

	return response, nil
}

func (s *CryptoServiceServer) DownVoteCrypto(ctx context.Context, request *pb.DownVoteCryptoRequest) (*pb.DownVoteCryptoResponse, error) {
	db := database.IDatabase{}
	data, err := db.DownVoteCrypto(request.Crypto.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.DownVoteCryptoResponse{
		Crypto: data,
	}

	return response, nil
}
