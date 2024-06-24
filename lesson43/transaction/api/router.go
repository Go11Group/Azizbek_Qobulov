package handler

import (
	"database/sql"
	"transaction/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(db *sql.DB) *http.Server {
	mux := gin.Default()
	h := handler.NewHandler(db)

	mux.POST("/transaction/create", h.CreateTransaction)
	mux.GET("/transaction/:id", h.GetTransactionById)
	mux.PUT("/transaction/:id", h.UpdateTransaction)
	mux.DELETE("/transaction/:id", h.DeleteTransaction)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
