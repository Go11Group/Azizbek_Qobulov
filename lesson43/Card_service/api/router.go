package api

import (
	"Card_service/api/handler"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()
	h := handler.NewHandler(db)
	mux.POST("/card/create", h.CreateCard)
	mux.GET("/card/:id", h.GetCardById)
	mux.PUT("/card/:id", h.UpdateCard)
	mux.DELETE("/card/:id", h.DeleteCard)
	return &http.Server{Handler: mux, Addr: ":8080"}
}
