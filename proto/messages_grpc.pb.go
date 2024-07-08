// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/messages.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	RemoteService_Receive_FullMethodName = "/proto.RemoteService/Receive"
	RemoteService_Command_FullMethodName = "/proto.RemoteService/Command"
)

// RemoteServiceClient is the client API for RemoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteServiceClient interface {
	Receive(ctx context.Context, opts ...grpc.CallOption) (RemoteService_ReceiveClient, error)
	Command(ctx context.Context, in *RemoteCommand, opts ...grpc.CallOption) (*CommandResponse, error)
}

type remoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteServiceClient(cc grpc.ClientConnInterface) RemoteServiceClient {
	return &remoteServiceClient{cc}
}

func (c *remoteServiceClient) Receive(ctx context.Context, opts ...grpc.CallOption) (RemoteService_ReceiveClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RemoteService_ServiceDesc.Streams[0], RemoteService_Receive_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &remoteServiceReceiveClient{ClientStream: stream}
	return x, nil
}

type RemoteService_ReceiveClient interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type remoteServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *remoteServiceReceiveClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *remoteServiceReceiveClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *remoteServiceClient) Command(ctx context.Context, in *RemoteCommand, opts ...grpc.CallOption) (*CommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, RemoteService_Command_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteServiceServer is the server API for RemoteService service.
// All implementations must embed UnimplementedRemoteServiceServer
// for forward compatibility
type RemoteServiceServer interface {
	Receive(RemoteService_ReceiveServer) error
	Command(context.Context, *RemoteCommand) (*CommandResponse, error)
	mustEmbedUnimplementedRemoteServiceServer()
}

// UnimplementedRemoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteServiceServer struct {
}

func (UnimplementedRemoteServiceServer) Receive(RemoteService_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedRemoteServiceServer) Command(context.Context, *RemoteCommand) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Command not implemented")
}
func (UnimplementedRemoteServiceServer) mustEmbedUnimplementedRemoteServiceServer() {}

// UnsafeRemoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteServiceServer will
// result in compilation errors.
type UnsafeRemoteServiceServer interface {
	mustEmbedUnimplementedRemoteServiceServer()
}

func RegisterRemoteServiceServer(s grpc.ServiceRegistrar, srv RemoteServiceServer) {
	s.RegisterService(&RemoteService_ServiceDesc, srv)
}

func _RemoteService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RemoteServiceServer).Receive(&remoteServiceReceiveServer{ServerStream: stream})
}

type RemoteService_ReceiveServer interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type remoteServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *remoteServiceReceiveServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *remoteServiceReceiveServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RemoteService_Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteService_Command_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).Command(ctx, req.(*RemoteCommand))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteService_ServiceDesc is the grpc.ServiceDesc for RemoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RemoteService",
	HandlerType: (*RemoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Command",
			Handler:    _RemoteService_Command_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _RemoteService_Receive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/messages.proto",
}
