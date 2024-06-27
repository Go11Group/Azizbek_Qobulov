package Services

import pb "WeatherAndBusService/server/genproto/BusService"

type BusServer struct {
	pb.UnimplementedTransportServiceServer
}

type WeatherServer struct {
	pb.UnimplementedTransportServiceServer
}

func (s *WeatherServer) mustEmbedUnimplementedWeatherServiceServer() {
	//TODO implement me
	panic("implement me")
}
