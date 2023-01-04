// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	// UpdateSigningStatus is a bidirectional connection where server is requesting that the client:
	// - Sign a release artifact.
	// - Provide an update on a previous request to sign a release artifact.
	// - Consider a previous request to sign a release artifact as obsolete.
	// The client initiates a connection with the server and waits for the server to issue a request
	// such as:
	// - An update on the status of a signing request (either running or completed).
	// - An acknowledgement that a request to sign a release artifact has been accepted and initiated.
	UpdateSigningStatus(ctx context.Context, opts ...grpc.CallOption) (ReleaseService_UpdateSigningStatusClient, error)
}

type releaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseServiceClient(cc grpc.ClientConnInterface) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) UpdateSigningStatus(ctx context.Context, opts ...grpc.CallOption) (ReleaseService_UpdateSigningStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReleaseService_ServiceDesc.Streams[0], "/protos.ReleaseService/UpdateSigningStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &releaseServiceUpdateSigningStatusClient{stream}
	return x, nil
}

type ReleaseService_UpdateSigningStatusClient interface {
	Send(*SigningStatus) error
	Recv() (*SigningRequest, error)
	grpc.ClientStream
}

type releaseServiceUpdateSigningStatusClient struct {
	grpc.ClientStream
}

func (x *releaseServiceUpdateSigningStatusClient) Send(m *SigningStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *releaseServiceUpdateSigningStatusClient) Recv() (*SigningRequest, error) {
	m := new(SigningRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
// All implementations must embed UnimplementedReleaseServiceServer
// for forward compatibility
type ReleaseServiceServer interface {
	// UpdateSigningStatus is a bidirectional connection where server is requesting that the client:
	// - Sign a release artifact.
	// - Provide an update on a previous request to sign a release artifact.
	// - Consider a previous request to sign a release artifact as obsolete.
	// The client initiates a connection with the server and waits for the server to issue a request
	// such as:
	// - An update on the status of a signing request (either running or completed).
	// - An acknowledgement that a request to sign a release artifact has been accepted and initiated.
	UpdateSigningStatus(ReleaseService_UpdateSigningStatusServer) error
	mustEmbedUnimplementedReleaseServiceServer()
}

// UnimplementedReleaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReleaseServiceServer struct {
}

func (UnimplementedReleaseServiceServer) UpdateSigningStatus(ReleaseService_UpdateSigningStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateSigningStatus not implemented")
}
func (UnimplementedReleaseServiceServer) mustEmbedUnimplementedReleaseServiceServer() {}

// UnsafeReleaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReleaseServiceServer will
// result in compilation errors.
type UnsafeReleaseServiceServer interface {
	mustEmbedUnimplementedReleaseServiceServer()
}

func RegisterReleaseServiceServer(s grpc.ServiceRegistrar, srv ReleaseServiceServer) {
	s.RegisterService(&ReleaseService_ServiceDesc, srv)
}

func _ReleaseService_UpdateSigningStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReleaseServiceServer).UpdateSigningStatus(&releaseServiceUpdateSigningStatusServer{stream})
}

type ReleaseService_UpdateSigningStatusServer interface {
	Send(*SigningRequest) error
	Recv() (*SigningStatus, error)
	grpc.ServerStream
}

type releaseServiceUpdateSigningStatusServer struct {
	grpc.ServerStream
}

func (x *releaseServiceUpdateSigningStatusServer) Send(m *SigningRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *releaseServiceUpdateSigningStatusServer) Recv() (*SigningStatus, error) {
	m := new(SigningStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReleaseService_ServiceDesc is the grpc.ServiceDesc for ReleaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReleaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateSigningStatus",
			Handler:       _ReleaseService_UpdateSigningStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "relui.proto",
}