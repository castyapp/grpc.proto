// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: ws.enums.proto

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

type PERSONAL_STATE int32

const (
	PERSONAL_STATE_OFFLINE   PERSONAL_STATE = 0
	PERSONAL_STATE_ONLINE    PERSONAL_STATE = 1
	PERSONAL_STATE_IDLE      PERSONAL_STATE = 2
	PERSONAL_STATE_BUSY      PERSONAL_STATE = 3
	PERSONAL_STATE_INVISIBLE PERSONAL_STATE = 4
)

// Enum value maps for PERSONAL_STATE.
var (
	PERSONAL_STATE_name = map[int32]string{
		0: "OFFLINE",
		1: "ONLINE",
		2: "IDLE",
		3: "BUSY",
		4: "INVISIBLE",
	}
	PERSONAL_STATE_value = map[string]int32{
		"OFFLINE":   0,
		"ONLINE":    1,
		"IDLE":      2,
		"BUSY":      3,
		"INVISIBLE": 4,
	}
)

func (x PERSONAL_STATE) Enum() *PERSONAL_STATE {
	p := new(PERSONAL_STATE)
	*p = x
	return p
}

func (x PERSONAL_STATE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PERSONAL_STATE) Descriptor() protoreflect.EnumDescriptor {
	return file_ws_enums_proto_enumTypes[0].Descriptor()
}

func (PERSONAL_STATE) Type() protoreflect.EnumType {
	return &file_ws_enums_proto_enumTypes[0]
}

func (x PERSONAL_STATE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PERSONAL_STATE.Descriptor instead.
func (PERSONAL_STATE) EnumDescriptor() ([]byte, []int) {
	return file_ws_enums_proto_rawDescGZIP(), []int{0}
}

type EMSG int32

const (
	EMSG_INVALID                   EMSG = 0
	EMSG_PING                      EMSG = 1
	EMSG_PONG                      EMSG = 2
	EMSG_LOGON                     EMSG = 3
	EMSG_LOGOUT                    EMSG = 4
	EMSG_PERSONAL_STATE_CHANGED    EMSG = 5
	EMSG_PERSONAL_ACTIVITY_CHANGED EMSG = 6
	EMSG_AUTHORIZED                EMSG = 7
	EMSG_UNAUTHORIZED              EMSG = 8
	EMSG_CHAT_MESSAGES             EMSG = 9
	EMSG_LOG_MESSAGES              EMSG = 10
	EMSG_NEW_CHAT_MESSAGE          EMSG = 11
	EMSG_NEW_LOG_MESSAGE           EMSG = 12
	EMSG_MEMBER_STATE_CHANGED      EMSG = 13
	EMSG_THEATER_MEMBERS           EMSG = 14
	EMSG_THEATER_PLAY              EMSG = 15
	EMSG_THEATER_PAUSE             EMSG = 16
	EMSG_NEW_NOTIFICATION          EMSG = 17
	EMSG_FRIEND_REQUEST_ACCEPTED   EMSG = 18
	// Theater video player events
	EMSG_SYNC_ME             EMSG = 19
	EMSG_SYNCED              EMSG = 20
	EMSG_BUFFERING           EMSG = 21
	EMSG_BUFFERED            EMSG = 22
	EMSG_WAITING_FOR_CLIENTS EMSG = 23
	EMSG_CLIENTS_SYNCYED     EMSG = 24
	EMSG_FINISHED_MOVIE      EMSG = 25
	EMSG_PLAYING             EMSG = 26
	EMSG_CLIENT_READY        EMSG = 27
)

// Enum value maps for EMSG.
var (
	EMSG_name = map[int32]string{
		0:  "INVALID",
		1:  "PING",
		2:  "PONG",
		3:  "LOGON",
		4:  "LOGOUT",
		5:  "PERSONAL_STATE_CHANGED",
		6:  "PERSONAL_ACTIVITY_CHANGED",
		7:  "AUTHORIZED",
		8:  "UNAUTHORIZED",
		9:  "CHAT_MESSAGES",
		10: "LOG_MESSAGES",
		11: "NEW_CHAT_MESSAGE",
		12: "NEW_LOG_MESSAGE",
		13: "MEMBER_STATE_CHANGED",
		14: "THEATER_MEMBERS",
		15: "THEATER_PLAY",
		16: "THEATER_PAUSE",
		17: "NEW_NOTIFICATION",
		18: "FRIEND_REQUEST_ACCEPTED",
		19: "SYNC_ME",
		20: "SYNCED",
		21: "BUFFERING",
		22: "BUFFERED",
		23: "WAITING_FOR_CLIENTS",
		24: "CLIENTS_SYNCYED",
		25: "FINISHED_MOVIE",
		26: "PLAYING",
		27: "CLIENT_READY",
	}
	EMSG_value = map[string]int32{
		"INVALID":                   0,
		"PING":                      1,
		"PONG":                      2,
		"LOGON":                     3,
		"LOGOUT":                    4,
		"PERSONAL_STATE_CHANGED":    5,
		"PERSONAL_ACTIVITY_CHANGED": 6,
		"AUTHORIZED":                7,
		"UNAUTHORIZED":              8,
		"CHAT_MESSAGES":             9,
		"LOG_MESSAGES":              10,
		"NEW_CHAT_MESSAGE":          11,
		"NEW_LOG_MESSAGE":           12,
		"MEMBER_STATE_CHANGED":      13,
		"THEATER_MEMBERS":           14,
		"THEATER_PLAY":              15,
		"THEATER_PAUSE":             16,
		"NEW_NOTIFICATION":          17,
		"FRIEND_REQUEST_ACCEPTED":   18,
		"SYNC_ME":                   19,
		"SYNCED":                    20,
		"BUFFERING":                 21,
		"BUFFERED":                  22,
		"WAITING_FOR_CLIENTS":       23,
		"CLIENTS_SYNCYED":           24,
		"FINISHED_MOVIE":            25,
		"PLAYING":                   26,
		"CLIENT_READY":              27,
	}
)

func (x EMSG) Enum() *EMSG {
	p := new(EMSG)
	*p = x
	return p
}

func (x EMSG) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EMSG) Descriptor() protoreflect.EnumDescriptor {
	return file_ws_enums_proto_enumTypes[1].Descriptor()
}

