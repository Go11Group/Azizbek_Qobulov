package main

import (
	"log"
	"user_service/api"
	"user_service/storage/postgres"
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
