// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: search_flight.proto

package server_proto

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

type SearchOneWayFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepartureDate           string `protobuf:"bytes,1,opt,name=DepartureDate,proto3" json:"DepartureDate,omitempty"`
	OriginLocationCode      string `protobuf:"bytes,2,opt,name=OriginLocationCode,proto3" json:"OriginLocationCode,omitempty"`
	DestinationLocationCode string `protobuf:"bytes,3,opt,name=DestinationLocationCode,proto3" json:"DestinationLocationCode,omitempty"`
}

func (x *SearchOneWayFlightRequest) Reset() {
	*x = SearchOneWayFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOneWayFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOneWayFlightRequest) ProtoMessage() {}

func (x *SearchOneWayFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOneWayFlightRequest.ProtoReflect.Descriptor instead.
func (*SearchOneWayFlightRequest) Descriptor() ([]byte, []int) {
	return file_search_flight_proto_rawDescGZIP(), []int{0}
}

func (x *SearchOneWayFlightRequest) GetDepartureDate() string {
	if x != nil {
		return x.DepartureDate
	}
	return ""
}

func (x *SearchOneWayFlightRequest) GetOriginLocationCode() string {
	if x != nil {
		return x.OriginLocationCode
	}
	return ""
}

func (x *SearchOneWayFlightRequest) GetDestinationLocationCode() string {
	if x != nil {
		return x.DestinationLocationCode
	}
	return ""
}

type SearchOneWayFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepartureTime        string `protobuf:"bytes,1,opt,name=DepartureTime,proto3" json:"DepartureTime,omitempty"`
	ArrivalTime          string `protobuf:"bytes,2,opt,name=ArrivalTime,proto3" json:"ArrivalTime,omitempty"`
	FlightIdentification string `protobuf:"bytes,3,opt,name=FlightIdentification,proto3" json:"FlightIdentification,omitempty"`
}

func (x *SearchOneWayFlightResponse) Reset() {
	*x = SearchOneWayFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOneWayFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOneWayFlightResponse) ProtoMessage() {}

func (x *SearchOneWayFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOneWayFlightResponse.ProtoReflect.Descriptor instead.
func (*SearchOneWayFlightResponse) Descriptor() ([]byte, []int) {
	return file_search_flight_proto_rawDescGZIP(), []int{1}
}

func (x *SearchOneWayFlightResponse) GetDepartureTime() string {
	if x != nil {
		return x.DepartureTime
	}
	return ""
}

func (x *SearchOneWayFlightResponse) GetArrivalTime() string {
	if x != nil {
		return x.ArrivalTime
	}
	return ""
}

func (x *SearchOneWayFlightResponse) GetFlightIdentification() string {
	if x != nil {
		return x.FlightIdentification
	}
	return ""
}

var File_search_flight_proto protoreflect.FileDescriptor

var file_search_flight_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x6e,
	0x65, 0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x60,
	0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x6e, 0x65, 0x57, 0x61, 0x79, 0x12, 0x1a, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x6e,
	0x65, 0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_flight_proto_rawDescOnce sync.Once
	file_search_flight_proto_rawDescData = file_search_flight_proto_rawDesc
)

func file_search_flight_proto_rawDescGZIP() []byte {
	file_search_flight_proto_rawDescOnce.Do(func() {
		file_search_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_flight_proto_rawDescData)
	})
	return file_search_flight_proto_rawDescData
}

var file_search_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_search_flight_proto_goTypes = []interface{}{
	(*SearchOneWayFlightRequest)(nil),  // 0: SearchOneWayFlightRequest
	(*SearchOneWayFlightResponse)(nil), // 1: SearchOneWayFlightResponse
}
var file_search_flight_proto_depIdxs = []int32{
	0, // 0: SearchFlightService.SearchOneWay:input_type -> SearchOneWayFlightRequest
	1, // 1: SearchFlightService.SearchOneWay:output_type -> SearchOneWayFlightResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_search_flight_proto_init() }
func file_search_flight_proto_init() {
	if File_search_flight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOneWayFlightRequest); i {
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
		file_search_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOneWayFlightResponse); i {
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
			RawDescriptor: file_search_flight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_flight_proto_goTypes,
		DependencyIndexes: file_search_flight_proto_depIdxs,
		MessageInfos:      file_search_flight_proto_msgTypes,
	}.Build()
	File_search_flight_proto = out.File
	file_search_flight_proto_rawDesc = nil
	file_search_flight_proto_goTypes = nil
	file_search_flight_proto_depIdxs = nil
}