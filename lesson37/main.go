package main

import (
	"example/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	p := postgres.NewPersonRepo(db)
	r.GET("/person/getAll", p.GetAll)
	r.Run()
}
