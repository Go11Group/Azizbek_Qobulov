package api

import (
	"database/sql"
	"user_service/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()
	h := handler.NewHandler(db)
	mux.POST("/user/create", h.CreateUser)
	mux.GET("/user/:id", h.GetUserById)
	return &http.Server{Handler: mux, Addr: ":8080"}
}
