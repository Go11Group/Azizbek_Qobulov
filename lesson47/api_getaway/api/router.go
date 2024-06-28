package api

import (
	"api/api/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pw "api/genproto/WeatherService"
	pb "api/genproto/BusService"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	weatherClient := pw.NewWeatherServiceClient(conn)
	transportClient := pb.NewTransportServiceClient(conn)
	h := handler.NewHandler(transportClient,weatherClient)

	router.POST("/reportWeather", h.ReportWeather)
	router.GET("/getCurrentWeather", h.GetCurrentWeather)
	router.GET("/getWeatherForecast", h.GetWeatherForecast)

	return router
}
