package main

import (
	"log"
	"net"

	"WeatherAndBusService/server/Services"
	pb "WeatherAndBusService/server/genproto/BusService"
	pw "WeatherAndBusService/server/genproto/WeatherService"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	b := grpc.NewServer()
	pb.RegisterTransportServiceServer(b, &Services.BusServer{})
	if err := b.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	w := grpc.NewServer()
	pw.RegisterWeatherServiceServer(w, &Services.WeatherServer{})
	if err = w.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
