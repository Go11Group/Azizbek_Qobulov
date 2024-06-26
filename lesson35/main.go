<<<<<<< HEAD
package main

import (
	"log"

	"github.com/Go11Group/at_lesson/lesson34/handler"
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	st := postgres.NewStudentRepo(db)

	r := gin.Default()
	h := handler.NewHandler(st)

	h.SetupRoutes(r)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
=======
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
		log.Fatal("Could not connect to the database: ", err)
	}
	st := postgres.NewStudentRepo(db)

	r := mux.NewRouter()
	h := handler.NewHandler(st)

	h.SetupRoutes(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
>>>>>>> origin/main
