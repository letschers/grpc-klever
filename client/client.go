package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DownVoteCrypto(ctx, &pb.DownVoteCryptoRequest{Id: 14})

	if err != nil {
		log.Println(err)
	}

	log.Println(response)
}
