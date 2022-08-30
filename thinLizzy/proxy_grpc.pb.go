// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proxy.proto

package thinLizzy

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

// CallEndpointClient is the client API for CallEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallEndpointClient interface {
	JobSubmit(ctx context.Context, in *SubmitBody, opts ...grpc.CallOption) (*Response, error)
}

type callEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewCallEndpointClient(cc grpc.ClientConnInterface) CallEndpointClient {
	return &callEndpointClient{cc}
}

func (c *callEndpointClient) JobSubmit(ctx context.Context, in *SubmitBody, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CallEndpoint/JobSubmit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallEndpointServer is the server API for CallEndpoint service.
// All implementations must embed UnimplementedCallEndpointServer
// for forward compatibility
type CallEndpointServer interface {
	JobSubmit(context.Context, *SubmitBody) (*Response, error)
	// mustEmbedUnimplementedCallEndpointServer()
}

// UnimplementedCallEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedCallEndpointServer struct {
}

func (UnimplementedCallEndpointServer) JobSubmit(context.Context, *SubmitBody) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSubmit not implemented")
}
func (UnimplementedCallEndpointServer) mustEmbedUnimplementedCallEndpointServer() {}

// UnsafeCallEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallEndpointServer will
// result in compilation errors.
type UnsafeCallEndpointServer interface {
	mustEmbedUnimplementedCallEndpointServer()
}

func RegisterCallEndpointServer(s grpc.ServiceRegistrar, srv CallEndpointServer) {
	s.RegisterService(&CallEndpoint_ServiceDesc, srv)
}

func _CallEndpoint_JobSubmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallEndpointServer).JobSubmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CallEndpoint/JobSubmit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallEndpointServer).JobSubmit(ctx, req.(*SubmitBody))
	}
	return interceptor(ctx, in, info, handler)
}

// CallEndpoint_ServiceDesc is the grpc.ServiceDesc for CallEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CallEndpoint",
	HandlerType: (*CallEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JobSubmit",
			Handler:    _CallEndpoint_JobSubmit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}
