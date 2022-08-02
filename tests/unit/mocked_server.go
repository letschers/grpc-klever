package tests

import (
	"context"
	"log"
	"net"

	pb "github.com/letschers/grpc-klever/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CryptoServiceServer struct {
	pb.UnimplementedCryptoServiceServer
}

func StartServer() {

	lis, err := net.Listen("tcp", ":8080")
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
	response := &pb.CreateCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    15,
			Name:  request.Name,
			Votes: 0,
		},
	}
	return response, nil
}

func (s *CryptoServiceServer) GetCrypto(ctx context.Context, request *pb.GetCryptoRequest) (*pb.GetCryptoResponse, error) {
	response := &pb.GetCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    request.Id,
			Name:  "TestCoin",
			Votes: 0,
		},
	}

	return response, nil
}

func (s *CryptoServiceServer) GetAllCrypto(ctx context.Context, request *pb.GetAllCryptoRequest) (*pb.GetAllCryptoResponse, error) {
	response := &pb.GetAllCryptoResponse{
		Cryptos: []*pb.Crypto{
			&pb.Crypto{
				Id:    15,
				Name:  "TestCoin",
				Votes: 0,
			},
			&pb.Crypto{
				Id:   10,
				Name: "TestCoin2",
			},
			&pb.Crypto{
				Id:    2,
				Name:  "TestCoin3",
				Votes: 0,
			},
		},
	}

	return response, nil
}

func (s *CryptoServiceServer) DeleteCrypto(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.DeleteCryptoResponse, error) {
	response := &pb.DeleteCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    request.Id,
			Name:  "Testcoin",
			Votes: 0,
		},
	}

	return response, nil
}

func (s *CryptoServiceServer) UpdateCrypto(ctx context.Context, request *pb.UpdateCryptoRequest) (*pb.UpdateCryptoResponse, error) {
	response := &pb.UpdateCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    request.Crypto.Id,
			Name:  request.Crypto.Name,
			Votes: request.Crypto.Votes,
		},
	}

	return response, nil
}

func (s *CryptoServiceServer) UpVoteCrypto(ctx context.Context, request *pb.UpVoteCryptoRequest) (*pb.UpVoteCryptoResponse, error) {
	response := &pb.UpVoteCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    request.Id,
			Name:  "Testcoin",
			Votes: 1,
		},
	}

	return response, nil
}

func (s *CryptoServiceServer) DownVoteCrypto(ctx context.Context, request *pb.DownVoteCryptoRequest) (*pb.DownVoteCryptoResponse, error) {
	response := &pb.DownVoteCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    request.Id,
			Name:  "Testcoin",
			Votes: -1,
		},
	}

	return response, nil
}

func (s *CryptoServiceServer) StreamCryptoVotes(request *pb.StreamCryptoVotesRequest, stream pb.CryptoService_StreamCryptoVotesServer) error {
	for {
		streamResponse := &pb.StreamCryptoVotesResponse{
			Votes: 1,
		}

		stream.Send(streamResponse)
	}

	return nil
}
