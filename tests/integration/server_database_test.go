package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/letschers/grpc-klever/proto"
	s "github.com/letschers/grpc-klever/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var cryptoServiceClient pb.CryptoServiceClient

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	go s.StartServer()
}

func connectToServer() (context.Context, *grpc.ClientConn, context.CancelFunc) {
	conn, err := grpc.Dial("localhost"+os.Getenv("SERVER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect to server: %v", err)
	}

	client := pb.NewCryptoServiceClient(conn)
	cryptoServiceClient = client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx, conn, cancel
}

func TestCreateCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.CreateCryptoRequest{
		Name:  "Testcoin",
		Votes: 0,
	}

	response, err := cryptoServiceClient.CreateCrypto(ctx, request)
	if err != nil {
		t.Errorf("Wasn't possible to create the crypto. Error: %v", err)
	}

	if response.Crypto.Name != "Testcoin" {
		t.Errorf("Data received from server is wrong. Expected: %v, received: %v", request.Name, response.Crypto.Name)
	}

	if _, err := cryptoServiceClient.GetCrypto(ctx, &pb.GetCryptoRequest{Id: response.Crypto.Id}); err != nil {
		t.Errorf("Wasn't possible to retrieve the new created crypto. Error: %v", err)
	}

	_, _ = cryptoServiceClient.DeleteCrypto(ctx, &pb.DeleteCryptoRequest{Id: response.Crypto.Id})
}

func TestGetCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.GetCryptoRequest{
		Id: 14,
	}

	response, err := cryptoServiceClient.GetCrypto(ctx, request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

	if response.Crypto.Name != "Bitcoin" {
		t.Errorf("Data received from server is wrong. Expected: Bitcoin, received: %v", response.Crypto.Name)
	}
}

func TestGetAllCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.GetAllCryptoRequest{}

	_, err := cryptoServiceClient.GetAllCrypto(ctx, request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

}

func TestDeleteCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	createRequest := &pb.CreateCryptoRequest{
		Name:  "Testcoin",
		Votes: 0,
	}

	createResponse, err := cryptoServiceClient.CreateCrypto(ctx, createRequest)
	if err != nil {
		t.Errorf("Wasn't possible to create the crypto. Error: %v", err)
	}

	deleteRequest := &pb.DeleteCryptoRequest{
		Id: createResponse.Crypto.Id,
	}

	response, err := cryptoServiceClient.DeleteCrypto(ctx, deleteRequest)
	if err != nil {
		t.Errorf("Wasn't possible to delete the crypto. Error: %v", err)
	}

	if _, err := cryptoServiceClient.GetCrypto(ctx, &pb.GetCryptoRequest{Id: response.Crypto.Id}); err == nil {
		t.Errorf("Crypto wasn't correctly deleted. Error: %v", err)
	}
}

func TestUpdateCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	createRequest := &pb.CreateCryptoRequest{
		Name:  "Testcoin",
		Votes: 0,
	}

	createResponse, err := cryptoServiceClient.CreateCrypto(ctx, createRequest)
	if err != nil {
		t.Errorf("Wasn't possible to create the crypto. Error: %v", err)
	}

	updateRequest := &pb.UpdateCryptoRequest{
		Crypto: &pb.Crypto{
			Id:    createResponse.Crypto.Id,
			Name:  "AlteredCoin",
			Votes: 0,
		},
	}

	response, err := cryptoServiceClient.UpdateCrypto(ctx, updateRequest)
	if err != nil {
		t.Errorf("Wasn't possible to delete the crypto. Error: %v", err)
	}

	if response, _ := cryptoServiceClient.GetCrypto(ctx, &pb.GetCryptoRequest{Id: response.Crypto.Id}); response.Crypto.Name != "AlteredCoin" {
		t.Errorf("Crypto wasn't correctly updated. Error: %v", err)
	}

	_, _ = cryptoServiceClient.DeleteCrypto(ctx, &pb.DeleteCryptoRequest{Id: response.Crypto.Id})
}

func TestUpvoteCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	createRequest := &pb.CreateCryptoRequest{
		Name:  "Testcoin",
		Votes: 0,
	}

	createResponse, err := cryptoServiceClient.CreateCrypto(ctx, createRequest)
	if err != nil {
		t.Errorf("Wasn't possible to create the crypto. Error: %v", err)
	}

	upvoteRequest := &pb.UpVoteCryptoRequest{
		Id: createResponse.Crypto.Id,
	}

	response, err := cryptoServiceClient.UpVoteCrypto(ctx, upvoteRequest)
	if err != nil {
		t.Errorf("Wasn't possible to upvote the crypto. Error: %v", err)
	}

	if response, _ := cryptoServiceClient.GetCrypto(ctx, &pb.GetCryptoRequest{Id: response.Crypto.Id}); createRequest.Votes+1 != response.Crypto.Votes {
		t.Errorf("Crypto wasn't correctly upvoted. Error: %v", err)
	}

	_, _ = cryptoServiceClient.DeleteCrypto(ctx, &pb.DeleteCryptoRequest{Id: response.Crypto.Id})
}

func TestDownvoteCrypto(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	createRequest := &pb.CreateCryptoRequest{
		Name:  "Testcoin",
		Votes: 0,
	}

	createResponse, err := cryptoServiceClient.CreateCrypto(ctx, createRequest)
	if err != nil {
		t.Errorf("Wasn't possible to create the crypto. Error: %v", err)
	}

	downVoteRequest := &pb.DownVoteCryptoRequest{
		Id: createResponse.Crypto.Id,
	}

	response, err := cryptoServiceClient.DownVoteCrypto(ctx, downVoteRequest)
	if err != nil {
		t.Errorf("Wasn't possible to upvote the crypto. Error: %v", err)
	}

	if response, _ := cryptoServiceClient.GetCrypto(ctx, &pb.GetCryptoRequest{Id: response.Crypto.Id}); createRequest.Votes-1 != response.Crypto.Votes {
		t.Errorf("Crypto wasn't correctly upvoted.")
	}

	_, _ = cryptoServiceClient.DeleteCrypto(ctx, &pb.DeleteCryptoRequest{Id: response.Crypto.Id})
}

func TestStreamCryptoVotes(t *testing.T) {
	ctx, conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.StreamCryptoVotesRequest{
		Id: 14,
	}

	response, err := cryptoServiceClient.StreamCryptoVotes(ctx, request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto stream. Error: %v", err)
	}

	if response == nil {
		t.Errorf("Server stream wasn't received")
	}

	serverResponse, err := response.Recv()
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto stream. Error: %v", err)
	}

	if serverResponse == nil {
		t.Errorf("Server stream wasn't received. Value: %v", serverResponse)
	}

	if response, _ := cryptoServiceClient.GetCrypto(ctx, &pb.GetCryptoRequest{Id: request.Id}); response.Crypto.Votes != serverResponse.Votes {
		t.Errorf("Value returned from server wasn't correct. Expected: %v, received: %v", response.Crypto.Id, serverResponse.Votes)

	}
}
