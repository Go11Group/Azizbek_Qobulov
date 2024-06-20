package main

import (
	"fmt"
	"person/handler"
	"person/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := postgres.NewPersonRepo(db)
	h := handler.Handler{Person: p}
	server := handler.NewHandler(h)
	// go handler.Get()
	// go handler.Post()
	// go handler.Put()
	// go handler.Delete()
	// go handler.GetAll()
	fmt.Println("Server is listening on port 8080...")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}