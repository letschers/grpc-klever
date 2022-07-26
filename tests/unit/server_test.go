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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var cryptoServiceClient pb.CryptoServiceClient

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	go StartServer()
}

func connectToServer() (*grpc.ClientConn, context.CancelFunc) {
	conn, err := grpc.Dial("localhost"+os.Getenv("SERVER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect to server: %v", err)
	}

	client := pb.NewCryptoServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cryptoServiceClient = client

	return conn, cancel
}

func TestCreateCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.CreateCryptoRequest{
		Name:  "TestCoin",
		Votes: 0,
	}

	response, err := cryptoServiceClient.CreateCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to create the crypto. Error: %v", err)
	}

	if response.Crypto.Name != request.Name {
		t.Errorf("Data returned from server is incorrect. Expected %v, received %v", request.Name, response.Crypto.Name)
	}
}

func TestGetCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.GetCryptoRequest{
		Id: 15,
	}

	response, err := cryptoServiceClient.GetCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

	if response.Crypto.Id != request.Id {
		t.Errorf("Data returned from server is incorrect. Expected %v, received %v", request.Id, response.Crypto.Id)
	}
}

func TestGetAllCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.GetAllCryptoRequest{}

	_, err := cryptoServiceClient.GetAllCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

}

func TestDeleteCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.DeleteCryptoRequest{
		Id: 15,
	}

	response, err := cryptoServiceClient.DeleteCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

	if response.Crypto.Id != request.Id {
		t.Errorf("Data returned from server is incorrect. Expected %v, received %v", request.Id, response.Crypto.Id)
	}
}

func TestUpdateCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.UpdateCryptoRequest{
		Crypto: &pb.Crypto{
			Id:    15,
			Name:  "Testcoin",
			Votes: 0,
		},
	}

	response, err := cryptoServiceClient.UpdateCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

	if response.Crypto.Id != request.Crypto.Id {
		t.Errorf("Data returned from server is incorrect. Expected %v, received %v", request.Crypto.Id, response.Crypto.Id)
	}
}

func TestUpvoteCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.UpVoteCryptoRequest{
		Id: 15,
	}

	response, err := cryptoServiceClient.UpVoteCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

	if response.Crypto.Id != request.Id {
		t.Errorf("Data returned from server is incorrect. Expected %v, received %v", request.Id, response.Crypto.Id)
	}

	if response.Crypto.Votes != 1 {
		t.Errorf("Wasn't incremented correctly. Expected %v, received %v", 1, response.Crypto.Votes)
	}
}

func TestDownvoteCrypto(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.DownVoteCryptoRequest{
		Id: 15,
	}

	response, err := cryptoServiceClient.DownVoteCrypto(context.Background(), request)
	if err != nil {
		t.Errorf("Wasn't possible to get the crypto. Error: %v", err)
	}

	if response.Crypto.Id != request.Id {
		t.Errorf("Data returned from server is incorrect. Expected %v, received %v", request.Id, response.Crypto.Id)
	}

	if response.Crypto.Votes != -1 {
		t.Errorf("Wasn't incremented correctly. Expected %v, received %v", -1, response.Crypto.Votes)
	}
}

func TestStreamCryptoVotes(t *testing.T) {
	conn, cancel := connectToServer()
	defer conn.Close()
	defer cancel()

	request := &pb.StreamCryptoVotesRequest{
		Id: 14,
	}

	response, err := cryptoServiceClient.StreamCryptoVotes(context.Background(), request)
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

	if serverResponse.Votes != 1 {
		t.Errorf("Value returned from server wasn't correct. Expected: %v, received: %v", 1, serverResponse.Votes)
	}
}
