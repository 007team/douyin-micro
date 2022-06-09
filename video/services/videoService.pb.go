// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: videoService.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
//Feed
type VideoFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"latest_time"
	LatestTime int64 `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VideoFeedRequest) Reset() {
	*x = VideoFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFeedRequest) ProtoMessage() {}

func (x *VideoFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFeedRequest.ProtoReflect.Descriptor instead.
func (*VideoFeedRequest) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{0}
}

func (x *VideoFeedRequest) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

func (x *VideoFeedRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VideoFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// @inject_tag: json:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// @inject_tag: json:"video_list"
	VideoList []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
	// @inject_tag: json:"next_time"
	NextTime int64 `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"`
}

func (x *VideoFeedResponse) Reset() {
	*x = VideoFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFeedResponse) ProtoMessage() {}

func (x *VideoFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFeedResponse.ProtoReflect.Descriptor instead.
func (*VideoFeedResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{1}
}

func (x *VideoFeedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VideoFeedResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *VideoFeedResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *VideoFeedResponse) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

//
//PublishAction
type VideoPublishActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// @inject_tag: json:"data"
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *VideoPublishActionRequest) Reset() {
	*x = VideoPublishActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPublishActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPublishActionRequest) ProtoMessage() {}

func (x *VideoPublishActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPublishActionRequest.ProtoReflect.Descriptor instead.
func (*VideoPublishActionRequest) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{2}
}

func (x *VideoPublishActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *VideoPublishActionRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *VideoPublishActionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type VideoPublishActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// @inject_tag: json:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *VideoPublishActionResponse) Reset() {
	*x = VideoPublishActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPublishActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPublishActionResponse) ProtoMessage() {}

