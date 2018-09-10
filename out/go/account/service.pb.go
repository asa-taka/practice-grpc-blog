// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/service.proto

package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Ignoring public import of User from account/data.proto

// PageRequest from public import common/page.proto
type PageRequest = common.PageRequest

// PageResponse from public import common/page.proto
type PageResponse = common.PageResponse

type QueryUsersRequest struct {
	Page                 *common.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *QueryUsersRequest) Reset()         { *m = QueryUsersRequest{} }
func (m *QueryUsersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUsersRequest) ProtoMessage()    {}
func (*QueryUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{0}
}
func (m *QueryUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUsersRequest.Unmarshal(m, b)
}
func (m *QueryUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUsersRequest.Marshal(b, m, deterministic)
}
func (dst *QueryUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsersRequest.Merge(dst, src)
}
func (m *QueryUsersRequest) XXX_Size() int {
	return xxx_messageInfo_QueryUsersRequest.Size(m)
}
func (m *QueryUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsersRequest proto.InternalMessageInfo

func (m *QueryUsersRequest) GetPage() *common.PageRequest {
	if m != nil {
		return m.Page
	}
	return nil
}

type QueryUsersResponse struct {
	Users                []*User              `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	Page                 *common.PageResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryUsersResponse) Reset()         { *m = QueryUsersResponse{} }
func (m *QueryUsersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUsersResponse) ProtoMessage()    {}
func (*QueryUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{1}
}
func (m *QueryUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUsersResponse.Unmarshal(m, b)
}
func (m *QueryUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUsersResponse.Marshal(b, m, deterministic)
}
func (dst *QueryUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsersResponse.Merge(dst, src)
}
func (m *QueryUsersResponse) XXX_Size() int {
	return xxx_messageInfo_QueryUsersResponse.Size(m)
}
func (m *QueryUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsersResponse proto.InternalMessageInfo

func (m *QueryUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *QueryUsersResponse) GetPage() *common.PageResponse {
	if m != nil {
		return m.Page
	}
	return nil
}

type GetUserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{2}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{3}
}
func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (dst *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(dst, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{4}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{5}
}
func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (dst *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(dst, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{6}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(dst, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{7}
}
func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(dst, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{8}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(dst, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1dac4139c3af3c08, []int{9}
}
func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(dst, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryUsersRequest)(nil), "account.QueryUsersRequest")
	proto.RegisterType((*QueryUsersResponse)(nil), "account.QueryUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "account.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "account.GetUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "account.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "account.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "account.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "account.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "account.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "account.DeleteUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	QueryUsers(ctx context.Context, in *QueryUsersRequest, opts ...grpc.CallOption) (*QueryUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) QueryUsers(ctx context.Context, in *QueryUsersRequest, opts ...grpc.CallOption) (*QueryUsersResponse, error) {
	out := new(QueryUsersResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/QueryUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	QueryUsers(context.Context, *QueryUsersRequest) (*QueryUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_QueryUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).QueryUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/QueryUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).QueryUsers(ctx, req.(*QueryUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryUsers",
			Handler:    _AccountService_QueryUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AccountService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AccountService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AccountService_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AccountService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/service.proto",
}

func init() { proto.RegisterFile("account/service.proto", fileDescriptor_service_1dac4139c3af3c08) }

var fileDescriptor_service_1dac4139c3af3c08 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xa6, 0x15, 0x24, 0x19, 0x22, 0xa6, 0xa3, 0x46, 0x52, 0x2f, 0x58, 0x0e, 0x72, 0x2a, 0x09,
	0x1a, 0xb9, 0x18, 0x13, 0xa3, 0x09, 0x57, 0xac, 0xe1, 0x01, 0x6a, 0x3b, 0x21, 0x24, 0xc2, 0xd6,
	0xdd, 0xad, 0x89, 0xef, 0xed, 0x03, 0x18, 0x76, 0xd7, 0x76, 0x61, 0x3d, 0xa0, 0xd7, 0x9d, 0xef,
	0xaf, 0xf3, 0x4d, 0xe1, 0x2c, 0xcd, 0x32, 0x56, 0xae, 0xe5, 0x48, 0x10, 0xff, 0x58, 0x66, 0x14,
	0x17, 0x9c, 0x49, 0x86, 0x6d, 0xf3, 0x1c, 0xe2, 0xcf, 0x3c, 0x4f, 0x65, 0xaa, 0x87, 0x61, 0x90,
	0xb1, 0xd5, 0x8a, 0xad, 0x47, 0x45, 0xba, 0x30, 0xf8, 0xe8, 0x0e, 0x82, 0xe7, 0x92, 0xf8, 0xe7,
	0x5c, 0x10, 0x17, 0x09, 0xbd, 0x97, 0x24, 0x24, 0x5e, 0x41, 0x73, 0x03, 0xe9, 0x79, 0x7d, 0x6f,
	0xd8, 0x19, 0x9f, 0xc4, 0x9a, 0x16, 0xcf, 0xd2, 0x05, 0x19, 0x48, 0xa2, 0x00, 0x51, 0x06, 0x68,
	0xb3, 0x45, 0xc1, 0xd6, 0x82, 0x70, 0x00, 0x2d, 0xf5, 0xd0, 0xf3, 0xfa, 0x07, 0xc3, 0xce, 0xf8,
	0x28, 0x36, 0x51, 0xe2, 0xcd, 0x6b, 0xa2, 0x67, 0x38, 0x34, 0x1e, 0xbe, 0xf2, 0x38, 0xdd, 0xf6,
	0xd0, 0x42, 0xc6, 0xa4, 0x0f, 0xdd, 0x29, 0x49, 0xc5, 0x35, 0xf9, 0xba, 0xe0, 0x2f, 0x73, 0x95,
	0xae, 0x95, 0xf8, 0xcb, 0x3c, 0xba, 0x81, 0xe3, 0x0a, 0x61, 0x32, 0x5c, 0x42, 0xb3, 0x14, 0xc4,
	0xcd, 0x27, 0xec, 0x44, 0x50, 0xa3, 0xe8, 0x16, 0x82, 0x47, 0x4e, 0xa9, 0x24, 0x5b, 0x7a, 0x0f,
	0xde, 0x04, 0xd0, 0xe6, 0xfd, 0xc9, 0x70, 0x5e, 0xe4, 0xff, 0x32, 0xb4, 0x79, 0xfb, 0x1b, 0x0e,
	0x20, 0x78, 0xa2, 0x37, 0xda, 0x36, 0xdc, 0x5d, 0xde, 0x04, 0xd0, 0x06, 0xed, 0xad, 0x3e, 0xfe,
	0xf2, 0xa1, 0xfb, 0xa0, 0x9f, 0x5f, 0xf4, 0x0d, 0xe2, 0x14, 0xa0, 0xbe, 0x07, 0x0c, 0x2b, 0x96,
	0x73, 0x62, 0xe1, 0xc5, 0xaf, 0x33, 0x6d, 0x1e, 0x35, 0xf0, 0x1e, 0xda, 0xa6, 0x51, 0x3c, 0xaf,
	0x90, 0xdb, 0x57, 0x10, 0xf6, 0xdc, 0x41, 0xc5, 0x9f, 0x02, 0xd4, 0x1d, 0x59, 0x41, 0x9c, 0xc2,
	0xad, 0x20, 0x6e, 0xa9, 0x5a, 0xa8, 0xde, 0x8e, 0x25, 0xe4, 0xec, 0xd5, 0x12, 0x72, 0xd7, 0xa9,
	0x85, 0xea, 0x12, 0x2d, 0x21, 0xe7, 0x22, 0x2c, 0x21, 0xb7, 0xf5, 0xa8, 0x31, 0x6b, 0xcc, 0xbc,
	0xd7, 0x43, 0xf5, 0xf3, 0x5e, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x0d, 0x42, 0x65, 0x05,
	0x04, 0x00, 0x00,
}
