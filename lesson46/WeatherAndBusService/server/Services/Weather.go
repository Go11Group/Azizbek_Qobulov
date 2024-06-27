package Services

import (
	pb "WeatherAndBusService/server/genproto/WeatherService"
	"context"
	"fmt"
)

func (s *WeatherServer) GetCurrentWeather(ctx context.Context, req *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	return &pb.CurrentWeatherResponse{
		Temperature: 25.5,
		Wet:         60,
		WindSpeed:   15,
	}, nil
}

func (s *WeatherServer) GetWeatherForecast(ctx context.Context, req *pb.WeatherForecastRequest) (*pb.WeatherForecastResponse, error) {
	forecasts := []*pb.Forecast{
		{Date: "2024-06-28", Temperature: 26.5, Wet: 55, WindSpeed: 12},
		{Date: "2024-06-29", Temperature: 27.0, Wet: 50, WindSpeed: 10},
	}
	return &pb.WeatherForecastResponse{Forecasts: forecasts}, nil
}

func (s *WeatherServer) ReportWeatherCondition(ctx context.Context, req *pb.WeatherConditionRequest) (*pb.WeatherConditionResponse, error) {
	fmt.Printf("Weather condition reported: %s - %s\n", req.Location, req.Condition)
	return &pb.WeatherConditionResponse{Success: true}, nil
}
