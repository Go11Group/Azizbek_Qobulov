package main

import (
	"WeatherAndBusService/server/Services"
	"log"
	"net"

	pb "WeatherAndBusService/server/genproto/WeatherService"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &Services.WeatherServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
