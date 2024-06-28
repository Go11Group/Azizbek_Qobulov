// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: weather.proto

package WeatherService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CurrentWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CurrentWeatherRequest) Reset() {
	*x = CurrentWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentWeatherRequest) ProtoMessage() {}

func (x *CurrentWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentWeatherRequest.ProtoReflect.Descriptor instead.
func (*CurrentWeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentWeatherRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type CurrentWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float32 `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Wet         float32 `protobuf:"fixed32,2,opt,name=wet,proto3" json:"wet,omitempty"`
	WindSpeed   float32 `protobuf:"fixed32,3,opt,name=WindSpeed,proto3" json:"WindSpeed,omitempty"`
}

func (x *CurrentWeatherResponse) Reset() {
	*x = CurrentWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentWeatherResponse) ProtoMessage() {}

func (x *CurrentWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentWeatherResponse.ProtoReflect.Descriptor instead.
func (*CurrentWeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentWeatherResponse) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *CurrentWeatherResponse) GetWet() float32 {
	if x != nil {
		return x.Wet
	}
	return 0
}

func (x *CurrentWeatherResponse) GetWindSpeed() float32 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

type WeatherForecastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *WeatherForecastRequest) Reset() {
	*x = WeatherForecastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherForecastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherForecastRequest) ProtoMessage() {}

func (x *WeatherForecastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherForecastRequest.ProtoReflect.Descriptor instead.
func (*WeatherForecastRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{2}
}

func (x *WeatherForecastRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type WeatherForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forecasts []*Forecast `protobuf:"bytes,1,rep,name=forecasts,proto3" json:"forecasts,omitempty"`
}

func (x *WeatherForecastResponse) Reset() {
	*x = WeatherForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherForecastResponse) ProtoMessage() {}

func (x *WeatherForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherForecastResponse.ProtoReflect.Descriptor instead.
func (*WeatherForecastResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{3}
}

func (x *WeatherForecastResponse) GetForecasts() []*Forecast {
	if x != nil {
		return x.Forecasts
	}
	return nil
}

type Forecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Temperature float32 `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Wet         float32 `protobuf:"fixed32,3,opt,name=wet,proto3" json:"wet,omitempty"`
	WindSpeed   float32 `protobuf:"fixed32,4,opt,name=WindSpeed,proto3" json:"WindSpeed,omitempty"`
}

func (x *Forecast) Reset() {
	*x = Forecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forecast) ProtoMessage() {}

func (x *Forecast) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forecast.ProtoReflect.Descriptor instead.
func (*Forecast) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{4}
}

func (x *Forecast) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Forecast) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Forecast) GetWet() float32 {
	if x != nil {
		return x.Wet
	}
	return 0
}

func (x *Forecast) GetWindSpeed() float32 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

type WeatherConditionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location  string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Condition string `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *WeatherConditionRequest) Reset() {
	*x = WeatherConditionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherConditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherConditionRequest) ProtoMessage() {}

func (x *WeatherConditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherConditionRequest.ProtoReflect.Descriptor instead.
func (*WeatherConditionRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{5}
}

func (x *WeatherConditionRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *WeatherConditionRequest) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

type WeatherConditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *WeatherConditionResponse) Reset() {
	*x = WeatherConditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherConditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherConditionResponse) ProtoMessage() {}

func (x *WeatherConditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherConditionResponse.ProtoReflect.Descriptor instead.
func (*WeatherConditionResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{6}
}

func (x *WeatherConditionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x33, 0x0a, 0x15, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a, 0x16, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x77, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x77,
	0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x22, 0x34, 0x0a, 0x16, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x17, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x09,
	0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x08, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x77, 0x65, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x53, 0x0a, 0x17, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x34, 0x0a, 0x18, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc8, 0x02, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x25,
	0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_weather_proto_goTypes = []interface{}{
	(*CurrentWeatherRequest)(nil),    // 0: WeatherService.CurrentWeatherRequest
	(*CurrentWeatherResponse)(nil),   // 1: WeatherService.CurrentWeatherResponse
	(*WeatherForecastRequest)(nil),   // 2: WeatherService.WeatherForecastRequest
	(*WeatherForecastResponse)(nil),  // 3: WeatherService.WeatherForecastResponse
	(*Forecast)(nil),                 // 4: WeatherService.Forecast
	(*WeatherConditionRequest)(nil),  // 5: WeatherService.WeatherConditionRequest
	(*WeatherConditionResponse)(nil), // 6: WeatherService.WeatherConditionResponse
}
var file_weather_proto_depIdxs = []int32{
	4, // 0: WeatherService.WeatherForecastResponse.forecasts:type_name -> WeatherService.Forecast
	0, // 1: WeatherService.WeatherService.GetCurrentWeather:input_type -> WeatherService.CurrentWeatherRequest
	2, // 2: WeatherService.WeatherService.GetWeatherForecast:input_type -> WeatherService.WeatherForecastRequest
	5, // 3: WeatherService.WeatherService.ReportWeatherCondition:input_type -> WeatherService.WeatherConditionRequest
	1, // 4: WeatherService.WeatherService.GetCurrentWeather:output_type -> WeatherService.CurrentWeatherResponse
	3, // 5: WeatherService.WeatherService.GetWeatherForecast:output_type -> WeatherService.WeatherForecastResponse
	6, // 6: WeatherService.WeatherService.ReportWeatherCondition:output_type -> WeatherService.WeatherConditionResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentWeatherRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentWeatherResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherForecastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherForecastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forecast); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherConditionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherConditionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}
