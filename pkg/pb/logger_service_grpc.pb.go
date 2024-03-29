// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: logger_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoggerServiceClient is the client API for LoggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerServiceClient interface {
	SendLoginTimestampToLogger(ctx context.Context, in *LoginTimestamp, opts ...grpc.CallOption) (*LoginTimestamp, error)
	CheckRediness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type loggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerServiceClient(cc grpc.ClientConnInterface) LoggerServiceClient {
	return &loggerServiceClient{cc}
}

func (c *loggerServiceClient) SendLoginTimestampToLogger(ctx context.Context, in *LoginTimestamp, opts ...grpc.CallOption) (*LoginTimestamp, error) {
	out := new(LoginTimestamp)
	err := c.cc.Invoke(ctx, "/pb.LoggerService/SendLoginTimestampToLogger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerServiceClient) CheckRediness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LoggerService/CheckRediness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServiceServer is the server API for LoggerService service.
// All implementations must embed UnimplementedLoggerServiceServer
// for forward compatibility
type LoggerServiceServer interface {
	SendLoginTimestampToLogger(context.Context, *LoginTimestamp) (*LoginTimestamp, error)
	CheckRediness(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedLoggerServiceServer()
}

// UnimplementedLoggerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerServiceServer struct {
}

func (UnimplementedLoggerServiceServer) SendLoginTimestampToLogger(context.Context, *LoginTimestamp) (*LoginTimestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLoginTimestampToLogger not implemented")
}
func (UnimplementedLoggerServiceServer) CheckRediness(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRediness not implemented")
}
func (UnimplementedLoggerServiceServer) mustEmbedUnimplementedLoggerServiceServer() {}

// UnsafeLoggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServiceServer will
// result in compilation errors.
type UnsafeLoggerServiceServer interface {
	mustEmbedUnimplementedLoggerServiceServer()
}

func RegisterLoggerServiceServer(s grpc.ServiceRegistrar, srv LoggerServiceServer) {
	s.RegisterService(&LoggerService_ServiceDesc, srv)
}

func _LoggerService_SendLoginTimestampToLogger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginTimestamp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServiceServer).SendLoginTimestampToLogger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoggerService/SendLoginTimestampToLogger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServiceServer).SendLoginTimestampToLogger(ctx, req.(*LoginTimestamp))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggerService_CheckRediness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServiceServer).CheckRediness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoggerService/CheckRediness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServiceServer).CheckRediness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// LoggerService_ServiceDesc is the grpc.ServiceDesc for LoggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LoggerService",
	HandlerType: (*LoggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLoginTimestampToLogger",
			Handler:    _LoggerService_SendLoginTimestampToLogger_Handler,
		},
		{
			MethodName: "CheckRediness",
			Handler:    _LoggerService_CheckRediness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logger_service.proto",
}
