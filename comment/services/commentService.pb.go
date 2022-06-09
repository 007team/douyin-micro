// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commentService.proto

package services

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommentActionRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	VideoId              int64    `protobuf:"varint,3,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	ActionType           int32    `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	CommentText          string   `protobuf:"bytes,5,opt,name=comment_text,json=commentText,proto3" json:"comment_text,omitempty"`
	CommentId            int64    `protobuf:"varint,6,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentActionRequest) Reset()         { *m = CommentActionRequest{} }
func (m *CommentActionRequest) String() string { return proto.CompactTextString(m) }
func (*CommentActionRequest) ProtoMessage()    {}
func (*CommentActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cf873f7f4ec518, []int{0}
}

func (m *CommentActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentActionRequest.Unmarshal(m, b)
}
func (m *CommentActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentActionRequest.Marshal(b, m, deterministic)
}
func (m *CommentActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentActionRequest.Merge(m, src)
}
func (m *CommentActionRequest) XXX_Size() int {
	return xxx_messageInfo_CommentActionRequest.Size(m)
}
func (m *CommentActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentActionRequest proto.InternalMessageInfo

func (m *CommentActionRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CommentActionRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommentActionRequest) GetVideoId() int64 {
	if m != nil {
		return m.VideoId
	}
	return 0
}

func (m *CommentActionRequest) GetActionType() int32 {
	if m != nil {
		return m.ActionType
	}
	return 0
}

func (m *CommentActionRequest) GetCommentText() string {
	if m != nil {
		return m.CommentText
	}
	return ""
}

func (m *CommentActionRequest) GetCommentId() int64 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

type CommentActionResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Comment              *Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentActionResponse) Reset()         { *m = CommentActionResponse{} }
func (m *CommentActionResponse) String() string { return proto.CompactTextString(m) }
func (*CommentActionResponse) ProtoMessage()    {}
func (*CommentActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cf873f7f4ec518, []int{1}
}

func (m *CommentActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentActionResponse.Unmarshal(m, b)
}
func (m *CommentActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentActionResponse.Marshal(b, m, deterministic)
}
func (m *CommentActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentActionResponse.Merge(m, src)
}
func (m *CommentActionResponse) XXX_Size() int {
	return xxx_messageInfo_CommentActionResponse.Size(m)
}
func (m *CommentActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentActionResponse proto.InternalMessageInfo

func (m *CommentActionResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CommentActionResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *CommentActionResponse) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

type CommentListRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VideoId              int64    `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentListRequest) Reset()         { *m = CommentListRequest{} }
func (m *CommentListRequest) String() string { return proto.CompactTextString(m) }
func (*CommentListRequest) ProtoMessage()    {}
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cf873f7f4ec518, []int{2}
}

func (m *CommentListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentListRequest.Unmarshal(m, b)
}
func (m *CommentListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentListRequest.Marshal(b, m, deterministic)
}
func (m *CommentListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentListRequest.Merge(m, src)
}
func (m *CommentListRequest) XXX_Size() int {
	return xxx_messageInfo_CommentListRequest.Size(m)
}
func (m *CommentListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentListRequest proto.InternalMessageInfo

func (m *CommentListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommentListRequest) GetVideoId() int64 {
	if m != nil {
		return m.VideoId
	}
	return 0
}

type CommentListResponse struct {
	StatusCode           int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	CommentList          []*Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CommentListResponse) Reset()         { *m = CommentListResponse{} }
func (m *CommentListResponse) String() string { return proto.CompactTextString(m) }
func (*CommentListResponse) ProtoMessage()    {}
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cf873f7f4ec518, []int{3}
}

