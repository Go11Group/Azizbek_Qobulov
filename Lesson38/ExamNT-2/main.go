package main

import (
	handler "ExamNT/Handler"
	"ExamNT/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()    // Connecting to PostgreSQL database
	if err != nil {
		panic(err)
	}

	defer db.Close()                  // Defer closing the database connection after main function exits

	handler := handler.NewHandler(db) // Creating a new handler with the database connection

	handler.SetupRoutes()             // Setting up HTTP routes for handling requests
}
