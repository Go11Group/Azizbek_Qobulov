<<<<<<< HEAD
package main

import (
	"log"
	"net/http"

	"github.com/Go11Group/at_lesson/lesson34/handler"
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"github.com/gorilla/mux"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	st := postgres.NewStudentRepo(db)

	r := mux.NewRouter()
	h := handler.NewHandler(st)

	h.SetupRoutes(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
=======
package main

import (
	"github.com/Go11Group/at_lesson/lesson34/handler"
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	st := postgres.NewStudentRepo(db)

	server := handler.NewHandler(handler.Handler{Talaba: st})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
>>>>>>> origin/main