func (m *CommentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentListResponse.Unmarshal(m, b)
}
func (m *CommentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentListResponse.Marshal(b, m, deterministic)
}
func (m *CommentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentListResponse.Merge(m, src)
}
func (m *CommentListResponse) XXX_Size() int {
	return xxx_messageInfo_CommentListResponse.Size(m)
}
func (m *CommentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentListResponse proto.InternalMessageInfo

func (m *CommentListResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CommentListResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *CommentListResponse) GetCommentList() []*Comment {
	if m != nil {
		return m.CommentList
	}
	return nil
}

func init() {
	proto.RegisterType((*CommentActionRequest)(nil), "services.CommentActionRequest")
	proto.RegisterType((*CommentActionResponse)(nil), "services.CommentActionResponse")
	proto.RegisterType((*CommentListRequest)(nil), "services.CommentListRequest")
	proto.RegisterType((*CommentListResponse)(nil), "services.CommentListResponse")
}

func init() { proto.RegisterFile("commentService.proto", fileDescriptor_b5cf873f7f4ec518) }

var fileDescriptor_b5cf873f7f4ec518 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0xcd, 0xc0, 0x2b, 0xd0, 0xdb, 0xf7, 0x5e, 0xf2, 0x06, 0x5e, 0xac, 0x44, 0x04, 0xbb, 0x22,
	0x31, 0xa9, 0x09, 0xba, 0x73, 0xa5, 0xc4, 0x05, 0x46, 0x12, 0x53, 0x59, 0xb9, 0x69, 0xb0, 0x73,
	0x43, 0x1a, 0xa1, 0x53, 0x99, 0x81, 0xc0, 0xde, 0x9d, 0xff, 0xe2, 0x4f, 0xf8, 0x63, 0xa6, 0x33,
	0x53, 0xa5, 0x22, 0x3b, 0x57, 0x93, 0x7b, 0xee, 0xc9, 0x39, 0x73, 0x4e, 0x2e, 0x34, 0x22, 0x3e,
	0x9b, 0x61, 0x22, 0xef, 0x70, 0xbe, 0x8c, 0x23, 0xf4, 0xd3, 0x39, 0x97, 0x9c, 0xd6, 0x84, 0x1e,
	0x45, 0xb3, 0x6e, 0xf6, 0x43, 0xce, 0x70, 0x2a, 0xf4, 0xda, 0x7b, 0x23, 0xd0, 0xe8, 0x6b, 0xfc,
	0x22, 0x92, 0x31, 0x4f, 0x02, 0x7c, 0x5a, 0xa0, 0x90, 0x74, 0x0f, 0xaa, 0x0b, 0x81, 0xf3, 0x30,
	0x66, 0x2e, 0xe9, 0x90, 0x6e, 0x39, 0xa8, 0x64, 0xe3, 0x80, 0xd1, 0x06, 0x58, 0x92, 0x3f, 0x62,
	0xe2, 0x96, 0x3a, 0xa4, 0x6b, 0x07, 0x7a, 0xa0, 0xfb, 0x50, 0x5b, 0xc6, 0x0c, 0x79, 0xc6, 0x2f,
	0x2b, 0x7e, 0x55, 0xcd, 0x03, 0x46, 0xdb, 0xe0, 0x8c, 0x95, 0x74, 0x28, 0xd7, 0x29, 0xba, 0xbf,
	0x3a, 0xa4, 0x6b, 0x05, 0xa0, 0xa1, 0xd1, 0x3a, 0x45, 0x7a, 0x04, 0xbf, 0xcd, 0xd7, 0x42, 0x89,
	0x2b, 0xe9, 0x5a, 0x4a, 0xd8, 0x31, 0xd8, 0x08, 0x57, 0x92, 0xb6, 0x00, 0x72, 0x4a, 0xcc, 0xdc,
	0x8a, 0x32, 0xb0, 0x0d, 0x32, 0x60, 0xde, 0x33, 0x81, 0xff, 0x5f, 0x52, 0x88, 0x94, 0x27, 0x02,
	0x33, 0x73, 0x21, 0xc7, 0x72, 0x21, 0xc2, 0x88, 0x33, 0x54, 0x51, 0xac, 0x00, 0x34, 0xd4, 0xe7,
	0x0c, 0x33, 0x65, 0x43, 0x98, 0x89, 0x89, 0xc9, 0x64, 0x6b, 0x64, 0x28, 0x26, 0xf4, 0x18, 0xaa,
	0xc6, 0x46, 0xc5, 0x72, 0x7a, 0xff, 0xfc, 0xbc, 0x50, 0xdf, 0x38, 0x06, 0x39, 0xc3, 0xbb, 0x02,
	0x6a, 0xb0, 0x9b, 0x58, 0xc8, 0xbc, 0xc9, 0x8f, 0xc2, 0xc8, 0xae, 0xc2, 0x4a, 0x85, 0xc2, 0xbc,
	0x17, 0x02, 0xf5, 0x82, 0xce, 0x0f, 0x65, 0x39, 0xfb, 0xec, 0x79, 0x1a, 0x8b, 0x2c, 0x50, 0xf9,
	0xfb, 0x40, 0x79, 0xf5, 0x99, 0x7b, 0xef, 0x95, 0xc0, 0xdf, 0x7e, 0xe1, 0xb2, 0xe8, 0x2d, 0xfc,
	0x29, 0xb4, 0x4d, 0x0f, 0xb7, 0x34, 0x0a, 0xc7, 0xd4, 0x6c, 0xef, 0xdc, 0x9b, 0x68, 0xd7, 0xe0,
	0x6c, 0x24, 0xa6, 0x07, 0x5b, 0xfc, 0x8d, 0x42, 0x9b, 0xad, 0x1d, 0x5b, 0xad, 0x75, 0xe9, 0xdc,
	0xdb, 0xfe, 0xc9, 0xb9, 0x3a, 0x6f, 0xf1, 0x50, 0x51, 0xef, 0xe9, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x12, 0x84, 0x0a, 0x1d, 0x03, 0x00, 0x00,
}
