// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/response/users.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateUserResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1610be0aa29382, []int{0}
}

func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (m *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(m, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

func (m *CreateUserResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListUsersResp struct {
	Users                []*common.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListUsersResp) Reset()         { *m = ListUsersResp{} }
func (m *ListUsersResp) String() string { return proto.CompactTextString(m) }
func (*ListUsersResp) ProtoMessage()    {}
func (*ListUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1610be0aa29382, []int{1}
}

func (m *ListUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResp.Unmarshal(m, b)
}
func (m *ListUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResp.Marshal(b, m, deterministic)
}
func (m *ListUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResp.Merge(m, src)
}
func (m *ListUsersResp) XXX_Size() int {
	return xxx_messageInfo_ListUsersResp.Size(m)
}
func (m *ListUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResp proto.InternalMessageInfo

func (m *ListUsersResp) GetUsers() []*common.User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeleteUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResp) Reset()         { *m = DeleteUserResp{} }
func (m *DeleteUserResp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResp) ProtoMessage()    {}
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1610be0aa29382, []int{2}
}

func (m *DeleteUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResp.Unmarshal(m, b)
}
func (m *DeleteUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResp.Marshal(b, m, deterministic)
}
func (m *DeleteUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResp.Merge(m, src)
}
func (m *DeleteUserResp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResp.Size(m)
}
func (m *DeleteUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResp proto.InternalMessageInfo

type GetUserResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1610be0aa29382, []int{3}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateUserResp) Reset()         { *m = UpdateUserResp{} }
func (m *UpdateUserResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResp) ProtoMessage()    {}
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1610be0aa29382, []int{4}
}

func (m *UpdateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResp.Unmarshal(m, b)
}
func (m *UpdateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResp.Merge(m, src)
}
func (m *UpdateUserResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResp.Size(m)
}
func (m *UpdateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResp proto.InternalMessageInfo

func (m *UpdateUserResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateSelfResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateSelfResp) Reset()         { *m = UpdateSelfResp{} }
func (m *UpdateSelfResp) String() string { return proto.CompactTextString(m) }
func (*UpdateSelfResp) ProtoMessage()    {}
func (*UpdateSelfResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1610be0aa29382, []int{5}
}

func (m *UpdateSelfResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSelfResp.Unmarshal(m, b)
}
func (m *UpdateSelfResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSelfResp.Marshal(b, m, deterministic)
}
func (m *UpdateSelfResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSelfResp.Merge(m, src)
}
func (m *UpdateSelfResp) XXX_Size() int {
	return xxx_messageInfo_UpdateSelfResp.Size(m)
}
func (m *UpdateSelfResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSelfResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSelfResp proto.InternalMessageInfo

func (m *UpdateSelfResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserResp)(nil), "chef.automate.api.iam.v2.CreateUserResp")
	proto.RegisterType((*ListUsersResp)(nil), "chef.automate.api.iam.v2.ListUsersResp")
	proto.RegisterType((*DeleteUserResp)(nil), "chef.automate.api.iam.v2.DeleteUserResp")
	proto.RegisterType((*GetUserResp)(nil), "chef.automate.api.iam.v2.GetUserResp")
	proto.RegisterType((*UpdateUserResp)(nil), "chef.automate.api.iam.v2.UpdateUserResp")
	proto.RegisterType((*UpdateSelfResp)(nil), "chef.automate.api.iam.v2.UpdateSelfResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/response/users.proto", fileDescriptor_ce1610be0aa29382)
}

var fileDescriptor_ce1610be0aa29382 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0x39, 0xfc, 0x53, 0x6c, 0xf0, 0x90, 0x54, 0x87, 0x85, 0x84, 0xab, 0x6c, 0xdc, 0x81,
	0xd3, 0x46, 0xac, 0xd4, 0x13, 0x45, 0xac, 0x94, 0x34, 0x76, 0x93, 0x73, 0x92, 0x2c, 0x64, 0x6f,
	0x96, 0x9d, 0xb9, 0x88, 0xdf, 0x5e, 0x6e, 0xc5, 0x68, 0x23, 0x44, 0xaf, 0x5d, 0xde, 0xfb, 0xbd,
	0xb7, 0xbc, 0x31, 0x97, 0x0d, 0xfb, 0xc0, 0x2d, 0xb5, 0x2a, 0x80, 0x9d, 0xb2, 0x47, 0xa5, 0xd3,
	0x05, 0x2a, 0xbd, 0xe1, 0x3b, 0x60, 0x70, 0xe0, 0xd0, 0xc3, 0xba, 0x82, 0x48, 0x12, 0xb8, 0x15,
	0x82, 0x4e, 0x28, 0x8a, 0x0d, 0x91, 0x95, 0xc7, 0x45, 0xb3, 0xa4, 0xb9, 0xfd, 0xb2, 0x59, 0x0c,
	0xce, 0x3a, 0xf4, 0x76, 0x5d, 0x1d, 0x5d, 0x6c, 0x89, 0x6d, 0xd8, 0x7b, 0x6e, 0x7f, 0x42, 0xcb,
	0xda, 0xe4, 0x37, 0x91, 0x50, 0x69, 0x2a, 0x14, 0x9f, 0x48, 0xc2, 0xb8, 0x32, 0xbb, 0xbd, 0xa0,
	0xc8, 0x26, 0xd9, 0xc9, 0xa8, 0x3a, 0xb6, 0xbf, 0xa5, 0xda, 0xe4, 0x48, 0xda, 0xf2, 0xd6, 0x1c,
	0x3c, 0x3a, 0xd1, 0xfe, 0x45, 0x12, 0xe4, 0xdc, 0xec, 0xa5, 0x94, 0x22, 0x9b, 0xec, 0x6c, 0x41,
	0xf9, 0x14, 0x97, 0x87, 0x26, 0xaf, 0x69, 0x45, 0xdf, 0x65, 0xca, 0x2b, 0x33, 0xba, 0x23, 0x1d,
	0xd4, 0xad, 0x36, 0xf9, 0x34, 0xbc, 0x0e, 0xfd, 0xe1, 0x86, 0xf2, 0x4c, 0xab, 0xf9, 0x7f, 0x29,
	0xd7, 0x0f, 0x2f, 0xf7, 0x0b, 0xa7, 0xcb, 0x6e, 0x66, 0x1b, 0xf6, 0xd0, 0x3b, 0x36, 0x7b, 0xc1,
	0x1f, 0x4f, 0x63, 0xb6, 0x9f, 0x06, 0x3c, 0xfb, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xa7, 0x04,
	0xc3, 0x54, 0x02, 0x00, 0x00,
}