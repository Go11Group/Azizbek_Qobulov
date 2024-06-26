// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: transport.proto

package BusService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportServiceClient interface {
	GetBusSchedule(ctx context.Context, in *BusScheduleRequest, opts ...grpc.CallOption) (*BusScheduleResponse, error)
	TrackBusLocation(ctx context.Context, in *BusLocationRequest, opts ...grpc.CallOption) (*BusLocationResponse, error)
	ReportTrafficJam(ctx context.Context, in *TrafficJamRequest, opts ...grpc.CallOption) (*TrafficJamResponse, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) GetBusSchedule(ctx context.Context, in *BusScheduleRequest, opts ...grpc.CallOption) (*BusScheduleResponse, error) {
	out := new(BusScheduleResponse)
	err := c.cc.Invoke(ctx, "/TransportService.TransportService/GetBusSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) TrackBusLocation(ctx context.Context, in *BusLocationRequest, opts ...grpc.CallOption) (*BusLocationResponse, error) {
	out := new(BusLocationResponse)
	err := c.cc.Invoke(ctx, "/TransportService.TransportService/TrackBusLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) ReportTrafficJam(ctx context.Context, in *TrafficJamRequest, opts ...grpc.CallOption) (*TrafficJamResponse, error) {
	out := new(TrafficJamResponse)
	err := c.cc.Invoke(ctx, "/TransportService.TransportService/ReportTrafficJam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility
type TransportServiceServer interface {
	GetBusSchedule(context.Context, *BusScheduleRequest) (*BusScheduleResponse, error)
	TrackBusLocation(context.Context, *BusLocationRequest) (*BusLocationResponse, error)
	ReportTrafficJam(context.Context, *TrafficJamRequest) (*TrafficJamResponse, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServiceServer struct {
}

func (UnimplementedTransportServiceServer) GetBusSchedule(context.Context, *BusScheduleRequest) (*BusScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusSchedule not implemented")
}
func (UnimplementedTransportServiceServer) TrackBusLocation(context.Context, *BusLocationRequest) (*BusLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackBusLocation not implemented")
}
func (UnimplementedTransportServiceServer) ReportTrafficJam(context.Context, *TrafficJamRequest) (*TrafficJamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTrafficJam not implemented")
}
func (UnimplementedTransportServiceServer) mustEmbedUnimplementedTransportServiceServer() {}

// UnsafeTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServiceServer will
// result in compilation errors.
type UnsafeTransportServiceServer interface {
	mustEmbedUnimplementedTransportServiceServer()
}

func RegisterTransportServiceServer(s grpc.ServiceRegistrar, srv TransportServiceServer) {
	s.RegisterService(&TransportService_ServiceDesc, srv)
}

func _TransportService_GetBusSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetBusSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportService.TransportService/GetBusSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetBusSchedule(ctx, req.(*BusScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_TrackBusLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).TrackBusLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportService.TransportService/TrackBusLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).TrackBusLocation(ctx, req.(*BusLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_ReportTrafficJam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficJamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).ReportTrafficJam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportService.TransportService/ReportTrafficJam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).ReportTrafficJam(ctx, req.(*TrafficJamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransportService.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusSchedule",
			Handler:    _TransportService_GetBusSchedule_Handler,
		},
		{
			MethodName: "TrackBusLocation",
			Handler:    _TransportService_TrackBusLocation_Handler,
		},
		{
			MethodName: "ReportTrafficJam",
			Handler:    _TransportService_ReportTrafficJam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
