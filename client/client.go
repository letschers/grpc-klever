package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/letschers/grpc-klever/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	conn, err := grpc.Dial("localhost"+os.Getenv("SERVER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCryptoServiceClient(conn)
	ctx := context.Background()

	response, err := client.StreamCryptoVotes(ctx, &pb.StreamCryptoVotesRequest{Id: 14})

	if err != nil {
		log.Println(err)
	}

	for {
		serverResponse, err := response.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		fmt.Printf("Server response: %v\n", serverResponse)
	}
}
