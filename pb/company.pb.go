// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: company.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompanySignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Otp      string `protobuf:"bytes,5,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *CompanySignupRequest) Reset() {
	*x = CompanySignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanySignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanySignupRequest) ProtoMessage() {}

func (x *CompanySignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanySignupRequest.ProtoReflect.Descriptor instead.
func (*CompanySignupRequest) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{0}
}

func (x *CompanySignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CompanySignupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanySignupRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CompanySignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CompanySignupRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type CompanySignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CompanySignupResponse) Reset() {
	*x = CompanySignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanySignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanySignupResponse) ProtoMessage() {}

func (x *CompanySignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanySignupResponse.ProtoReflect.Descriptor instead.
func (*CompanySignupResponse) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{1}
}

func (x *CompanySignupResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompanySignupResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CompanySignupResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanySignupResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CompanyLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CompanyLoginRequest) Reset() {
	*x = CompanyLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyLoginRequest) ProtoMessage() {}

func (x *CompanyLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyLoginRequest.ProtoReflect.Descriptor instead.
func (*CompanyLoginRequest) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{2}
}

func (x *CompanyLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CompanyLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SalaryRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinSalary int64 `protobuf:"varint,1,opt,name=min_salary,json=minSalary,proto3" json:"min_salary,omitempty"`
	MaxSalary int64 `protobuf:"varint,2,opt,name=max_salary,json=maxSalary,proto3" json:"max_salary,omitempty"`
}

func (x *SalaryRange) Reset() {
	*x = SalaryRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalaryRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalaryRange) ProtoMessage() {}

func (x *SalaryRange) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalaryRange.ProtoReflect.Descriptor instead.
func (*SalaryRange) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{3}
}

func (x *SalaryRange) GetMinSalary() int64 {
	if x != nil {
		return x.MinSalary
	}
	return 0
}

func (x *SalaryRange) GetMaxSalary() int64 {
	if x != nil {
		return x.MaxSalary
	}
	return 0
}

type AddJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vacancy       int32                  `protobuf:"varint,1,opt,name=vacancy,proto3" json:"vacancy,omitempty"`
	Designation   string                 `protobuf:"bytes,2,opt,name=designation,proto3" json:"designation,omitempty"`
	Salaryrange   *SalaryRange           `protobuf:"bytes,3,opt,name=salaryrange,proto3" json:"salaryrange,omitempty"`
	MinExperience string                 `protobuf:"bytes,4,opt,name=min_experience,json=minExperience,proto3" json:"min_experience,omitempty"`
	ValidUntil    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
}

func (x *AddJobRequest) Reset() {
	*x = AddJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddJobRequest) ProtoMessage() {}

func (x *AddJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddJobRequest.ProtoReflect.Descriptor instead.
func (*AddJobRequest) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{4}
}

func (x *AddJobRequest) GetVacancy() int32 {
	if x != nil {
		return x.Vacancy
	}
	return 0
}

func (x *AddJobRequest) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

func (x *AddJobRequest) GetSalaryrange() *SalaryRange {
	if x != nil {
		return x.Salaryrange
	}
	return nil
}

func (x *AddJobRequest) GetMinExperience() string {
	if x != nil {
		return x.MinExperience
	}
	return ""
}

func (x *AddJobRequest) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

type JobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company     string                 `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Designation string                 `protobuf:"bytes,2,opt,name=designation,proto3" json:"designation,omitempty"`
	Salaryrange *SalaryRange           `protobuf:"bytes,3,opt,name=salaryrange,proto3" json:"salaryrange,omitempty"`
	Vacancy     int32                  `protobuf:"varint,4,opt,name=vacancy,proto3" json:"vacancy,omitempty"`
	Hired       int32                  `protobuf:"varint,5,opt,name=hired,proto3" json:"hired,omitempty"`
	PostedOn    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=posted_on,json=postedOn,proto3" json:"posted_on,omitempty"`
	ValidUntil  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	Status      string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *JobResponse) Reset() {
	*x = JobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResponse) ProtoMessage() {}

func (x *JobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResponse.ProtoReflect.Descriptor instead.
func (*JobResponse) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{5}
}

func (x *JobResponse) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *JobResponse) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

func (x *JobResponse) GetSalaryrange() *SalaryRange {
	if x != nil {
		return x.Salaryrange
	}
	return nil
}

func (x *JobResponse) GetVacancy() int32 {
	if x != nil {
		return x.Vacancy
	}
	return 0
}

func (x *JobResponse) GetHired() int32 {
	if x != nil {
		return x.Hired
	}
	return 0
}

func (x *JobResponse) GetPostedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.PostedOn
	}
	return nil
}

func (x *JobResponse) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