func (EMSG) Type() protoreflect.EnumType {
	return &file_ws_enums_proto_enumTypes[1]
}

func (x EMSG) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EMSG.Descriptor instead.
func (EMSG) EnumDescriptor() ([]byte, []int) {
	return file_ws_enums_proto_rawDescGZIP(), []int{1}
}

var File_ws_enums_proto protoreflect.FileDescriptor

var file_ws_enums_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4c, 0x0a, 0x0e, 0x50, 0x45, 0x52, 0x53, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x55, 0x53, 0x59, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x56, 0x49, 0x53, 0x49,
	0x42, 0x4c, 0x45, 0x10, 0x04, 0x2a, 0x89, 0x04, 0x0a, 0x04, 0x45, 0x4d, 0x53, 0x47, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f,
	0x47, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10,
	0x07, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45,
	0x44, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x53, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4f, 0x47, 0x5f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x53, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x45, 0x57, 0x5f,
	0x43, 0x48, 0x41, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x13,
	0x0a, 0x0f, 0x4e, 0x45, 0x57, 0x5f, 0x4c, 0x4f, 0x47, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x13, 0x0a,
	0x0f, 0x54, 0x48, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53,
	0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x48, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x4c,
	0x41, 0x59, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x48, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x10, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x45, 0x57, 0x5f, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x11, 0x12, 0x1b, 0x0a,
	0x17, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x12, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x59,
	0x4e, 0x43, 0x5f, 0x4d, 0x45, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x4e, 0x43, 0x45,
	0x44, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x15, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52, 0x45, 0x44, 0x10, 0x16,
	0x12, 0x17, 0x0a, 0x13, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x53, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x59, 0x45, 0x44, 0x10, 0x18, 0x12, 0x12,
	0x0a, 0x0e, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x5f, 0x4d, 0x4f, 0x56, 0x49, 0x45,
	0x10, 0x19, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x1a, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x1b, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ws_enums_proto_rawDescOnce sync.Once
	file_ws_enums_proto_rawDescData = file_ws_enums_proto_rawDesc
)

func file_ws_enums_proto_rawDescGZIP() []byte {
	file_ws_enums_proto_rawDescOnce.Do(func() {
		file_ws_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_ws_enums_proto_rawDescData)
	})
	return file_ws_enums_proto_rawDescData
}

var file_ws_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ws_enums_proto_goTypes = []interface{}{
	(PERSONAL_STATE)(0), // 0: proto.PERSONAL_STATE
	(EMSG)(0),           // 1: proto.EMSG
}
var file_ws_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ws_enums_proto_init() }
func file_ws_enums_proto_init() {
	if File_ws_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ws_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ws_enums_proto_goTypes,
		DependencyIndexes: file_ws_enums_proto_depIdxs,
		EnumInfos:         file_ws_enums_proto_enumTypes,
	}.Build()
	File_ws_enums_proto = out.File
	file_ws_enums_proto_rawDesc = nil
	file_ws_enums_proto_goTypes = nil
	file_ws_enums_proto_depIdxs = nil
}
