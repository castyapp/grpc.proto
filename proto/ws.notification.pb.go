// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: ws.notification.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NotificationMsgEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *NotificationMsgEvent) Reset() {
	*x = NotificationMsgEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationMsgEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationMsgEvent) ProtoMessage() {}

func (x *NotificationMsgEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ws_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationMsgEvent.ProtoReflect.Descriptor instead.
func (*NotificationMsgEvent) Descriptor() ([]byte, []int) {
	return file_ws_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationMsgEvent) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type FriendRequestAcceptedMsgEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friend *User `protobuf:"bytes,1,opt,name=friend,proto3" json:"friend,omitempty"`
}

func (x *FriendRequestAcceptedMsgEvent) Reset() {
	*x = FriendRequestAcceptedMsgEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestAcceptedMsgEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestAcceptedMsgEvent) ProtoMessage() {}

func (x *FriendRequestAcceptedMsgEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ws_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestAcceptedMsgEvent.ProtoReflect.Descriptor instead.
func (*FriendRequestAcceptedMsgEvent) Descriptor() ([]byte, []int) {
	return file_ws_notification_proto_rawDescGZIP(), []int{1}
}

func (x *FriendRequestAcceptedMsgEvent) GetFriend() *User {
	if x != nil {
		return x.Friend
	}
	return nil
}

var File_ws_notification_proto protoreflect.FileDescriptor

var file_ws_notification_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x73, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x44, 0x0a, 0x1d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ws_notification_proto_rawDescOnce sync.Once
	file_ws_notification_proto_rawDescData = file_ws_notification_proto_rawDesc
)

func file_ws_notification_proto_rawDescGZIP() []byte {
	file_ws_notification_proto_rawDescOnce.Do(func() {
		file_ws_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_ws_notification_proto_rawDescData)
	})
	return file_ws_notification_proto_rawDescData
}

var file_ws_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ws_notification_proto_goTypes = []interface{}{
	(*NotificationMsgEvent)(nil),          // 0: proto.NotificationMsgEvent
	(*FriendRequestAcceptedMsgEvent)(nil), // 1: proto.FriendRequestAcceptedMsgEvent
	(*Notification)(nil),                  // 2: proto.Notification
	(*User)(nil),                          // 3: proto.User
}
var file_ws_notification_proto_depIdxs = []int32{
	2, // 0: proto.NotificationMsgEvent.notification:type_name -> proto.Notification
	3, // 1: proto.FriendRequestAcceptedMsgEvent.friend:type_name -> proto.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ws_notification_proto_init() }
func file_ws_notification_proto_init() {
	if File_ws_notification_proto != nil {
		return
	}
	file_grpc_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ws_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationMsgEvent); i {
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
		file_ws_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestAcceptedMsgEvent); i {
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
			RawDescriptor: file_ws_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ws_notification_proto_goTypes,
		DependencyIndexes: file_ws_notification_proto_depIdxs,
		MessageInfos:      file_ws_notification_proto_msgTypes,
	}.Build()
	File_ws_notification_proto = out.File
	file_ws_notification_proto_rawDesc = nil
	file_ws_notification_proto_goTypes = nil
	file_ws_notification_proto_depIdxs = nil
}
