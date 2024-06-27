package main

import (
	"log"
	"net"

	"WeatherAndBusService/server/Services"
	pb "WeatherAndBusService/server/genproto/BusService"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransportServiceServer(s, &Services.BusServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
