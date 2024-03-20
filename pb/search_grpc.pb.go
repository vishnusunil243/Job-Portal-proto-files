// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: search.proto

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	AddSearchHistory(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSearchHistory(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*SearchResponse, error)
	UserAddReview(ctx context.Context, in *UserReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCompanyReview(ctx context.Context, in *ReviewByCompanyId, opts ...grpc.CallOption) (SearchService_GetCompanyReviewClient, error)
	RemoveReview(ctx context.Context, in *UserReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) AddSearchHistory(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.SearchService/AddSearchHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetSearchHistory(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/user.SearchService/GetSearchHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) UserAddReview(ctx context.Context, in *UserReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.SearchService/UserAddReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetCompanyReview(ctx context.Context, in *ReviewByCompanyId, opts ...grpc.CallOption) (SearchService_GetCompanyReviewClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[0], "/user.SearchService/GetCompanyReview", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceGetCompanyReviewClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchService_GetCompanyReviewClient interface {
	Recv() (*ReviewResponse, error)
	grpc.ClientStream
}

type searchServiceGetCompanyReviewClient struct {
	grpc.ClientStream
}

func (x *searchServiceGetCompanyReviewClient) Recv() (*ReviewResponse, error) {
	m := new(ReviewResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchServiceClient) RemoveReview(ctx context.Context, in *UserReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.SearchService/RemoveReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	AddSearchHistory(context.Context, *SearchRequest) (*emptypb.Empty, error)
	GetSearchHistory(context.Context, *UserId) (*SearchResponse, error)
	UserAddReview(context.Context, *UserReviewRequest) (*emptypb.Empty, error)
	GetCompanyReview(*ReviewByCompanyId, SearchService_GetCompanyReviewServer) error
	RemoveReview(context.Context, *UserReviewRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) AddSearchHistory(context.Context, *SearchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSearchHistory not implemented")
}
func (UnimplementedSearchServiceServer) GetSearchHistory(context.Context, *UserId) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchHistory not implemented")
}
func (UnimplementedSearchServiceServer) UserAddReview(context.Context, *UserReviewRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAddReview not implemented")
}
func (UnimplementedSearchServiceServer) GetCompanyReview(*ReviewByCompanyId, SearchService_GetCompanyReviewServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCompanyReview not implemented")
}
func (UnimplementedSearchServiceServer) RemoveReview(context.Context, *UserReviewRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReview not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_AddSearchHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).AddSearchHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.SearchService/AddSearchHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).AddSearchHistory(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetSearchHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetSearchHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.SearchService/GetSearchHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetSearchHistory(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_UserAddReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).UserAddReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.SearchService/UserAddReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).UserAddReview(ctx, req.(*UserReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetCompanyReview_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReviewByCompanyId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServiceServer).GetCompanyReview(m, &searchServiceGetCompanyReviewServer{stream})
}

type SearchService_GetCompanyReviewServer interface {
	Send(*ReviewResponse) error
	grpc.ServerStream
}

type searchServiceGetCompanyReviewServer struct {
	grpc.ServerStream
}

func (x *searchServiceGetCompanyReviewServer) Send(m *ReviewResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SearchService_RemoveReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).RemoveReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.SearchService/RemoveReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).RemoveReview(ctx, req.(*UserReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSearchHistory",
			Handler:    _SearchService_AddSearchHistory_Handler,
		},
		{
			MethodName: "GetSearchHistory",
			Handler:    _SearchService_GetSearchHistory_Handler,
		},
		{
			MethodName: "UserAddReview",
			Handler:    _SearchService_UserAddReview_Handler,
		},
		{
			MethodName: "RemoveReview",
			Handler:    _SearchService_RemoveReview_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCompanyReview",
			Handler:       _SearchService_GetCompanyReview_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "search.proto",
}
