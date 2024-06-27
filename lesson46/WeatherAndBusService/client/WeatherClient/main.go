package main

import (
	pb "WeatherAndBusService/server/genproto/WeatherService"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to weather service: %v", err)
	}
	defer conn.Close()
	client := pb.NewWeatherServiceClient(conn)
	var num int
	fmt.Println("1.GetCurrentWeather\n2.GetWeatherForecast\n3.ReportWeatherCondition")
	fmt.Print("ENTER NUMBER:")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		getCurrentWeather(client)
	case 2:
		getWeatherForecast(client)
	case 3:
		reportWeatherCondition(client)
	default:
		log.Fatal("Please enter a valid number")
	}
}

func getCurrentWeather(client pb.WeatherServiceClient) {
	req := &pb.CurrentWeatherRequest{Location: "Tashkent"}
	res, err := client.GetCurrentWeather(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get current weather: %v", err)
	}
	fmt.Printf("Current Weather in Tashkent: %.2f°C, %.2f%% wet, %.2f m/s wind speed\n", res.Temperature, res.Wet, res.WindSpeed)
}

func getWeatherForecast(client pb.WeatherServiceClient) {
	req := &pb.WeatherForecastRequest{Location: "Tashkent"}
	res, err := client.GetWeatherForecast(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get weather forecast: %v", err)
	}
	fmt.Println("Weather Forecast in Tashkent:")
	for _, forecast := range res.Forecasts {
		fmt.Printf("Date: %s - %.2f°C, %.2f%% wet, %.2f m/s wind speed\n", forecast.Date, forecast.Temperature, forecast.Wet, forecast.WindSpeed)
	}
}

func reportWeatherCondition(client pb.WeatherServiceClient) {
	req := &pb.WeatherConditionRequest{
		Location:  "Tashkent",
		Condition: "Sunny",
	}
	res, err := client.ReportWeatherCondition(context.Background(), req)
	if err != nil {
		log.Fatalf("could not report weather condition: %v", err)
	}
	if res.Success {
		fmt.Println("Weather condition reported successfully")
	} else {
		fmt.Println("Failed to report weather condition")
	}
}
