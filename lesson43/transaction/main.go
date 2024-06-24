package main

import (
	"log"
	"transaction/api"
	"transaction/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	server := handler.Routers(db)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
