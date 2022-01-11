// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/command.proto

package command

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

// RunRPCCommandsClient is the client API for RunRPCCommands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunRPCCommandsClient interface {
	RunShellCommands(ctx context.Context, in *ShellCommandsOutputRequest, opts ...grpc.CallOption) (*ShellCommandsOutputResponse, error)
}

type runRPCCommandsClient struct {
	cc grpc.ClientConnInterface
}

func NewRunRPCCommandsClient(cc grpc.ClientConnInterface) RunRPCCommandsClient {
	return &runRPCCommandsClient{cc}
}

func (c *runRPCCommandsClient) RunShellCommands(ctx context.Context, in *ShellCommandsOutputRequest, opts ...grpc.CallOption) (*ShellCommandsOutputResponse, error) {
	out := new(ShellCommandsOutputResponse)
	err := c.cc.Invoke(ctx, "/command.RunRPCCommands/RunShellCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunRPCCommandsServer is the server API for RunRPCCommands service.
// All implementations must embed UnimplementedRunRPCCommandsServer
// for forward compatibility
type RunRPCCommandsServer interface {
	RunShellCommands(context.Context, *ShellCommandsOutputRequest) (*ShellCommandsOutputResponse, error)
	mustEmbedUnimplementedRunRPCCommandsServer()
}

// UnimplementedRunRPCCommandsServer must be embedded to have forward compatible implementations.
type UnimplementedRunRPCCommandsServer struct {
}

func (UnimplementedRunRPCCommandsServer) RunShellCommands(context.Context, *ShellCommandsOutputRequest) (*ShellCommandsOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunShellCommands not implemented")
}
func (UnimplementedRunRPCCommandsServer) mustEmbedUnimplementedRunRPCCommandsServer() {}

// UnsafeRunRPCCommandsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunRPCCommandsServer will
// result in compilation errors.
type UnsafeRunRPCCommandsServer interface {
	mustEmbedUnimplementedRunRPCCommandsServer()
}

func RegisterRunRPCCommandsServer(s grpc.ServiceRegistrar, srv RunRPCCommandsServer) {
	s.RegisterService(&RunRPCCommands_ServiceDesc, srv)
}

func _RunRPCCommands_RunShellCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellCommandsOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunRPCCommandsServer).RunShellCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.RunRPCCommands/RunShellCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunRPCCommandsServer).RunShellCommands(ctx, req.(*ShellCommandsOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RunRPCCommands_ServiceDesc is the grpc.ServiceDesc for RunRPCCommands service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RunRPCCommands_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "command.RunRPCCommands",
	HandlerType: (*RunRPCCommandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunShellCommands",
			Handler:    _RunRPCCommands_RunShellCommands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/command.proto",
}
