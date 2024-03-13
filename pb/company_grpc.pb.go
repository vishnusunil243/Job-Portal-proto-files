// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: company.proto

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

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	CompanySignup(ctx context.Context, in *CompanySignupRequest, opts ...grpc.CallOption) (*CompanySignupResponse, error)
	CompanyLogin(ctx context.Context, in *CompanyLoginRequest, opts ...grpc.CallOption) (*CompanySignupResponse, error)
	AddJobs(ctx context.Context, in *AddJobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	UpdateJobs(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteJob(ctx context.Context, in *GetJobById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllJobs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CompanyService_GetAllJobsClient, error)
	GetJob(ctx context.Context, in *GetJobById, opts ...grpc.CallOption) (*JobResponse, error)
	GetAllJobsForCompany(ctx context.Context, in *GetJobByCompanyId, opts ...grpc.CallOption) (CompanyService_GetAllJobsForCompanyClient, error)
	CompanyAddJobSkill(ctx context.Context, in *AddJobSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CompanyUpdateJobSkill(ctx context.Context, in *AddJobSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllJobSkill(ctx context.Context, in *GetJobById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteJobSkill(ctx context.Context, in *JobSkillId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CompanySignup(ctx context.Context, in *CompanySignupRequest, opts ...grpc.CallOption) (*CompanySignupResponse, error) {
	out := new(CompanySignupResponse)
	err := c.cc.Invoke(ctx, "/user.CompanyService/CompanySignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) CompanyLogin(ctx context.Context, in *CompanyLoginRequest, opts ...grpc.CallOption) (*CompanySignupResponse, error) {
	out := new(CompanySignupResponse)
	err := c.cc.Invoke(ctx, "/user.CompanyService/CompanyLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) AddJobs(ctx context.Context, in *AddJobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/user.CompanyService/AddJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateJobs(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.CompanyService/UpdateJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteJob(ctx context.Context, in *GetJobById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.CompanyService/DeleteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetAllJobs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CompanyService_GetAllJobsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompanyService_ServiceDesc.Streams[0], "/user.CompanyService/GetAllJobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &companyServiceGetAllJobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompanyService_GetAllJobsClient interface {
	Recv() (*JobResponse, error)
	grpc.ClientStream
}

type companyServiceGetAllJobsClient struct {
	grpc.ClientStream
}

func (x *companyServiceGetAllJobsClient) Recv() (*JobResponse, error) {
	m := new(JobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *companyServiceClient) GetJob(ctx context.Context, in *GetJobById, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/user.CompanyService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetAllJobsForCompany(ctx context.Context, in *GetJobByCompanyId, opts ...grpc.CallOption) (CompanyService_GetAllJobsForCompanyClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompanyService_ServiceDesc.Streams[1], "/user.CompanyService/GetAllJobsForCompany", opts...)
	if err != nil {
		return nil, err
	}
	x := &companyServiceGetAllJobsForCompanyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompanyService_GetAllJobsForCompanyClient interface {
	Recv() (*JobResponse, error)
	grpc.ClientStream
}

type companyServiceGetAllJobsForCompanyClient struct {
	grpc.ClientStream
}

func (x *companyServiceGetAllJobsForCompanyClient) Recv() (*JobResponse, error) {
	m := new(JobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *companyServiceClient) CompanyAddJobSkill(ctx context.Context, in *AddJobSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.CompanyService/CompanyAddJobSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) CompanyUpdateJobSkill(ctx context.Context, in *AddJobSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.CompanyService/CompanyUpdateJobSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetAllJobSkill(ctx context.Context, in *GetJobById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.CompanyService/GetAllJobSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteJobSkill(ctx context.Context, in *JobSkillId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.CompanyService/DeleteJobSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations must embed UnimplementedCompanyServiceServer
// for forward compatibility
type CompanyServiceServer interface {
	CompanySignup(context.Context, *CompanySignupRequest) (*CompanySignupResponse, error)
	CompanyLogin(context.Context, *CompanyLoginRequest) (*CompanySignupResponse, error)
	AddJobs(context.Context, *AddJobRequest) (*JobResponse, error)
	UpdateJobs(context.Context, *UpdateJobRequest) (*emptypb.Empty, error)
	DeleteJob(context.Context, *GetJobById) (*emptypb.Empty, error)
	GetAllJobs(*emptypb.Empty, CompanyService_GetAllJobsServer) error
	GetJob(context.Context, *GetJobById) (*JobResponse, error)
	GetAllJobsForCompany(*GetJobByCompanyId, CompanyService_GetAllJobsForCompanyServer) error
	CompanyAddJobSkill(context.Context, *AddJobSkillRequest) (*emptypb.Empty, error)
	CompanyUpdateJobSkill(context.Context, *AddJobSkillRequest) (*emptypb.Empty, error)
	GetAllJobSkill(context.Context, *GetJobById) (*emptypb.Empty, error)
	DeleteJobSkill(context.Context, *JobSkillId) (*emptypb.Empty, error)
	mustEmbedUnimplementedCompanyServiceServer()
}

// UnimplementedCompanyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (UnimplementedCompanyServiceServer) CompanySignup(context.Context, *CompanySignupRequest) (*CompanySignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanySignup not implemented")
}
func (UnimplementedCompanyServiceServer) CompanyLogin(context.Context, *CompanyLoginRequest) (*CompanySignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyLogin not implemented")
}
func (UnimplementedCompanyServiceServer) AddJobs(context.Context, *AddJobRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddJobs not implemented")
}
func (UnimplementedCompanyServiceServer) UpdateJobs(context.Context, *UpdateJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJobs not implemented")
}
func (UnimplementedCompanyServiceServer) DeleteJob(context.Context, *GetJobById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedCompanyServiceServer) GetAllJobs(*emptypb.Empty, CompanyService_GetAllJobsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllJobs not implemented")
}
func (UnimplementedCompanyServiceServer) GetJob(context.Context, *GetJobById) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedCompanyServiceServer) GetAllJobsForCompany(*GetJobByCompanyId, CompanyService_GetAllJobsForCompanyServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllJobsForCompany not implemented")
}
func (UnimplementedCompanyServiceServer) CompanyAddJobSkill(context.Context, *AddJobSkillRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyAddJobSkill not implemented")
}
func (UnimplementedCompanyServiceServer) CompanyUpdateJobSkill(context.Context, *AddJobSkillRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyUpdateJobSkill not implemented")
}
func (UnimplementedCompanyServiceServer) GetAllJobSkill(context.Context, *GetJobById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllJobSkill not implemented")
}
func (UnimplementedCompanyServiceServer) DeleteJobSkill(context.Context, *JobSkillId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJobSkill not implemented")
}
func (UnimplementedCompanyServiceServer) mustEmbedUnimplementedCompanyServiceServer() {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_CompanySignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanySignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CompanySignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/CompanySignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CompanySignup(ctx, req.(*CompanySignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_CompanyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CompanyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/CompanyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CompanyLogin(ctx, req.(*CompanyLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_AddJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).AddJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/AddJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).AddJobs(ctx, req.(*AddJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/UpdateJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateJobs(ctx, req.(*UpdateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/DeleteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteJob(ctx, req.(*GetJobById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetAllJobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompanyServiceServer).GetAllJobs(m, &companyServiceGetAllJobsServer{stream})
}

type CompanyService_GetAllJobsServer interface {
	Send(*JobResponse) error
	grpc.ServerStream
}

type companyServiceGetAllJobsServer struct {
	grpc.ServerStream
}

func (x *companyServiceGetAllJobsServer) Send(m *JobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CompanyService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetJob(ctx, req.(*GetJobById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetAllJobsForCompany_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetJobByCompanyId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompanyServiceServer).GetAllJobsForCompany(m, &companyServiceGetAllJobsForCompanyServer{stream})
}

type CompanyService_GetAllJobsForCompanyServer interface {
	Send(*JobResponse) error
	grpc.ServerStream
}

type companyServiceGetAllJobsForCompanyServer struct {
	grpc.ServerStream
}

func (x *companyServiceGetAllJobsForCompanyServer) Send(m *JobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CompanyService_CompanyAddJobSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJobSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CompanyAddJobSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/CompanyAddJobSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CompanyAddJobSkill(ctx, req.(*AddJobSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_CompanyUpdateJobSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJobSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CompanyUpdateJobSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/CompanyUpdateJobSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CompanyUpdateJobSkill(ctx, req.(*AddJobSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetAllJobSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetAllJobSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/GetAllJobSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetAllJobSkill(ctx, req.(*GetJobById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteJobSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSkillId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteJobSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CompanyService/DeleteJobSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteJobSkill(ctx, req.(*JobSkillId))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompanySignup",
			Handler:    _CompanyService_CompanySignup_Handler,
		},
		{
			MethodName: "CompanyLogin",
			Handler:    _CompanyService_CompanyLogin_Handler,
		},
		{
			MethodName: "AddJobs",
			Handler:    _CompanyService_AddJobs_Handler,
		},
		{
			MethodName: "UpdateJobs",
			Handler:    _CompanyService_UpdateJobs_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _CompanyService_DeleteJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _CompanyService_GetJob_Handler,
		},
		{
			MethodName: "CompanyAddJobSkill",
			Handler:    _CompanyService_CompanyAddJobSkill_Handler,
		},
		{
			MethodName: "CompanyUpdateJobSkill",
			Handler:    _CompanyService_CompanyUpdateJobSkill_Handler,
		},
		{
			MethodName: "GetAllJobSkill",
			Handler:    _CompanyService_GetAllJobSkill_Handler,
		},
		{
			MethodName: "DeleteJobSkill",
			Handler:    _CompanyService_DeleteJobSkill_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllJobs",
			Handler:       _CompanyService_GetAllJobs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllJobsForCompany",
			Handler:       _CompanyService_GetAllJobsForCompany_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "company.proto",
}
