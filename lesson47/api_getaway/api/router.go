package api

import (
	"api/api/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine{
	router := gin.Default()

	weatherClient := genproto.NewWeatherServiceClient(conn)
	transportClient := genproto.NewTransportServiceClient(conn)
	handler.Handler{Weather: weatherClient, Transport: transportClient}


	router.GET("/get",handler.GetWeather)

	return router
}