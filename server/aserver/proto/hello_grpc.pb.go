//go get -u google.golang.org/protobuf/cmd/protoc-gen-go && go install google.golang.org/protobuf/cmd/protoc-gen-go && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

// export PATH="$PATH:$(go env GOPATH)/bin"
// protoc --go_out=./proto --go-grpc_out=./proto ./proto/hello.proto

// npm install grpc_tools_node_protoc_ts --save-dev
// npm install @grpc/grpc-js
// protoc --plugin=protoc-gen-ts=../node_modules/.bin/protoc-gen-ts --ts_out=grpc_js:. hello.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.3
// source: proto/hello.proto

package service

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
	ServerStatusService_ServerStatus_FullMethodName = "/ServerStatusService/ServerStatus"
)

// ServerStatusServiceClient is the client API for ServerStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStatusServiceClient interface {
	ServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (ServerStatusService_ServerStatusClient, error)
}

type serverStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStatusServiceClient(cc grpc.ClientConnInterface) ServerStatusServiceClient {
	return &serverStatusServiceClient{cc}
}

func (c *serverStatusServiceClient) ServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (ServerStatusService_ServerStatusClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ServerStatusService_ServiceDesc.Streams[0], ServerStatusService_ServerStatus_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &serverStatusServiceServerStatusClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStatusService_ServerStatusClient interface {
	Recv() (*ServerStatusResponse, error)
	grpc.ClientStream
}

type serverStatusServiceServerStatusClient struct {
	grpc.ClientStream
}

func (x *serverStatusServiceServerStatusClient) Recv() (*ServerStatusResponse, error) {
	m := new(ServerStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerStatusServiceServer is the server API for ServerStatusService service.
// All implementations must embed UnimplementedServerStatusServiceServer
// for forward compatibility
type ServerStatusServiceServer interface {
	ServerStatus(*ServerStatusRequest, ServerStatusService_ServerStatusServer) error
	mustEmbedUnimplementedServerStatusServiceServer()
}

// UnimplementedServerStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerStatusServiceServer struct {
}

func (UnimplementedServerStatusServiceServer) ServerStatus(*ServerStatusRequest, ServerStatusService_ServerStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStatus not implemented")
}
func (UnimplementedServerStatusServiceServer) mustEmbedUnimplementedServerStatusServiceServer() {}

// UnsafeServerStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStatusServiceServer will
// result in compilation errors.
type UnsafeServerStatusServiceServer interface {
	mustEmbedUnimplementedServerStatusServiceServer()
}

func RegisterServerStatusServiceServer(s grpc.ServiceRegistrar, srv ServerStatusServiceServer) {
	s.RegisterService(&ServerStatusService_ServiceDesc, srv)
}

func _ServerStatusService_ServerStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStatusServiceServer).ServerStatus(m, &serverStatusServiceServerStatusServer{ServerStream: stream})
}

type ServerStatusService_ServerStatusServer interface {
	Send(*ServerStatusResponse) error
	grpc.ServerStream
}

type serverStatusServiceServerStatusServer struct {
	grpc.ServerStream
}

func (x *serverStatusServiceServerStatusServer) Send(m *ServerStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServerStatusService_ServiceDesc is the grpc.ServiceDesc for ServerStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServerStatusService",
	HandlerType: (*ServerStatusServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStatus",
			Handler:       _ServerStatusService_ServerStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/hello.proto",
}