func (x *JobResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vacancy       int32                  `protobuf:"varint,1,opt,name=vacancy,proto3" json:"vacancy,omitempty"`
	Designation   string                 `protobuf:"bytes,2,opt,name=designation,proto3" json:"designation,omitempty"`
	Salaryrange   *SalaryRange           `protobuf:"bytes,3,opt,name=salaryrange,proto3" json:"salaryrange,omitempty"`
	MinExperience string                 `protobuf:"bytes,4,opt,name=min_experience,json=minExperience,proto3" json:"min_experience,omitempty"`
	ValidUntil    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateJobRequest) GetVacancy() int32 {
	if x != nil {
		return x.Vacancy
	}
	return 0
}

func (x *UpdateJobRequest) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

func (x *UpdateJobRequest) GetSalaryrange() *SalaryRange {
	if x != nil {
		return x.Salaryrange
	}
	return nil
}

func (x *UpdateJobRequest) GetMinExperience() string {
	if x != nil {
		return x.MinExperience
	}
	return ""
}

func (x *UpdateJobRequest) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

func (x *UpdateJobRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetJobById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetJobById) Reset() {
	*x = GetJobById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobById) ProtoMessage() {}

func (x *GetJobById) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobById.ProtoReflect.Descriptor instead.
func (*GetJobById) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{7}
}

func (x *GetJobById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_company_proto protoreflect.FileDescriptor

var file_company_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x67, 0x0a, 0x15, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x0b,
	0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x69, 0x6e, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6d, 0x69, 0x6e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x41, 0x64,
	0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c,
	0x22, 0xbc, 0x02, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0b,
	0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x68,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x68, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xff, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d,
	0x69, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x32,
	0xb2, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x73, 0x12,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x10,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_proto_rawDescOnce sync.Once
	file_company_proto_rawDescData = file_company_proto_rawDesc
)

func file_company_proto_rawDescGZIP() []byte {
	file_company_proto_rawDescOnce.Do(func() {
		file_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_proto_rawDescData)
	})
	return file_company_proto_rawDescData
}

var file_company_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_company_proto_goTypes = []interface{}{
	(*CompanySignupRequest)(nil),  // 0: user.CompanySignupRequest
	(*CompanySignupResponse)(nil), // 1: user.CompanySignupResponse
	(*CompanyLoginRequest)(nil),   // 2: user.CompanyLoginRequest
	(*SalaryRange)(nil),           // 3: user.SalaryRange
	(*AddJobRequest)(nil),         // 4: user.AddJobRequest
	(*JobResponse)(nil),           // 5: user.JobResponse
	(*UpdateJobRequest)(nil),      // 6: user.UpdateJobRequest
	(*GetJobById)(nil),            // 7: user.GetJobById
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_company_proto_depIdxs = []int32{
	3,  // 0: user.AddJobRequest.salaryrange:type_name -> user.SalaryRange
	8,  // 1: user.AddJobRequest.valid_until:type_name -> google.protobuf.Timestamp
	3,  // 2: user.JobResponse.salaryrange:type_name -> user.SalaryRange
	8,  // 3: user.JobResponse.posted_on:type_name -> google.protobuf.Timestamp
	8,  // 4: user.JobResponse.valid_until:type_name -> google.protobuf.Timestamp
	3,  // 5: user.UpdateJobRequest.salaryrange:type_name -> user.SalaryRange
	8,  // 6: user.UpdateJobRequest.valid_until:type_name -> google.protobuf.Timestamp
	0,  // 7: user.CompanyService.CompanySignup:input_type -> user.CompanySignupRequest
	2,  // 8: user.CompanyService.CompanyLogin:input_type -> user.CompanyLoginRequest
	4,  // 9: user.CompanyService.AddJobs:input_type -> user.AddJobRequest
	6,  // 10: user.CompanyService.UpdateJobs:input_type -> user.UpdateJobRequest
	7,  // 11: user.CompanyService.DeleteJobs:input_type -> user.GetJobById
	9,  // 12: user.CompanyService.GetAllJobs:input_type -> google.protobuf.Empty
	7,  // 13: user.CompanyService.GetJob:input_type -> user.GetJobById
	1,  // 14: user.CompanyService.CompanySignup:output_type -> user.CompanySignupResponse
	1,  // 15: user.CompanyService.CompanyLogin:output_type -> user.CompanySignupResponse
	5,  // 16: user.CompanyService.AddJobs:output_type -> user.JobResponse
	5,  // 17: user.CompanyService.UpdateJobs:output_type -> user.JobResponse
	9,  // 18: user.CompanyService.DeleteJobs:output_type -> google.protobuf.Empty
	5,  // 19: user.CompanyService.GetAllJobs:output_type -> user.JobResponse
	5,  // 20: user.CompanyService.GetJob:output_type -> user.JobResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_company_proto_init() }
func file_company_proto_init() {
	if File_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanySignupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanySignupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalaryRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_company_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobById); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_proto_goTypes,
		DependencyIndexes: file_company_proto_depIdxs,
		MessageInfos:      file_company_proto_msgTypes,
	}.Build()
	File_company_proto = out.File
	file_company_proto_rawDesc = nil
	file_company_proto_goTypes = nil
	file_company_proto_depIdxs = nil
}
