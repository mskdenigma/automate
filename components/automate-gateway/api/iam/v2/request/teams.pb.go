// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/request/teams.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type ListTeamsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTeamsReq) Reset()         { *m = ListTeamsReq{} }
func (m *ListTeamsReq) String() string { return proto.CompactTextString(m) }
func (*ListTeamsReq) ProtoMessage()    {}
func (*ListTeamsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{0}
}

func (m *ListTeamsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamsReq.Unmarshal(m, b)
}
func (m *ListTeamsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamsReq.Marshal(b, m, deterministic)
}
func (m *ListTeamsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamsReq.Merge(m, src)
}
func (m *ListTeamsReq) XXX_Size() int {
	return xxx_messageInfo_ListTeamsReq.Size(m)
}
func (m *ListTeamsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamsReq proto.InternalMessageInfo

type GetTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamReq) Reset()         { *m = GetTeamReq{} }
func (m *GetTeamReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamReq) ProtoMessage()    {}
func (*GetTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{1}
}

func (m *GetTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamReq.Unmarshal(m, b)
}
func (m *GetTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamReq.Marshal(b, m, deterministic)
}
func (m *GetTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamReq.Merge(m, src)
}
func (m *GetTeamReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamReq.Size(m)
}
func (m *GetTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamReq proto.InternalMessageInfo

func (m *GetTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamReq) Reset()         { *m = CreateTeamReq{} }
func (m *CreateTeamReq) String() string { return proto.CompactTextString(m) }
func (*CreateTeamReq) ProtoMessage()    {}
func (*CreateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{2}
}

func (m *CreateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamReq.Unmarshal(m, b)
}
func (m *CreateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamReq.Marshal(b, m, deterministic)
}
func (m *CreateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamReq.Merge(m, src)
}
func (m *CreateTeamReq) XXX_Size() int {
	return xxx_messageInfo_CreateTeamReq.Size(m)
}
func (m *CreateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamReq proto.InternalMessageInfo

func (m *CreateTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTeamReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTeamReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type UpdateTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamReq) Reset()         { *m = UpdateTeamReq{} }
func (m *UpdateTeamReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamReq) ProtoMessage()    {}
func (*UpdateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{3}
}

func (m *UpdateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamReq.Unmarshal(m, b)
}
func (m *UpdateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamReq.Marshal(b, m, deterministic)
}
func (m *UpdateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamReq.Merge(m, src)
}
func (m *UpdateTeamReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamReq.Size(m)
}
func (m *UpdateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamReq proto.InternalMessageInfo

func (m *UpdateTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTeamReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTeamReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTeamReq) Reset()         { *m = DeleteTeamReq{} }
func (m *DeleteTeamReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTeamReq) ProtoMessage()    {}
func (*DeleteTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{4}
}

func (m *DeleteTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTeamReq.Unmarshal(m, b)
}
func (m *DeleteTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTeamReq.Marshal(b, m, deterministic)
}
func (m *DeleteTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTeamReq.Merge(m, src)
}
func (m *DeleteTeamReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTeamReq.Size(m)
}
func (m *DeleteTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTeamReq proto.InternalMessageInfo

func (m *DeleteTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddTeamMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MembershipIds        []string `protobuf:"bytes,2,rep,name=membership_ids,json=membershipIds,proto3" json:"membership_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTeamMembersReq) Reset()         { *m = AddTeamMembersReq{} }
func (m *AddTeamMembersReq) String() string { return proto.CompactTextString(m) }
func (*AddTeamMembersReq) ProtoMessage()    {}
func (*AddTeamMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{5}
}

func (m *AddTeamMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTeamMembersReq.Unmarshal(m, b)
}
func (m *AddTeamMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTeamMembersReq.Marshal(b, m, deterministic)
}
func (m *AddTeamMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTeamMembersReq.Merge(m, src)
}
func (m *AddTeamMembersReq) XXX_Size() int {
	return xxx_messageInfo_AddTeamMembersReq.Size(m)
}
func (m *AddTeamMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTeamMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddTeamMembersReq proto.InternalMessageInfo

func (m *AddTeamMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddTeamMembersReq) GetMembershipIds() []string {
	if m != nil {
		return m.MembershipIds
	}
	return nil
}

type GetTeamMembershipReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamMembershipReq) Reset()         { *m = GetTeamMembershipReq{} }
func (m *GetTeamMembershipReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamMembershipReq) ProtoMessage()    {}
func (*GetTeamMembershipReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{6}
}

func (m *GetTeamMembershipReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamMembershipReq.Unmarshal(m, b)
}
func (m *GetTeamMembershipReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamMembershipReq.Marshal(b, m, deterministic)
}
func (m *GetTeamMembershipReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamMembershipReq.Merge(m, src)
}
func (m *GetTeamMembershipReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamMembershipReq.Size(m)
}
func (m *GetTeamMembershipReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamMembershipReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamMembershipReq proto.InternalMessageInfo

func (m *GetTeamMembershipReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveTeamMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MembershipIds        []string `protobuf:"bytes,2,rep,name=membership_ids,json=membershipIds,proto3" json:"membership_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamMembersReq) Reset()         { *m = RemoveTeamMembersReq{} }
func (m *RemoveTeamMembersReq) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamMembersReq) ProtoMessage()    {}
func (*RemoveTeamMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{7}
}

func (m *RemoveTeamMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamMembersReq.Unmarshal(m, b)
}
func (m *RemoveTeamMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamMembersReq.Marshal(b, m, deterministic)
}
func (m *RemoveTeamMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamMembersReq.Merge(m, src)
}
func (m *RemoveTeamMembersReq) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamMembersReq.Size(m)
}
func (m *RemoveTeamMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamMembersReq proto.InternalMessageInfo

func (m *RemoveTeamMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RemoveTeamMembersReq) GetMembershipIds() []string {
	if m != nil {
		return m.MembershipIds
	}
	return nil
}

type GetTeamsForMemberReq struct {
	MembershipId         string   `protobuf:"bytes,1,opt,name=membership_id,json=membershipId,proto3" json:"membership_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsForMemberReq) Reset()         { *m = GetTeamsForMemberReq{} }
func (m *GetTeamsForMemberReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamsForMemberReq) ProtoMessage()    {}
func (*GetTeamsForMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{8}
}

func (m *GetTeamsForMemberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsForMemberReq.Unmarshal(m, b)
}
func (m *GetTeamsForMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsForMemberReq.Marshal(b, m, deterministic)
}
func (m *GetTeamsForMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsForMemberReq.Merge(m, src)
}
func (m *GetTeamsForMemberReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamsForMemberReq.Size(m)
}
func (m *GetTeamsForMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsForMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsForMemberReq proto.InternalMessageInfo

func (m *GetTeamsForMemberReq) GetMembershipId() string {
	if m != nil {
		return m.MembershipId
	}
	return ""
}

type ApplyV2DataMigrationsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyV2DataMigrationsReq) Reset()         { *m = ApplyV2DataMigrationsReq{} }
func (m *ApplyV2DataMigrationsReq) String() string { return proto.CompactTextString(m) }
func (*ApplyV2DataMigrationsReq) ProtoMessage()    {}
func (*ApplyV2DataMigrationsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{9}
}

func (m *ApplyV2DataMigrationsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyV2DataMigrationsReq.Unmarshal(m, b)
}
func (m *ApplyV2DataMigrationsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyV2DataMigrationsReq.Marshal(b, m, deterministic)
}
func (m *ApplyV2DataMigrationsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyV2DataMigrationsReq.Merge(m, src)
}
func (m *ApplyV2DataMigrationsReq) XXX_Size() int {
	return xxx_messageInfo_ApplyV2DataMigrationsReq.Size(m)
}
func (m *ApplyV2DataMigrationsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyV2DataMigrationsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyV2DataMigrationsReq proto.InternalMessageInfo

type ResetAllTeamProjectsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetAllTeamProjectsReq) Reset()         { *m = ResetAllTeamProjectsReq{} }
func (m *ResetAllTeamProjectsReq) String() string { return proto.CompactTextString(m) }
func (*ResetAllTeamProjectsReq) ProtoMessage()    {}
func (*ResetAllTeamProjectsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5377b66d36621f2, []int{10}
}

func (m *ResetAllTeamProjectsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetAllTeamProjectsReq.Unmarshal(m, b)
}
func (m *ResetAllTeamProjectsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetAllTeamProjectsReq.Marshal(b, m, deterministic)
}
func (m *ResetAllTeamProjectsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetAllTeamProjectsReq.Merge(m, src)
}
func (m *ResetAllTeamProjectsReq) XXX_Size() int {
	return xxx_messageInfo_ResetAllTeamProjectsReq.Size(m)
}
func (m *ResetAllTeamProjectsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetAllTeamProjectsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetAllTeamProjectsReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListTeamsReq)(nil), "chef.automate.api.iam.v2.ListTeamsReq")
	proto.RegisterType((*GetTeamReq)(nil), "chef.automate.api.iam.v2.GetTeamReq")
	proto.RegisterType((*CreateTeamReq)(nil), "chef.automate.api.iam.v2.CreateTeamReq")
	proto.RegisterType((*UpdateTeamReq)(nil), "chef.automate.api.iam.v2.UpdateTeamReq")
	proto.RegisterType((*DeleteTeamReq)(nil), "chef.automate.api.iam.v2.DeleteTeamReq")
	proto.RegisterType((*AddTeamMembersReq)(nil), "chef.automate.api.iam.v2.AddTeamMembersReq")
	proto.RegisterType((*GetTeamMembershipReq)(nil), "chef.automate.api.iam.v2.GetTeamMembershipReq")
	proto.RegisterType((*RemoveTeamMembersReq)(nil), "chef.automate.api.iam.v2.RemoveTeamMembersReq")
	proto.RegisterType((*GetTeamsForMemberReq)(nil), "chef.automate.api.iam.v2.GetTeamsForMemberReq")
	proto.RegisterType((*ApplyV2DataMigrationsReq)(nil), "chef.automate.api.iam.v2.ApplyV2DataMigrationsReq")
	proto.RegisterType((*ResetAllTeamProjectsReq)(nil), "chef.automate.api.iam.v2.ResetAllTeamProjectsReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/request/teams.proto", fileDescriptor_b5377b66d36621f2)
}

var fileDescriptor_b5377b66d36621f2 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xdf, 0x6a, 0xd4, 0x4e,
	0x14, 0x26, 0xdb, 0xdf, 0x4f, 0xed, 0xd0, 0x2d, 0x1a, 0xaa, 0xae, 0xbd, 0x0a, 0x23, 0x42, 0xc1,
	0x66, 0xd3, 0x9d, 0xfd, 0xd3, 0x36, 0x17, 0x62, 0xb4, 0x58, 0x0a, 0x2e, 0xd6, 0x58, 0xbd, 0x50,
	0x44, 0x4e, 0x32, 0x27, 0xe9, 0xc8, 0x26, 0x93, 0xcd, 0x4c, 0x5b, 0x4a, 0x11, 0xbc, 0xf0, 0x09,
	0xf6, 0x19, 0x7c, 0x05, 0x1f, 0xc4, 0xe2, 0x03, 0xc9, 0x64, 0xb7, 0x6b, 0xb7, 0x58, 0x44, 0xf4,
	0xca, 0xbb, 0x39, 0xe7, 0x7c, 0x73, 0xce, 0xf7, 0x9d, 0x19, 0x3e, 0xe2, 0xc7, 0x32, 0x2b, 0x64,
	0x8e, 0xb9, 0x56, 0x1e, 0x1c, 0x68, 0x99, 0x81, 0x46, 0x37, 0x05, 0x8d, 0x47, 0x70, 0xec, 0x41,
	0x21, 0x3c, 0x01, 0x99, 0x77, 0xc8, 0xbc, 0x12, 0x87, 0x07, 0xa8, 0xb4, 0xa7, 0x11, 0x32, 0xd5,
	0x2c, 0x4a, 0xa9, 0xa5, 0xdd, 0x88, 0xf7, 0x31, 0x69, 0x9e, 0xdd, 0x6a, 0x42, 0x21, 0x9a, 0x02,
	0xb2, 0xe6, 0x21, 0x5b, 0x5e, 0xad, 0x00, 0xb1, 0x9b, 0x62, 0xee, 0xaa, 0x23, 0x48, 0x53, 0x2c,
	0x3d, 0x59, 0x68, 0x21, 0x73, 0xe5, 0x41, 0x9e, 0x4b, 0x0d, 0xd5, 0x79, 0xdc, 0x87, 0x2e, 0x92,
	0x85, 0xa7, 0x42, 0xe9, 0x3d, 0xd3, 0x3a, 0xc4, 0x21, 0x5d, 0x21, 0x64, 0x1b, 0xab, 0x30, 0xc4,
	0xa1, 0xbd, 0x48, 0x6a, 0x82, 0x37, 0x2c, 0xc7, 0x5a, 0x99, 0x0f, 0x6b, 0x82, 0xfb, 0x64, 0x14,
	0x5c, 0x25, 0xff, 0x9f, 0x5a, 0x35, 0xc1, 0xe9, 0x17, 0x8b, 0xd4, 0x1f, 0x97, 0x08, 0x1a, 0x2f,
	0x41, 0xdb, 0x36, 0xf9, 0x2f, 0x87, 0x0c, 0x1b, 0xb5, 0x2a, 0x53, 0x9d, 0xed, 0x65, 0x72, 0xad,
	0x28, 0xe5, 0x7b, 0x8c, 0xb5, 0x6a, 0xcc, 0x39, 0x73, 0x2b, 0xf3, 0xe1, 0x34, 0xf6, 0xf9, 0x28,
	0x00, 0xb2, 0x50, 0x75, 0x3f, 0xb5, 0x2a, 0x38, 0x7b, 0x6e, 0x3f, 0x3b, 0xa1, 0x82, 0x53, 0xdf,
	0xa1, 0x1a, 0x95, 0x76, 0x05, 0xa7, 0xab, 0x0e, 0x35, 0x25, 0x93, 0xea, 0x1f, 0x3b, 0x7b, 0xa8,
	0xb4, 0x63, 0x08, 0x98, 0xfc, 0x59, 0x27, 0xea, 0x3b, 0x6f, 0xce, 0x82, 0xd6, 0xb9, 0x02, 0xa3,
	0x6f, 0x3f, 0xd0, 0xcf, 0x16, 0xa9, 0xbf, 0x2c, 0xf8, 0x5f, 0xe4, 0xfd, 0x62, 0x14, 0xec, 0x5e,
	0xe0, 0xfd, 0xd0, 0x7e, 0x70, 0xf2, 0x67, 0x34, 0xef, 0x93, 0xfa, 0x16, 0x0e, 0xf0, 0x52, 0x96,
	0x33, 0x6f, 0xf1, 0xd5, 0x22, 0x37, 0x02, 0xce, 0x0d, 0xb4, 0x8f, 0x59, 0x84, 0xa5, 0xfa, 0x99,
	0xae, 0x7b, 0x64, 0x31, 0x1b, 0x57, 0xf7, 0x45, 0xf1, 0x4e, 0x70, 0xd5, 0xa8, 0x55, 0x4a, 0xea,
	0x3f, 0xb2, 0x3b, 0x5c, 0xf9, 0x9f, 0xac, 0x51, 0xf0, 0xd1, 0x22, 0xb7, 0x26, 0x82, 0x2e, 0xdc,
	0x61, 0x89, 0xcd, 0x4f, 0xe8, 0x6c, 0xae, 0x12, 0xd2, 0x65, 0xeb, 0xc8, 0x37, 0x7b, 0x89, 0xcb,
	0x30, 0x8e, 0xdc, 0x4e, 0xb2, 0x91, 0xb8, 0x10, 0xf1, 0x75, 0x77, 0x2d, 0x4a, 0x7a, 0xdd, 0x56,
	0xab, 0xd3, 0xdd, 0x84, 0xd8, 0x88, 0x6c, 0x77, 0xdb, 0xd0, 0x63, 0xbc, 0xe3, 0x6e, 0x74, 0x13,
	0x70, 0x3b, 0x1d, 0xd6, 0x76, 0xa3, 0x16, 0x03, 0x37, 0xe9, 0xf5, 0xd6, 0x36, 0xa0, 0xdb, 0x63,
	0x80, 0x9b, 0x66, 0x01, 0x8c, 0x2c, 0x4d, 0x7e, 0x62, 0x7f, 0x3a, 0xec, 0x57, 0x7b, 0xf8, 0x66,
	0x91, 0xa5, 0x10, 0x33, 0x79, 0x88, 0xff, 0xd4, 0x2a, 0xc2, 0xe9, 0x2a, 0xd4, 0x13, 0x59, 0x8e,
	0x65, 0x19, 0x55, 0x77, 0x49, 0x7d, 0x66, 0xfa, 0x44, 0xe0, 0xc2, 0x79, 0x11, 0xfe, 0xcd, 0x51,
	0x60, 0x93, 0xeb, 0xa7, 0xd6, 0x2c, 0x94, 0x2e, 0x93, 0x46, 0x50, 0x14, 0x83, 0xe3, 0x57, 0x6c,
	0x0b, 0x34, 0xf4, 0x45, 0x5a, 0x8e, 0x7d, 0xc1, 0x98, 0xc0, 0x1d, 0x72, 0x3b, 0x44, 0x85, 0x3a,
	0x18, 0x0c, 0xcc, 0xd0, 0xdd, 0xc9, 0x7f, 0x0d, 0x71, 0xf8, 0x68, 0xe7, 0xf5, 0x76, 0x2a, 0xf4,
	0xfe, 0x41, 0xd4, 0x8c, 0x65, 0xe6, 0x19, 0x13, 0x9a, 0x5a, 0x97, 0xf7, 0x7b, 0x76, 0x16, 0x5d,
	0xa9, 0x1c, 0xa8, 0xfd, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x0b, 0x55, 0x03, 0x07, 0x05, 0x00,
	0x00,
}
