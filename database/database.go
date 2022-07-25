package database

import (
	"database/sql"
	"fmt"
	"os"

	pb "github.com/letschers/grpc-klever/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetDatabaseClient() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateCrypto(name string, votes int32) (*pb.Crypto, error) {
	db, err := GetDatabaseClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var cryptoResponse pb.Crypto
	err = db.QueryRow("INSERT INTO cryptos(crypto_name, crypto_votes) VALUES($1, $2) RETURNING *", name, votes).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error here: %v", err))
	}

	return &cryptoResponse, nil
}

func GetCrypto(id int32) (*pb.Crypto, error) {
	db, err := GetDatabaseClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var cryptoResponse pb.Crypto
	err = db.QueryRow("SELECT * FROM cryptos WHERE crypto_id = $1", id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &cryptoResponse, nil
}

func DeleteCrypto(id int32) (*pb.Crypto, error) {
	db, err := GetDatabaseClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var cryptoResponse pb.Crypto
	err = db.QueryRow("DELETE FROM cryptos WHERE crypto_id = $1 RETURNING *", id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error here: %v", err))
	}

	return &cryptoResponse, nil
}

func UpdateCrypto(request *pb.Crypto) (*pb.Crypto, error) {
	db, err := GetDatabaseClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var cryptoResponse pb.Crypto
	err = db.QueryRow("UPDATE cryptos SET crypto_name = $1, crypto_votes = $2 WHERE crypto_id = $3 RETURNING *", request.Name, request.Votes, request.Id).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error here: %v", err))
	}

	return &cryptoResponse, nil
}

func UpVoteCrypto(request int32) (*pb.Crypto, error) {
	db, err := GetDatabaseClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var cryptoResponse pb.Crypto
	err = db.QueryRow("UPDATE cryptos SET crypto_votes = (crypto_votes + 1) WHERE crypto_id = $1 RETURNING *", request).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error here: %v", err))
	}

	return &cryptoResponse, nil
}

func DownVoteCrypto(request int32) (*pb.Crypto, error) {
	db, err := GetDatabaseClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var cryptoResponse pb.Crypto
	err = db.QueryRow("UPDATE cryptos SET crypto_votes = crypto_votes - 1 WHERE crypto_id = $1 RETURNING *", request).Scan(&cryptoResponse.Id, &cryptoResponse.Name, &cryptoResponse.Votes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error here: %v", err))
	}

	return &cryptoResponse, nil
}
