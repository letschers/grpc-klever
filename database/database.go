package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/letschers/grpc-klever/models"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getCrypto(id int64) (*models.Crypto, error) {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var cryptoResponse models.Crypto
	err = db.QueryRow("SELECT * FROM cryptos WHERE crypto_id = $1", id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &cryptoResponse, nil
}

func main() {
	response, err := getCrypto(2)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Println(response)

}
