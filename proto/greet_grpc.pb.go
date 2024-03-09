// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/greet.proto

package proto

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

const (
	Service_RequestResponse_FullMethodName        = "/my_service.Service/RequestResponse"
	Service_ServerSideStreaming_FullMethodName    = "/my_service.Service/ServerSideStreaming"
	Service_ClientSideStreaming_FullMethodName    = "/my_service.Service/ClientSideStreaming"
	Service_BidirectionalStreaming_FullMethodName = "/my_service.Service/BidirectionalStreaming"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	RequestResponse(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	ServerSideStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (Service_ServerSideStreamingClient, error)
	ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (Service_ClientSideStreamingClient, error)
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (Service_BidirectionalStreamingClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) RequestResponse(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Service_RequestResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ServerSideStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (Service_ServerSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], Service_ServerSideStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServerSideStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_ServerSideStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type serviceServerSideStreamingClient struct {
	grpc.ClientStream
}

func (x *serviceServerSideStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (Service_ClientSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[1], Service_ClientSideStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceClientSideStreamingClient{stream}
	return x, nil
}

type Service_ClientSideStreamingClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type serviceClientSideStreamingClient struct {
	grpc.ClientStream
}

func (x *serviceClientSideStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceClientSideStreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (Service_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[2], Service_BidirectionalStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceBidirectionalStreamingClient{stream}
	return x, nil
}

type Service_BidirectionalStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type serviceBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *serviceBidirectionalStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceBidirectionalStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	RequestResponse(context.Context, *NoParam) (*HelloResponse, error)
	ServerSideStreaming(*NamesList, Service_ServerSideStreamingServer) error
	ClientSideStreaming(Service_ClientSideStreamingServer) error
	BidirectionalStreaming(Service_BidirectionalStreamingServer) error
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) RequestResponse(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestResponse not implemented")
}
func (UnimplementedServiceServer) ServerSideStreaming(*NamesList, Service_ServerSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreaming not implemented")
}
func (UnimplementedServiceServer) ClientSideStreaming(Service_ClientSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreaming not implemented")
}
func (UnimplementedServiceServer) BidirectionalStreaming(Service_BidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreaming not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_RequestResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RequestResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RequestResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RequestResponse(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ServerSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).ServerSideStreaming(m, &serviceServerSideStreamingServer{stream})
}

type Service_ServerSideStreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type serviceServerSideStreamingServer struct {
	grpc.ServerStream
}

func (x *serviceServerSideStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_ClientSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).ClientSideStreaming(&serviceClientSideStreamingServer{stream})
}

type Service_ClientSideStreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type serviceClientSideStreamingServer struct {
	grpc.ServerStream
}

func (x *serviceClientSideStreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceClientSideStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).BidirectionalStreaming(&serviceBidirectionalStreamingServer{stream})
}

type Service_BidirectionalStreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type serviceBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *serviceBidirectionalStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceBidirectionalStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "my_service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestResponse",
			Handler:    _Service_RequestResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreaming",
			Handler:       _Service_ServerSideStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStreaming",
			Handler:       _Service_ClientSideStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _Service_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}