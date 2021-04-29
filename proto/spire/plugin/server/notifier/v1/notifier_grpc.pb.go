// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package notifierv1

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

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifierClient interface {
	// Notify notifies the plugin that an event occurred. Errors returned by
	// the plugin are logged but otherwise ignored.
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	// NotifyAndAdvise notifies the plugin that an event occurred and waits
	// for a response. Errors returned by the plugin control SPIRE Server
	// behavior. See NotifyAndAdviseRequest for per-event details.
	NotifyAndAdvise(ctx context.Context, in *NotifyAndAdviseRequest, opts ...grpc.CallOption) (*NotifyAndAdviseResponse, error)
}

type notifierClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifierClient(cc grpc.ClientConnInterface) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/spire.plugin.server.notifier.v1.Notifier/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) NotifyAndAdvise(ctx context.Context, in *NotifyAndAdviseRequest, opts ...grpc.CallOption) (*NotifyAndAdviseResponse, error) {
	out := new(NotifyAndAdviseResponse)
	err := c.cc.Invoke(ctx, "/spire.plugin.server.notifier.v1.Notifier/NotifyAndAdvise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
// All implementations must embed UnimplementedNotifierServer
// for forward compatibility
type NotifierServer interface {
	// Notify notifies the plugin that an event occurred. Errors returned by
	// the plugin are logged but otherwise ignored.
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	// NotifyAndAdvise notifies the plugin that an event occurred and waits
	// for a response. Errors returned by the plugin control SPIRE Server
	// behavior. See NotifyAndAdviseRequest for per-event details.
	NotifyAndAdvise(context.Context, *NotifyAndAdviseRequest) (*NotifyAndAdviseResponse, error)
	mustEmbedUnimplementedNotifierServer()
}

// UnimplementedNotifierServer must be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (UnimplementedNotifierServer) Notify(context.Context, *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedNotifierServer) NotifyAndAdvise(context.Context, *NotifyAndAdviseRequest) (*NotifyAndAdviseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAndAdvise not implemented")
}
func (UnimplementedNotifierServer) mustEmbedUnimplementedNotifierServer() {}

// UnsafeNotifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifierServer will
// result in compilation errors.
type UnsafeNotifierServer interface {
	mustEmbedUnimplementedNotifierServer()
}

func RegisterNotifierServer(s grpc.ServiceRegistrar, srv NotifierServer) {
	s.RegisterService(&Notifier_ServiceDesc, srv)
}

func _Notifier_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.plugin.server.notifier.v1.Notifier/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_NotifyAndAdvise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyAndAdviseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NotifyAndAdvise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.plugin.server.notifier.v1.Notifier/NotifyAndAdvise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NotifyAndAdvise(ctx, req.(*NotifyAndAdviseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifier_ServiceDesc is the grpc.ServiceDesc for Notifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spire.plugin.server.notifier.v1.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Notifier_Notify_Handler,
		},
		{
			MethodName: "NotifyAndAdvise",
			Handler:    _Notifier_NotifyAndAdvise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/plugin/server/notifier/v1/notifier.proto",
}