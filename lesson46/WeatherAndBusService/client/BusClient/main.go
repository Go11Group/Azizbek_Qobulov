package main

import (
	pb "WeatherAndBusService/server/genproto/BusService"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to transport service: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransportServiceClient(conn)
	var num int
	fmt.Println("1.getBusSchedule\n2.trackBusLocation\n3.reportTrafficJam")
	fmt.Println("ENTER NUMBER ")
	fmt.Scanln(&num)
	switch num {
	case 1:
		getBusSchedule(client)
	case 2:
		trackBusLocation(client)
	case 3:
		reportTrafficJam(client)
	default:
		log.Fatal("Please enter a valid number.")
	}
}

func getBusSchedule(client pb.TransportServiceClient) {
	req := &pb.BusScheduleRequest{BusNumber: "12"}
	res, err := client.GetBusSchedule(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get bus schedule: %v", err)
	}
	fmt.Println("Bus Schedule for Bus 12:")
	for _, time := range res.Schedule {
		fmt.Println(time)
	}
}

func trackBusLocation(client pb.TransportServiceClient) {
	req := &pb.BusLocationRequest{BusNumber: "12"}
	res, err := client.TrackBusLocation(context.Background(), req)
	if err != nil {
		log.Fatalf("could not track bus location: %v", err)
	}
	fmt.Printf("Bus 12 Location: Latitude %.4f, Longitude %.4f\n", res.Latitude, res.Longitude)
}

func reportTrafficJam(client pb.TransportServiceClient) {
	req := &pb.TrafficJamRequest{
		Location:    "Main Street",
		Description: "Heavy traffic due to accident",
	}
	res, err := client.ReportTrafficJam(context.Background(), req)
	if err != nil {
		log.Fatalf("could not report traffic jam: %v", err)
	}
	if res.Success {
		fmt.Println("Traffic jam reported successfully")
	} else {
		fmt.Println("Failed to report traffic jam")
	}
}
