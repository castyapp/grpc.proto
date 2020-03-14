// Code generated by protoc-gen-go. DO NOT EDIT.
// source: theater.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PRIVACY int32

const (
	PRIVACY_PRIVATE  PRIVACY = 0
	PRIVACY_EVERYONE PRIVACY = 1
	PRIVACY_PUBLIC   PRIVACY = 2
)

var PRIVACY_name = map[int32]string{
	0: "PRIVATE",
	1: "EVERYONE",
	2: "PUBLIC",
}

var PRIVACY_value = map[string]int32{
	"PRIVATE":  0,
	"EVERYONE": 1,
	"PUBLIC":   2,
}

func (x PRIVACY) String() string {
	return proto.EnumName(PRIVACY_name, int32(x))
}

func (PRIVACY) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ca216068025a679, []int{0}
}

type Subtitle struct {
	Lang                 string   `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	File                 string   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subtitle) Reset()         { *m = Subtitle{} }
func (m *Subtitle) String() string { return proto.CompactTextString(m) }
func (*Subtitle) ProtoMessage()    {}
func (*Subtitle) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ca216068025a679, []int{0}
}

func (m *Subtitle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subtitle.Unmarshal(m, b)
}
func (m *Subtitle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subtitle.Marshal(b, m, deterministic)
}
func (m *Subtitle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subtitle.Merge(m, src)
}
func (m *Subtitle) XXX_Size() int {
	return xxx_messageInfo_Subtitle.Size(m)
}
func (m *Subtitle) XXX_DiscardUnknown() {
	xxx_messageInfo_Subtitle.DiscardUnknown(m)
}

var xxx_messageInfo_Subtitle proto.InternalMessageInfo

func (m *Subtitle) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Subtitle) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

type Movie struct {
	MovieUri             string      `protobuf:"bytes,1,opt,name=movie_uri,json=movieUri,proto3" json:"movie_uri,omitempty"`
	Poster               string      `protobuf:"bytes,2,opt,name=poster,proto3" json:"poster,omitempty"`
	Size                 int64       `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Length               int64       `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	LastPlayedTime       int64       `protobuf:"varint,5,opt,name=last_played_time,json=lastPlayedTime,proto3" json:"last_played_time,omitempty"`
	Subtitles            []*Subtitle `protobuf:"bytes,6,rep,name=subtitles,proto3" json:"subtitles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ca216068025a679, []int{1}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetMovieUri() string {
	if m != nil {
		return m.MovieUri
	}
	return ""
}

func (m *Movie) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *Movie) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Movie) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Movie) GetLastPlayedTime() int64 {
	if m != nil {
		return m.LastPlayedTime
	}
	return 0
}

func (m *Movie) GetSubtitles() []*Subtitle {
	if m != nil {
		return m.Subtitles
	}
	return nil
}

type Theater struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Hash                 string               `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Privacy              PRIVACY              `protobuf:"varint,4,opt,name=privacy,proto3,enum=messages.PRIVACY" json:"privacy,omitempty"`
	VideoPlayerAccess    PRIVACY              `protobuf:"varint,5,opt,name=video_player_access,json=videoPlayerAccess,proto3,enum=messages.PRIVACY" json:"video_player_access,omitempty"`
	UserId               string               `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User                 *User                `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	Movie                *Movie               `protobuf:"bytes,8,opt,name=movie,proto3" json:"movie,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Theater) Reset()         { *m = Theater{} }
func (m *Theater) String() string { return proto.CompactTextString(m) }
func (*Theater) ProtoMessage()    {}
func (*Theater) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ca216068025a679, []int{2}
}

func (m *Theater) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Theater.Unmarshal(m, b)
}
func (m *Theater) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Theater.Marshal(b, m, deterministic)
}
func (m *Theater) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Theater.Merge(m, src)
}
func (m *Theater) XXX_Size() int {
	return xxx_messageInfo_Theater.Size(m)
}
func (m *Theater) XXX_DiscardUnknown() {
	xxx_messageInfo_Theater.DiscardUnknown(m)
}

var xxx_messageInfo_Theater proto.InternalMessageInfo

func (m *Theater) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Theater) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Theater) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Theater) GetPrivacy() PRIVACY {
	if m != nil {
		return m.Privacy
	}
	return PRIVACY_PRIVATE
}

func (m *Theater) GetVideoPlayerAccess() PRIVACY {
	if m != nil {
		return m.VideoPlayerAccess
	}
	return PRIVACY_PRIVATE
}

func (m *Theater) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Theater) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Theater) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

func (m *Theater) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Theater) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("messages.PRIVACY", PRIVACY_name, PRIVACY_value)
	proto.RegisterType((*Subtitle)(nil), "messages.Subtitle")
	proto.RegisterType((*Movie)(nil), "messages.Movie")
	proto.RegisterType((*Theater)(nil), "messages.Theater")
}

func init() { proto.RegisterFile("theater.proto", fileDescriptor_9ca216068025a679) }

