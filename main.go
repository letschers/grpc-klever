package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/letschers/grpc-klever/server"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	server.StartServer()
}
