package main

import (
	"log"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/api"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	server := api.Routes(db)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
