// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: meta.proto

package proto_gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// META //
type RequestMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//  bool mine = 2;
	UserId      int64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TopicId     int64   `protobuf:"varint,4,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	FileId      int64   `protobuf:"varint,5,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	SequenceId  int64   `protobuf:"varint,6,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	Paging      *Paging `protobuf:"bytes,7,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *RequestMeta) Reset() {
	*x = RequestMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMeta) ProtoMessage() {}

func (x *RequestMeta) ProtoReflect() protoreflect.Message {
	mi := &file_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMeta.ProtoReflect.Descriptor instead.
func (*RequestMeta) Descriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMeta) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RequestMeta) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RequestMeta) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RequestMeta) GetTopicId() int64 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *RequestMeta) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *RequestMeta) GetSequenceId() int64 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *RequestMeta) GetPaging() *Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type ResponseMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header        *Header         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Topic         []*Topic        `protobuf:"bytes,2,rep,name=topic,proto3" json:"topic,omitempty"`
	Messages      []*Message      `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	Users         []*User         `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	Profiles      []*Profile      `protobuf:"bytes,5,rep,name=profiles,proto3" json:"profiles,omitempty"`
	Profile       *Profile        `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	Notifications []*Notification `protobuf:"bytes,7,rep,name=notifications,proto3" json:"notifications,omitempty"`
	Result        *Result         `protobuf:"bytes,8,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResponseMeta) Reset() {
	*x = ResponseMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMeta) ProtoMessage() {}

func (x *ResponseMeta) ProtoReflect() protoreflect.Message {
	mi := &file_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMeta.ProtoReflect.Descriptor instead.
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMeta) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ResponseMeta) GetTopic() []*Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *ResponseMeta) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ResponseMeta) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ResponseMeta) GetProfiles() []*Profile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

func (x *ResponseMeta) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *ResponseMeta) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *ResponseMeta) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_meta_proto protoreflect.FileDescriptor

var file_meta_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x80, 0x03, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65,
	0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2e,
	0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x67, 0x65, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xb6, 0x02,
	0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_proto_rawDescOnce sync.Once
	file_meta_proto_rawDescData = file_meta_proto_rawDesc
)

func file_meta_proto_rawDescGZIP() []byte {
	file_meta_proto_rawDescOnce.Do(func() {
		file_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_proto_rawDescData)
	})
	return file_meta_proto_rawDescData
}

var file_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meta_proto_goTypes = []interface{}{
	(*RequestMeta)(nil),  // 0: proto_gen.RequestMeta
	(*ResponseMeta)(nil), // 1: proto_gen.ResponseMeta
	(*Header)(nil),       // 2: proto_gen.Header
	(*Paging)(nil),       // 3: proto_gen.Paging
	(*Topic)(nil),        // 4: proto_gen.Topic
	(*Message)(nil),      // 5: proto_gen.Message
	(*User)(nil),         // 6: proto_gen.User
	(*Profile)(nil),      // 7: proto_gen.Profile
	(*Notification)(nil), // 8: proto_gen.Notification
	(*Result)(nil),       // 9: proto_gen.Result
}
var file_meta_proto_depIdxs = []int32{
	2,  // 0: proto_gen.RequestMeta.header:type_name -> proto_gen.Header
	3,  // 1: proto_gen.RequestMeta.paging:type_name -> proto_gen.Paging
	2,  // 2: proto_gen.ResponseMeta.header:type_name -> proto_gen.Header
	4,  // 3: proto_gen.ResponseMeta.topic:type_name -> proto_gen.Topic
	5,  // 4: proto_gen.ResponseMeta.messages:type_name -> proto_gen.Message
	6,  // 5: proto_gen.ResponseMeta.users:type_name -> proto_gen.User
	7,  // 6: proto_gen.ResponseMeta.profiles:type_name -> proto_gen.Profile
	7,  // 7: proto_gen.ResponseMeta.profile:type_name -> proto_gen.Profile
	8,  // 8: proto_gen.ResponseMeta.notifications:type_name -> proto_gen.Notification
	9,  // 9: proto_gen.ResponseMeta.result:type_name -> proto_gen.Result
	0,  // 10: proto_gen.Meta.Topic:input_type -> proto_gen.RequestMeta
	0,  // 11: proto_gen.Meta.Message:input_type -> proto_gen.RequestMeta
	0,  // 12: proto_gen.Meta.User:input_type -> proto_gen.RequestMeta
	0,  // 13: proto_gen.Meta.Profile:input_type -> proto_gen.RequestMeta
	0,  // 14: proto_gen.Meta.Notify:input_type -> proto_gen.RequestMeta
	1,  // 15: proto_gen.Meta.Topic:output_type -> proto_gen.ResponseMeta
	1,  // 16: proto_gen.Meta.Message:output_type -> proto_gen.ResponseMeta
	1,  // 17: proto_gen.Meta.User:output_type -> proto_gen.ResponseMeta
	1,  // 18: proto_gen.Meta.Profile:output_type -> proto_gen.ResponseMeta
	1,  // 19: proto_gen.Meta.Notify:output_type -> proto_gen.ResponseMeta
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_meta_proto_init() }
func file_meta_proto_init() {
	if File_meta_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMeta); i {
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
		file_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMeta); i {
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
			RawDescriptor: file_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meta_proto_goTypes,
		DependencyIndexes: file_meta_proto_depIdxs,
		MessageInfos:      file_meta_proto_msgTypes,
	}.Build()
	File_meta_proto = out.File
	file_meta_proto_rawDesc = nil
	file_meta_proto_goTypes = nil
	file_meta_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetaClient is the client API for Meta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetaClient interface {
	Topic(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error)
	Message(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error)
	User(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error)
	Profile(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error)
	Notify(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error)
}

type metaClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaClient(cc grpc.ClientConnInterface) MetaClient {
	return &metaClient{cc}
}

func (c *metaClient) Topic(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error) {
	out := new(ResponseMeta)
	err := c.cc.Invoke(ctx, "/proto_gen.Meta/Topic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Message(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error) {
	out := new(ResponseMeta)
	err := c.cc.Invoke(ctx, "/proto_gen.Meta/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) User(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error) {
	out := new(ResponseMeta)
	err := c.cc.Invoke(ctx, "/proto_gen.Meta/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Profile(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error) {
	out := new(ResponseMeta)
	err := c.cc.Invoke(ctx, "/proto_gen.Meta/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Notify(ctx context.Context, in *RequestMeta, opts ...grpc.CallOption) (*ResponseMeta, error) {
	out := new(ResponseMeta)
	err := c.cc.Invoke(ctx, "/proto_gen.Meta/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServer is the server API for Meta service.
type MetaServer interface {
	Topic(context.Context, *RequestMeta) (*ResponseMeta, error)
	Message(context.Context, *RequestMeta) (*ResponseMeta, error)
	User(context.Context, *RequestMeta) (*ResponseMeta, error)
	Profile(context.Context, *RequestMeta) (*ResponseMeta, error)
	Notify(context.Context, *RequestMeta) (*ResponseMeta, error)
}

// UnimplementedMetaServer can be embedded to have forward compatible implementations.
type UnimplementedMetaServer struct {
}

func (*UnimplementedMetaServer) Topic(context.Context, *RequestMeta) (*ResponseMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topic not implemented")
}
func (*UnimplementedMetaServer) Message(context.Context, *RequestMeta) (*ResponseMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (*UnimplementedMetaServer) User(context.Context, *RequestMeta) (*ResponseMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (*UnimplementedMetaServer) Profile(context.Context, *RequestMeta) (*ResponseMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedMetaServer) Notify(context.Context, *RequestMeta) (*ResponseMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}

func RegisterMetaServer(s *grpc.Server, srv MetaServer) {
	s.RegisterService(&_Meta_serviceDesc, srv)
}

func _Meta_Topic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Topic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.Meta/Topic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Topic(ctx, req.(*RequestMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.Meta/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Message(ctx, req.(*RequestMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.Meta/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).User(ctx, req.(*RequestMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.Meta/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Profile(ctx, req.(*RequestMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.Meta/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Notify(ctx, req.(*RequestMeta))
	}
	return interceptor(ctx, in, info, handler)
}

var _Meta_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_gen.Meta",
	HandlerType: (*MetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topic",
			Handler:    _Meta_Topic_Handler,
		},
		{
			MethodName: "Message",
			Handler:    _Meta_Message_Handler,
		},
		{
			MethodName: "User",
			Handler:    _Meta_User_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _Meta_Profile_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Meta_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta.proto",
}