func (x *VideoPublishActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPublishActionResponse.ProtoReflect.Descriptor instead.
func (*VideoPublishActionResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{3}
}

func (x *VideoPublishActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VideoPublishActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

//
//PublishList
type VideoPublishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VideoPublishListRequest) Reset() {
	*x = VideoPublishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPublishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPublishListRequest) ProtoMessage() {}

func (x *VideoPublishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPublishListRequest.ProtoReflect.Descriptor instead.
func (*VideoPublishListRequest) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{4}
}

func (x *VideoPublishListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VideoPublishListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VideoPublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// @inject_tag: json:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// @inject_tag: json:"video_list"
	VideoList []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *VideoPublishListResponse) Reset() {
	*x = VideoPublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPublishListResponse) ProtoMessage() {}

func (x *VideoPublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPublishListResponse.ProtoReflect.Descriptor instead.
func (*VideoPublishListResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{5}
}

func (x *VideoPublishListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VideoPublishListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *VideoPublishListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

//
//FavoriteAction
type VideoFavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// @inject_tag: json:"video_id"
	VideoId int64 `protobuf:"varint,3,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	// @inject_tag: json:"action_type"
	ActionType int32 `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (x *VideoFavoriteActionRequest) Reset() {
	*x = VideoFavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFavoriteActionRequest) ProtoMessage() {}

func (x *VideoFavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*VideoFavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{6}
}

func (x *VideoFavoriteActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VideoFavoriteActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *VideoFavoriteActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *VideoFavoriteActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type VideoFavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// @inject_tag: json:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *VideoFavoriteActionResponse) Reset() {
	*x = VideoFavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFavoriteActionResponse) ProtoMessage() {}

func (x *VideoFavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*VideoFavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{7}
}

func (x *VideoFavoriteActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VideoFavoriteActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

//
//FavoriteList
type VideoFavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VideoFavoriteListRequest) Reset() {
	*x = VideoFavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFavoriteListRequest) ProtoMessage() {}

func (x *VideoFavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFavoriteListRequest.ProtoReflect.Descriptor instead.
func (*VideoFavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{8}
}

func (x *VideoFavoriteListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VideoFavoriteListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VideoFavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// @inject_tag: json:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// @inject_tag: json:"video_list"
	VideoList []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *VideoFavoriteListResponse) Reset() {
	*x = VideoFavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFavoriteListResponse) ProtoMessage() {}

func (x *VideoFavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFavoriteListResponse.ProtoReflect.Descriptor instead.
func (*VideoFavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{9}
}

func (x *VideoFavoriteListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VideoFavoriteListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *VideoFavoriteListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_videoService_proto protoreflect.FileDescriptor

var file_videoService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x10, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa0, 0x01, 0x0a,
	0x11, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x2e, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x5b, 0x0a, 0x19, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5c, 0x0a, 0x1a,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x48, 0x0a, 0x17, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x18, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x2e, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x87, 0x01, 0x0a, 0x1a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5d, 0x0a, 0x1b, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x49, 0x0a, 0x18, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0xb9, 0x03, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_videoService_proto_rawDescOnce sync.Once
	file_videoService_proto_rawDescData = file_videoService_proto_rawDesc
)

func file_videoService_proto_rawDescGZIP() []byte {
	file_videoService_proto_rawDescOnce.Do(func() {
		file_videoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_videoService_proto_rawDescData)
	})
	return file_videoService_proto_rawDescData
}

var file_videoService_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_videoService_proto_goTypes = []interface{}{
	(*VideoFeedRequest)(nil),            // 0: services.VideoFeedRequest
	(*VideoFeedResponse)(nil),           // 1: services.VideoFeedResponse
	(*VideoPublishActionRequest)(nil),   // 2: services.VideoPublishActionRequest
	(*VideoPublishActionResponse)(nil),  // 3: services.VideoPublishActionResponse
	(*VideoPublishListRequest)(nil),     // 4: services.VideoPublishListRequest
	(*VideoPublishListResponse)(nil),    // 5: services.VideoPublishListResponse
	(*VideoFavoriteActionRequest)(nil),  // 6: services.VideoFavoriteActionRequest
	(*VideoFavoriteActionResponse)(nil), // 7: services.VideoFavoriteActionResponse
	(*VideoFavoriteListRequest)(nil),    // 8: services.VideoFavoriteListRequest
	(*VideoFavoriteListResponse)(nil),   // 9: services.VideoFavoriteListResponse
	(*Video)(nil),                       // 10: services.Video
}
var file_videoService_proto_depIdxs = []int32{
	10, // 0: services.VideoFeedResponse.video_list:type_name -> services.Video
	10, // 1: services.VideoPublishListResponse.video_list:type_name -> services.Video
	10, // 2: services.VideoFavoriteListResponse.video_list:type_name -> services.Video
	0,  // 3: services.VideoService.Feed:input_type -> services.VideoFeedRequest
	2,  // 4: services.VideoService.PublishAction:input_type -> services.VideoPublishActionRequest
	4,  // 5: services.VideoService.PublishList:input_type -> services.VideoPublishListRequest
	6,  // 6: services.VideoService.FavoriteAction:input_type -> services.VideoFavoriteActionRequest
	8,  // 7: services.VideoService.FavoriteList:input_type -> services.VideoFavoriteListRequest
	1,  // 8: services.VideoService.Feed:output_type -> services.VideoFeedResponse
	3,  // 9: services.VideoService.PublishAction:output_type -> services.VideoPublishActionResponse
	5,  // 10: services.VideoService.PublishList:output_type -> services.VideoPublishListResponse
	7,  // 11: services.VideoService.FavoriteAction:output_type -> services.VideoFavoriteActionResponse
	9,  // 12: services.VideoService.FavoriteList:output_type -> services.VideoFavoriteListResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_videoService_proto_init() }
func file_videoService_proto_init() {
	if File_videoService_proto != nil {
		return
	}
	file_videoModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_videoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFeedRequest); i {
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
		file_videoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFeedResponse); i {
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
		file_videoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPublishActionRequest); i {
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
		file_videoService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPublishActionResponse); i {
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
		file_videoService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPublishListRequest); i {
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
		file_videoService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPublishListResponse); i {
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
		file_videoService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFavoriteActionRequest); i {
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
		file_videoService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFavoriteActionResponse); i {
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
		file_videoService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFavoriteListRequest); i {
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
		file_videoService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFavoriteListResponse); i {
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
			RawDescriptor: file_videoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoService_proto_goTypes,
		DependencyIndexes: file_videoService_proto_depIdxs,
		MessageInfos:      file_videoService_proto_msgTypes,
	}.Build()
	File_videoService_proto = out.File
	file_videoService_proto_rawDesc = nil
	file_videoService_proto_goTypes = nil
	file_videoService_proto_depIdxs = nil
}
