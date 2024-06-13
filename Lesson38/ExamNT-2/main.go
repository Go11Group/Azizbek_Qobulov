package main

import (
	handler "ExamNT/Handler"
	"ExamNT/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	handler := handler.NewHandler(db)

	handler.SetupRoutes()
}
