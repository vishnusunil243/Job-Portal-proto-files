// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: email.proto

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

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	SendOTP(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
	AddNotification(ctx context.Context, in *AddNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllNotifications(ctx context.Context, in *GetNotificationsByUserId, opts ...grpc.CallOption) (EmailService_GetAllNotificationsClient, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) SendOTP(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.EmailService/SendOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, "/user.EmailService/VerifyOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) AddNotification(ctx context.Context, in *AddNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.EmailService/AddNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetAllNotifications(ctx context.Context, in *GetNotificationsByUserId, opts ...grpc.CallOption) (EmailService_GetAllNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmailService_ServiceDesc.Streams[0], "/user.EmailService/GetAllNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &emailServiceGetAllNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmailService_GetAllNotificationsClient interface {
	Recv() (*NotificationResponse, error)
	grpc.ClientStream
}

type emailServiceGetAllNotificationsClient struct {
	grpc.ClientStream
}

func (x *emailServiceGetAllNotificationsClient) Recv() (*NotificationResponse, error) {
	m := new(NotificationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	SendOTP(context.Context, *SendOtpRequest) (*emptypb.Empty, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	AddNotification(context.Context, *AddNotificationRequest) (*emptypb.Empty, error)
	GetAllNotifications(*GetNotificationsByUserId, EmailService_GetAllNotificationsServer) error
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) SendOTP(context.Context, *SendOtpRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTP not implemented")
}
func (UnimplementedEmailServiceServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedEmailServiceServer) AddNotification(context.Context, *AddNotificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNotification not implemented")
}
func (UnimplementedEmailServiceServer) GetAllNotifications(*GetNotificationsByUserId, EmailService_GetAllNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllNotifications not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_SendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.EmailService/SendOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendOTP(ctx, req.(*SendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.EmailService/VerifyOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_AddNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).AddNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.EmailService/AddNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).AddNotification(ctx, req.(*AddNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetAllNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNotificationsByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmailServiceServer).GetAllNotifications(m, &emailServiceGetAllNotificationsServer{stream})
}

type EmailService_GetAllNotificationsServer interface {
	Send(*NotificationResponse) error
	grpc.ServerStream
}

type emailServiceGetAllNotificationsServer struct {
	grpc.ServerStream
}

func (x *emailServiceGetAllNotificationsServer) Send(m *NotificationResponse) error {
	return x.ServerStream.SendMsg(m)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOTP",
			Handler:    _EmailService_SendOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _EmailService_VerifyOTP_Handler,
		},
		{
			MethodName: "AddNotification",
			Handler:    _EmailService_AddNotification_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllNotifications",
			Handler:       _EmailService_GetAllNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "email.proto",
}
