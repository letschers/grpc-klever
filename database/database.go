package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Crypto struct {
	id    int64
	name  string
	votes string
}

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Printf("Error: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM cryptos")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var crypto Crypto

		err = rows.Scan(&crypto.id, &crypto.name, &crypto.votes)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d\t%s\t%s \n", crypto.id, crypto.name, crypto.votes)
	}
}
