// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/contract.proto

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

// ProcessServiceClient is the client API for ProcessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessServiceClient interface {
	CallUnary(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Process, error)
	CallSStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ProcessService_CallSStreamClient, error)
	CallCStream(ctx context.Context, opts ...grpc.CallOption) (ProcessService_CallCStreamClient, error)
	CallBStream(ctx context.Context, opts ...grpc.CallOption) (ProcessService_CallBStreamClient, error)
}

type processServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessServiceClient(cc grpc.ClientConnInterface) ProcessServiceClient {
	return &processServiceClient{cc}
}

func (c *processServiceClient) CallUnary(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/proto.ProcessService/CallUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processServiceClient) CallSStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ProcessService_CallSStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessService_ServiceDesc.Streams[0], "/proto.ProcessService/CallSStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &processServiceCallSStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessService_CallSStreamClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processServiceCallSStreamClient struct {
	grpc.ClientStream
}

func (x *processServiceCallSStreamClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processServiceClient) CallCStream(ctx context.Context, opts ...grpc.CallOption) (ProcessService_CallCStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessService_ServiceDesc.Streams[1], "/proto.ProcessService/CallCStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &processServiceCallCStreamClient{stream}
	return x, nil
}

type ProcessService_CallCStreamClient interface {
	Send(*Process) error
	CloseAndRecv() (*ListProcess, error)
	grpc.ClientStream
}

type processServiceCallCStreamClient struct {
	grpc.ClientStream
}

func (x *processServiceCallCStreamClient) Send(m *Process) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processServiceCallCStreamClient) CloseAndRecv() (*ListProcess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ListProcess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processServiceClient) CallBStream(ctx context.Context, opts ...grpc.CallOption) (ProcessService_CallBStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessService_ServiceDesc.Streams[2], "/proto.ProcessService/CallBStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &processServiceCallBStreamClient{stream}
	return x, nil
}

type ProcessService_CallBStreamClient interface {
	Send(*Process) error
	Recv() (*Process, error)
	grpc.ClientStream
}

type processServiceCallBStreamClient struct {
	grpc.ClientStream
}

func (x *processServiceCallBStreamClient) Send(m *Process) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processServiceCallBStreamClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessServiceServer is the server API for ProcessService service.
// All implementations must embed UnimplementedProcessServiceServer
// for forward compatibility
type ProcessServiceServer interface {
	CallUnary(context.Context, *Empty) (*Process, error)
	CallSStream(*Empty, ProcessService_CallSStreamServer) error
	CallCStream(ProcessService_CallCStreamServer) error
	CallBStream(ProcessService_CallBStreamServer) error
	mustEmbedUnimplementedProcessServiceServer()
}

// UnimplementedProcessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessServiceServer struct {
}

func (UnimplementedProcessServiceServer) CallUnary(context.Context, *Empty) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallUnary not implemented")
}
func (UnimplementedProcessServiceServer) CallSStream(*Empty, ProcessService_CallSStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CallSStream not implemented")
}
func (UnimplementedProcessServiceServer) CallCStream(ProcessService_CallCStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CallCStream not implemented")
}
func (UnimplementedProcessServiceServer) CallBStream(ProcessService_CallBStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CallBStream not implemented")
}
func (UnimplementedProcessServiceServer) mustEmbedUnimplementedProcessServiceServer() {}

// UnsafeProcessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessServiceServer will
// result in compilation errors.
type UnsafeProcessServiceServer interface {
	mustEmbedUnimplementedProcessServiceServer()
}

func RegisterProcessServiceServer(s grpc.ServiceRegistrar, srv ProcessServiceServer) {
	s.RegisterService(&ProcessService_ServiceDesc, srv)
}

func _ProcessService_CallUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).CallUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProcessService/CallUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).CallUnary(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessService_CallSStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessServiceServer).CallSStream(m, &processServiceCallSStreamServer{stream})
}

type ProcessService_CallSStreamServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processServiceCallSStreamServer struct {
	grpc.ServerStream
}

func (x *processServiceCallSStreamServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessService_CallCStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessServiceServer).CallCStream(&processServiceCallCStreamServer{stream})
}

type ProcessService_CallCStreamServer interface {
	SendAndClose(*ListProcess) error
	Recv() (*Process, error)
	grpc.ServerStream
}

type processServiceCallCStreamServer struct {
	grpc.ServerStream
}

func (x *processServiceCallCStreamServer) SendAndClose(m *ListProcess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processServiceCallCStreamServer) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProcessService_CallBStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessServiceServer).CallBStream(&processServiceCallBStreamServer{stream})
}

type ProcessService_CallBStreamServer interface {
	Send(*Process) error
	Recv() (*Process, error)
	grpc.ServerStream
}

type processServiceCallBStreamServer struct {
	grpc.ServerStream
}

func (x *processServiceCallBStreamServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processServiceCallBStreamServer) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessService_ServiceDesc is the grpc.ServiceDesc for ProcessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProcessService",
	HandlerType: (*ProcessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallUnary",
			Handler:    _ProcessService_CallUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CallSStream",
			Handler:       _ProcessService_CallSStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CallCStream",
			Handler:       _ProcessService_CallCStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CallBStream",
			Handler:       _ProcessService_CallBStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/contract.proto",
}