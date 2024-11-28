// proto/fire_detection.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/fire_detection.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FireDetectionService_StreamVideo_FullMethodName = "/firedetection.FireDetectionService/StreamVideo"
)

// FireDetectionServiceClient is the client API for FireDetectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FireDetectionServiceClient interface {
	StreamVideo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[VideoChunk, VideoResponse], error)
}

type fireDetectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFireDetectionServiceClient(cc grpc.ClientConnInterface) FireDetectionServiceClient {
	return &fireDetectionServiceClient{cc}
}

func (c *fireDetectionServiceClient) StreamVideo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[VideoChunk, VideoResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FireDetectionService_ServiceDesc.Streams[0], FireDetectionService_StreamVideo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[VideoChunk, VideoResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FireDetectionService_StreamVideoClient = grpc.BidiStreamingClient[VideoChunk, VideoResponse]

// FireDetectionServiceServer is the server API for FireDetectionService service.
// All implementations must embed UnimplementedFireDetectionServiceServer
// for forward compatibility.
type FireDetectionServiceServer interface {
	StreamVideo(grpc.BidiStreamingServer[VideoChunk, VideoResponse]) error
	mustEmbedUnimplementedFireDetectionServiceServer()
}

// UnimplementedFireDetectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFireDetectionServiceServer struct{}

func (UnimplementedFireDetectionServiceServer) StreamVideo(grpc.BidiStreamingServer[VideoChunk, VideoResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamVideo not implemented")
}
func (UnimplementedFireDetectionServiceServer) mustEmbedUnimplementedFireDetectionServiceServer() {}
func (UnimplementedFireDetectionServiceServer) testEmbeddedByValue()                              {}

// UnsafeFireDetectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FireDetectionServiceServer will
// result in compilation errors.
type UnsafeFireDetectionServiceServer interface {
	mustEmbedUnimplementedFireDetectionServiceServer()
}

func RegisterFireDetectionServiceServer(s grpc.ServiceRegistrar, srv FireDetectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedFireDetectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FireDetectionService_ServiceDesc, srv)
}

func _FireDetectionService_StreamVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FireDetectionServiceServer).StreamVideo(&grpc.GenericServerStream[VideoChunk, VideoResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FireDetectionService_StreamVideoServer = grpc.BidiStreamingServer[VideoChunk, VideoResponse]

// FireDetectionService_ServiceDesc is the grpc.ServiceDesc for FireDetectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FireDetectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "firedetection.FireDetectionService",
	HandlerType: (*FireDetectionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamVideo",
			Handler:       _FireDetectionService_StreamVideo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/fire_detection.proto",
}