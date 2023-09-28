// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: search_flight.proto

package server_proto

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

// SearchFlightServiceClient is the client API for SearchFlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchFlightServiceClient interface {
	SearchOneWay(ctx context.Context, in *SearchOneWayFlightRequest, opts ...grpc.CallOption) (*SearchOneWayFlightResponse, error)
}

type searchFlightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchFlightServiceClient(cc grpc.ClientConnInterface) SearchFlightServiceClient {
	return &searchFlightServiceClient{cc}
}

func (c *searchFlightServiceClient) SearchOneWay(ctx context.Context, in *SearchOneWayFlightRequest, opts ...grpc.CallOption) (*SearchOneWayFlightResponse, error) {
	out := new(SearchOneWayFlightResponse)
	err := c.cc.Invoke(ctx, "/SearchFlightService/SearchOneWay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchFlightServiceServer is the server API for SearchFlightService service.
// All implementations must embed UnimplementedSearchFlightServiceServer
// for forward compatibility
type SearchFlightServiceServer interface {
	SearchOneWay(context.Context, *SearchOneWayFlightRequest) (*SearchOneWayFlightResponse, error)
	mustEmbedUnimplementedSearchFlightServiceServer()
}

// UnimplementedSearchFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchFlightServiceServer struct {
}

func (UnimplementedSearchFlightServiceServer) SearchOneWay(context.Context, *SearchOneWayFlightRequest) (*SearchOneWayFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOneWay not implemented")
}
func (UnimplementedSearchFlightServiceServer) mustEmbedUnimplementedSearchFlightServiceServer() {}

// UnsafeSearchFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchFlightServiceServer will
// result in compilation errors.
type UnsafeSearchFlightServiceServer interface {
	mustEmbedUnimplementedSearchFlightServiceServer()
}

func RegisterSearchFlightServiceServer(s grpc.ServiceRegistrar, srv SearchFlightServiceServer) {
	s.RegisterService(&SearchFlightService_ServiceDesc, srv)
}

func _SearchFlightService_SearchOneWay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOneWayFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchFlightServiceServer).SearchOneWay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchFlightService/SearchOneWay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchFlightServiceServer).SearchOneWay(ctx, req.(*SearchOneWayFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchFlightService_ServiceDesc is the grpc.ServiceDesc for SearchFlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchFlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SearchFlightService",
	HandlerType: (*SearchFlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchOneWay",
			Handler:    _SearchFlightService_SearchOneWay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search_flight.proto",
}