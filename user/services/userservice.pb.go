// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userservice.protos

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/protos"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protos package it is being compiled against.
// A compilation error at this line likely means your copy of the
// protos package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the protos package

type UserRequest struct {
	// @inject_tag: json:"user_name" form:"user_name" uri:"user_name"
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// @inject_tag: json:"password" form:"password" uri:"password"
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserDetailResponse struct {
	UserDetail           *UserModel `protobuf:"bytes,1,opt,name=UserDetail,proto3" json:"UserDetail,omitempty"`
	Code                 uint32     `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserDetailResponse) Reset()         { *m = UserDetailResponse{} }
func (m *UserDetailResponse) String() string { return proto.CompactTextString(m) }
func (*UserDetailResponse) ProtoMessage()    {}
func (*UserDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{1}
}

func (m *UserDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetailResponse.Unmarshal(m, b)
}
func (m *UserDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetailResponse.Marshal(b, m, deterministic)
}
func (m *UserDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetailResponse.Merge(m, src)
}
func (m *UserDetailResponse) XXX_Size() int {
	return xxx_messageInfo_UserDetailResponse.Size(m)
}
func (m *UserDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetailResponse proto.InternalMessageInfo

func (m *UserDetailResponse) GetUserDetail() *UserModel {
	if m != nil {
		return m.UserDetail
	}
	return nil
}

func (m *UserDetailResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "services.UserRequest")
	proto.RegisterType((*UserDetailResponse)(nil), "services.UserDetailResponse")
}

func init() { proto.RegisterFile("userservice.protos", fileDescriptor_68a7ca558839fd2b) }

var fileDescriptor_68a7ca558839fd2b = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xa9, 0xff, 0xd8, 0xce, 0xaa, 0x68, 0x44, 0x28, 0x8b, 0x07, 0xe9, 0xc9, 0xd3, 0x0a,
	0xed, 0x51, 0x41, 0x6a, 0xad, 0x20, 0x54, 0x91, 0x48, 0x2f, 0x82, 0x87, 0xb5, 0x3b, 0x96, 0x40,
	0xcc, 0xd4, 0x4c, 0x6a, 0xbf, 0x8c, 0x1f, 0x56, 0x92, 0x74, 0xd5, 0x7a, 0xcc, 0x65, 0x37, 0xef,
	0xbd, 0xc9, 0x2f, 0x2f, 0x10, 0x38, 0x5c, 0x30, 0x5a, 0x46, 0xfb, 0xa9, 0xa6, 0x58, 0xce, 0x2d,
	0x39, 0x12, 0xd9, 0x4a, 0x72, 0x71, 0xe0, 0xc3, 0x7b, 0xaa, 0x51, 0x73, 0xcc, 0xba, 0x23, 0xc8,
	0x27, 0x8c, 0x56, 0xe2, 0xc7, 0x02, 0xd9, 0x89, 0x02, 0x32, 0x2f, 0x1f, 0xaa, 0x77, 0xec, 0xb4,
	0x4e, 0x5b, 0x67, 0x6d, 0xf9, 0xa3, 0x7d, 0xf6, 0x58, 0x31, 0x2f, 0xc9, 0xd6, 0x9d, 0x8d, 0x98,
	0x35, 0xba, 0xfb, 0x02, 0xc2, 0xcf, 0xdd, 0xa0, 0xab, 0x94, 0x96, 0xc8, 0x73, 0x32, 0x8c, 0xa2,
	0x0f, 0xf0, 0xeb, 0x06, 0x5e, 0xde, 0x3b, 0x2a, 0x9b, 0x36, 0xe5, 0xa4, 0x29, 0x23, 0xff, 0x8c,
	0x09, 0x01, 0x5b, 0x43, 0xaa, 0x31, 0x1c, 0xb1, 0x27, 0xc3, 0xba, 0xf7, 0xb5, 0x19, 0x6b, 0x3e,
	0xc5, 0xad, 0xe2, 0x12, 0xb6, 0xc7, 0x34, 0x53, 0x46, 0x1c, 0xaf, 0xd3, 0x56, 0xd7, 0x28, 0x4e,
	0xd6, 0xed, 0x7f, 0xb5, 0xae, 0x20, 0x93, 0x38, 0x53, 0xec, 0xd0, 0x26, 0x03, 0xbc, 0x7b, 0x67,
	0xde, 0x28, 0x0d, 0x30, 0x82, 0x7d, 0x89, 0xba, 0x72, 0x8a, 0xcc, 0x60, 0xea, 0xbf, 0x69, 0x98,
	0x01, 0xc0, 0x2d, 0x69, 0x4d, 0xcb, 0xb1, 0x62, 0x97, 0x86, 0x18, 0xc2, 0x6e, 0x44, 0xa0, 0x4d,
	0x86, 0x5c, 0xe7, 0xcf, 0xed, 0xf2, 0xfc, 0x22, 0x3c, 0x28, 0x7e, 0xdd, 0x09, 0xff, 0xfe, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x99, 0xec, 0x8b, 0x89, 0x02, 0x00, 0x00,
}
