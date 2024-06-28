package Services

import (
	pb "WeatherAndBusService/server/genproto/BusService"
	"context"
	"fmt"
)

func (s *BusServer) GetBusSchedule(ctx context.Context, req *pb.BusScheduleRequest) (*pb.BusScheduleResponse, error) {
	schedule := []string{"08:00", "12:00", "16:00", "20:00"}
	return &pb.BusScheduleResponse{Schedule: schedule}, nil
}

func (s *BusServer) TrackBusLocation(ctx context.Context, req *pb.BusLocationRequest) (*pb.BusLocationResponse, error) {
	return &pb.BusLocationResponse{
		Latitude:  40.7128,
		Longitude: -74.0060,
	}, nil
}

func (s *BusServer) ReportTrafficJam(ctx context.Context, req *pb.TrafficJamRequest) (*pb.TrafficJamResponse, error) {
	fmt.Printf("Traffic jam reported at %s: %s\n", req.Location, req.Description)
	return &pb.TrafficJamResponse{Success: true}, nil
}
