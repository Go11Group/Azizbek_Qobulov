package handler

import (
	pb "api/genproto/BusService"
	pw "api/genproto/WeatherService"
)

type Handler struct {
	Weather   pw.WeatherServiceClient
	Transport pb.TransportServiceClient
}

func NewHandler(transport pb.TransportServiceClient, weather pw.WeatherServiceClient) *Handler {
	return &Handler{Transport: transport, Weather: weather}
}
