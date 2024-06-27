package Services

import (
	pb "WeatherAndBusService/server/genproto/BusService"
	pw "WeatherAndBusService/server/genproto/WeatherService"
)

type BusServer struct {
	pb.UnimplementedTransportServiceServer
}

type WeatherServer struct {
	pw.UnimplementedWeatherServiceServer
}