var fileDescriptor_9ca216068025a679 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0xeb, 0x64, 0xe2, 0x9f, 0x9b, 0x36, 0xcd, 0xa8, 0xa5, 0x15, 0xe9, 0xa2, 0x21, 0x50,
	0x30, 0x2d, 0x78, 0x86, 0x74, 0xd5, 0xa5, 0x3b, 0x64, 0x11, 0xe8, 0x4f, 0x50, 0x93, 0x81, 0x59,
	0x19, 0x25, 0xd6, 0x38, 0x02, 0x7b, 0x6c, 0x24, 0x39, 0x30, 0x7d, 0x87, 0xbe, 0x55, 0x1f, 0xac,
	0xe8, 0x5a, 0x6e, 0x36, 0x85, 0xee, 0x8e, 0x8e, 0xbf, 0x23, 0xeb, 0x5c, 0x09, 0x9e, 0x99, 0xa3,
	0xe0, 0x46, 0xa8, 0xa4, 0x51, 0xb5, 0xa9, 0x49, 0x58, 0x09, 0xad, 0x79, 0x21, 0xf4, 0xec, 0x6d,
	0x51, 0xd7, 0x45, 0x29, 0xae, 0xd0, 0xdf, 0xb7, 0xf7, 0x57, 0x46, 0x56, 0x42, 0x1b, 0x5e, 0x35,
	0x1d, 0x3a, 0x83, 0x56, 0xf7, 0xb1, 0xc5, 0x12, 0xc2, 0x1f, 0xed, 0xde, 0x48, 0x53, 0x0a, 0x42,
	0xe0, 0xa2, 0xe4, 0x0f, 0x05, 0xf5, 0xe6, 0x5e, 0x1c, 0x31, 0xd4, 0xd6, 0xbb, 0x97, 0xa5, 0xa0,
	0x83, 0xce, 0xb3, 0x7a, 0xf1, 0xdb, 0x83, 0xd1, 0xd7, 0xfa, 0x24, 0x05, 0x79, 0x03, 0x51, 0x65,
	0x45, 0xd6, 0x2a, 0xe9, 0x62, 0x21, 0x1a, 0x3b, 0x25, 0xc9, 0x2b, 0xf0, 0x9b, 0x5a, 0x1b, 0xa1,
	0x5c, 0xd8, 0xad, 0xec, 0x96, 0x5a, 0xfe, 0x14, 0x74, 0x38, 0xf7, 0xe2, 0x21, 0x43, 0x6d, 0xd9,
	0x52, 0x3c, 0x14, 0xe6, 0x48, 0x2f, 0xd0, 0x75, 0x2b, 0x12, 0xc3, 0xb4, 0xe4, 0xda, 0x64, 0x4d,
	0xc9, 0x1f, 0x45, 0x9e, 0xd9, 0x26, 0x74, 0x84, 0xc4, 0xc4, 0xfa, 0x1b, 0xb4, 0xb7, 0xb2, 0x12,
	0xe4, 0x1a, 0x22, 0xed, 0x8a, 0x68, 0xea, 0xcf, 0x87, 0xf1, 0x78, 0x49, 0x92, 0x7e, 0x26, 0x49,
	0xdf, 0x91, 0x9d, 0xa1, 0xc5, 0xaf, 0x21, 0x04, 0xdb, 0x6e, 0x86, 0x64, 0x02, 0x03, 0x99, 0xbb,
	0x06, 0x03, 0x99, 0x93, 0x97, 0x30, 0x42, 0xca, 0x1d, 0x7d, 0xf4, 0x77, 0x40, 0x47, 0xae, 0x8f,
	0x78, 0xf2, 0x88, 0xa1, 0x26, 0x1f, 0x20, 0x68, 0x94, 0x3c, 0xf1, 0xc3, 0x23, 0x1e, 0x7d, 0xb2,
	0xbc, 0x3c, 0xff, 0x75, 0xc3, 0xd6, 0xb7, 0xe9, 0xcd, 0x1d, 0xeb, 0x09, 0x92, 0xc2, 0x8b, 0x93,
	0xcc, 0x45, 0xdd, 0xf5, 0x51, 0x19, 0x3f, 0x1c, 0x84, 0xd6, 0xd8, 0xe8, 0x9f, 0xc1, 0x4b, 0xa4,
	0xb1, 0xa5, 0x4a, 0x91, 0x25, 0xaf, 0x21, 0xb0, 0xd7, 0x97, 0xc9, 0x9c, 0xfa, 0xdd, 0x58, 0xed,
	0x72, 0x9d, 0x93, 0x05, 0x5c, 0x58, 0x45, 0x83, 0xb9, 0x17, 0x8f, 0x97, 0x93, 0xf3, 0x66, 0x3b,
	0x2d, 0x14, 0xc3, 0x6f, 0xe4, 0x1d, 0x8c, 0xf0, 0x7a, 0x68, 0x88, 0xd0, 0xf3, 0x33, 0x84, 0xf7,
	0xc9, 0xba, 0xaf, 0xe4, 0x13, 0xc0, 0x41, 0xd9, 0xc1, 0xe4, 0x19, 0x37, 0x34, 0x42, 0x76, 0x96,
	0x74, 0xcf, 0x2a, 0xe9, 0x9f, 0x55, 0xb2, 0xed, 0x9f, 0x15, 0x8b, 0x1c, 0x9d, 0x1a, 0x1b, 0x6d,
	0x9b, 0xbc, 0x8f, 0xc2, 0xff, 0xa3, 0x8e, 0x4e, 0xcd, 0xfb, 0x6b, 0x08, 0x5c, 0x6f, 0x32, 0x76,
	0x72, 0xbb, 0x9a, 0x3e, 0x21, 0x4f, 0x21, 0x5c, 0xdd, 0xae, 0xd8, 0xdd, 0xf7, 0x6f, 0xab, 0xa9,
	0x47, 0x00, 0xfc, 0xcd, 0xee, 0xf3, 0x97, 0xf5, 0xcd, 0x74, 0xb0, 0xf7, 0x71, 0xc3, 0x8f, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x77, 0x91, 0xb5, 0x08, 0x0b, 0x03, 0x00, 0x00,
}