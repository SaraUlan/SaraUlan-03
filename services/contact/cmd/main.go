package main

import (
	"log"

	"github.com/SaraUlan/SaraUlan-03/pkg/store/postgres"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "12345678"
	dbname := "go_ass3"

	db, err := postgres.ConnectToPostgres(host, port, user, password, dbname)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Connected to PostgreSQL from main!")
}
