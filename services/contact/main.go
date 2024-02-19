// services/contact/main.go

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/SaraUlan/SaraUlan-03/pkg/store/postgres"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal"
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

	container := internal.NewContainer(db)

	go func() {
		router := initializeRouter(container)

		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	for {
		select {
		case <-time.After(time.Minute):
			log.Println("Performing background tasks...")
		}
	}
}

func initializeRouter(container *internal.Container) http.Handler {
	// You need to implement this part using your preferred router (e.g., Gorilla Mux, Chi, etc.)
	// Example using Gorilla Mux:
	// router := mux.NewRouter()
	// contactHandler := container.ContactDelivery.GetContactHandler()
	// router.HandleFunc("/contacts", contactHandler.HandleGetContacts).Methods("GET")
	// ...

	// Return your router
	// return router
	return nil
}
