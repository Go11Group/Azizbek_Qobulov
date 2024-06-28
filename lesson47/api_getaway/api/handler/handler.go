package handler

import (
	genproto "WeatherAndBusService/server/genproto/BusService"
	pb "WeatherAndBusService/server/genproto/WeatherService"
	"context"
)

type Handler struct {
	Weather   genproto.WeatherServiceClient
	Transport genproto.TransportServiceClient
}

func NewHandler(transport genproto.TransportServiceClient, weather genproto.WeatherServiceClient) *Handler {
	return &Handler{Transport: transport, Weather: weather}
}

func (h *Handler) GetWeather(ctx context.Context, req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	return &pb.WeatherResponse{
		Temperature: 25.5,
		Wet:         60,
		WindSpeed:   15,
	}, nil
}
