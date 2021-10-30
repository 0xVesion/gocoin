// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	StreamTasks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Node_StreamTasksClient, error)
	SubmitTask(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) StreamTasks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Node_StreamTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], "/gocoin.Node/StreamTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeStreamTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_StreamTasksClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type nodeStreamTasksClient struct {
	grpc.ClientStream
}

func (x *nodeStreamTasksClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) SubmitTask(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gocoin.Node/SubmitTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	StreamTasks(*emptypb.Empty, Node_StreamTasksServer) error
	SubmitTask(context.Context, *Submission) (*emptypb.Empty, error)
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) StreamTasks(*emptypb.Empty, Node_StreamTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTasks not implemented")
}
func (UnimplementedNodeServer) SubmitTask(context.Context, *Submission) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTask not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_StreamTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).StreamTasks(m, &nodeStreamTasksServer{stream})
}

type Node_StreamTasksServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type nodeStreamTasksServer struct {
	grpc.ServerStream
}

func (x *nodeStreamTasksServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_SubmitTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Submission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SubmitTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gocoin.Node/SubmitTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SubmitTask(ctx, req.(*Submission))
	}
	return interceptor(ctx, in, info, handler)
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gocoin.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTask",
			Handler:    _Node_SubmitTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTasks",
			Handler:       _Node_StreamTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/gocoin.proto",
}