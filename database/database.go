package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	pb "github.com/letschers/grpc-klever/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IDatabase struct {
}

var (
	dbConnectionOnce  sync.Once
	DbConnection      *sql.DB
	dbConnectionError error
)

func GetDatabaseClient() error {
	dbConnectionOnce.Do(func() {
		conn, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))

		if err != nil {
			dbConnectionError = err
		}

		DbConnection = conn
	})

	return dbConnectionError
}

func (*IDatabase) CreateCrypto(name string, votes int32) (*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	var cryptoResponse pb.Crypto
	err := DbConnection.QueryRow("INSERT INTO cryptos(crypto_name, crypto_votes) VALUES($1, $2) RETURNING *",
		name, votes).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	return &cryptoResponse, nil
}

func (*IDatabase) GetCrypto(id int32) (*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	var cryptoResponse pb.Crypto
	err := DbConnection.QueryRow("SELECT * FROM cryptos WHERE crypto_id = $1", id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	return &cryptoResponse, nil
}

func (*IDatabase) GetAllCrypto() ([]*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	response, err := DbConnection.Query("SELECT * FROM cryptos")

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	cryptos := []*pb.Crypto{}
	for response.Next() {
		var crypto pb.Crypto
		err = response.Scan(&crypto.Id, &crypto.Name, &crypto.Votes)
		if err != nil {
			log.Fatal(err)
		}

		cryptos = append(cryptos, &crypto)
	}

	return cryptos, nil
}

func (*IDatabase) DeleteCrypto(id int32) (*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	var cryptoResponse pb.Crypto
	err := DbConnection.QueryRow("DELETE FROM cryptos WHERE crypto_id = $1 RETURNING *", id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	return &cryptoResponse, nil
}

func (*IDatabase) UpdateCrypto(request *pb.Crypto) (*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	var cryptoResponse pb.Crypto
	err := DbConnection.QueryRow("UPDATE cryptos SET crypto_name = $1, crypto_votes = $2 WHERE crypto_id = $3 RETURNING *", request.Name, request.Votes, request.Id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	return &cryptoResponse, nil
}

func (*IDatabase) UpVoteCrypto(request int32) (*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	var cryptoResponse pb.Crypto
	err := DbConnection.QueryRow("UPDATE cryptos SET crypto_votes = (crypto_votes + 1) WHERE crypto_id = $1 RETURNING *", request).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	return &cryptoResponse, nil
}

func (db *IDatabase) DownVoteCrypto(request int32) (*pb.Crypto, error) {
	if err := GetDatabaseClient(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}

	var cryptoResponse pb.Crypto
	err := DbConnection.QueryRow("UPDATE cryptos SET crypto_votes = crypto_votes - 1 WHERE crypto_id = $1 RETURNING *", request).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	return &cryptoResponse, nil
}
