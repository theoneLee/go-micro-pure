// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello_world.proto

package micro_pure_helloworld

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EchoRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_296ee14c59be48da, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EchoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *EchoRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type EchoResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_296ee14c59be48da, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EchoResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *EchoResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_296ee14c59be48da, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	SayContent           string   `protobuf:"bytes,3,opt,name=sayContent,proto3" json:"sayContent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_296ee14c59be48da, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetSayContent() string {
	if m != nil {
		return m.SayContent
	}
	return ""
}

type PrintRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintRequest) Reset()         { *m = PrintRequest{} }
func (m *PrintRequest) String() string { return proto.CompactTextString(m) }
func (*PrintRequest) ProtoMessage()    {}
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_296ee14c59be48da, []int{4}
}

func (m *PrintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintRequest.Unmarshal(m, b)
}
func (m *PrintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintRequest.Marshal(b, m, deterministic)
}
func (m *PrintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintRequest.Merge(m, src)
}
func (m *PrintRequest) XXX_Size() int {
	return xxx_messageInfo_PrintRequest.Size(m)
}
func (m *PrintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrintRequest proto.InternalMessageInfo

func (m *PrintRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PrintResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	PrintContent         string   `protobuf:"bytes,3,opt,name=printContent,proto3" json:"printContent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintResponse) Reset()         { *m = PrintResponse{} }
func (m *PrintResponse) String() string { return proto.CompactTextString(m) }
func (*PrintResponse) ProtoMessage()    {}
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_296ee14c59be48da, []int{5}
}

func (m *PrintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintResponse.Unmarshal(m, b)
}
func (m *PrintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintResponse.Marshal(b, m, deterministic)
}
func (m *PrintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintResponse.Merge(m, src)
}
func (m *PrintResponse) XXX_Size() int {
	return xxx_messageInfo_PrintResponse.Size(m)
}
func (m *PrintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrintResponse proto.InternalMessageInfo

func (m *PrintResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PrintResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *PrintResponse) GetPrintContent() string {
	if m != nil {
		return m.PrintContent
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "micro.pure.helloworld.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "micro.pure.helloworld.EchoResponse")
	proto.RegisterType((*Error)(nil), "micro.pure.helloworld.Error")
	proto.RegisterType((*User)(nil), "micro.pure.helloworld.user")
	proto.RegisterType((*PrintRequest)(nil), "micro.pure.helloworld.PrintRequest")
	proto.RegisterType((*PrintResponse)(nil), "micro.pure.helloworld.PrintResponse")
}

func init() { proto.RegisterFile("proto/hello_world.proto", fileDescriptor_296ee14c59be48da) }

var fileDescriptor_296ee14c59be48da = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xb6, 0xa9, 0x75, 0x5a, 0x3d, 0x0c, 0xa8, 0xa1, 0x4a, 0x29, 0xab, 0x87, 0x9e,
	0x52, 0x48, 0x1f, 0x41, 0x0a, 0x9e, 0x44, 0x57, 0xc1, 0xa3, 0xd4, 0xcd, 0x42, 0x03, 0x69, 0x36,
	0xee, 0x6e, 0x10, 0x8f, 0x3e, 0x80, 0x8f, 0xe3, 0xfb, 0xc9, 0x4e, 0x5a, 0x6d, 0xc0, 0x14, 0x3c,
	0x78, 0xcb, 0xcc, 0x7c, 0xf3, 0xef, 0xcc, 0x9f, 0x81, 0xd3, 0xc2, 0x68, 0xa7, 0xa7, 0x4b, 0x95,
	0x65, 0xfa, 0xe9, 0x55, 0x9b, 0x2c, 0x89, 0x28, 0x83, 0xc7, 0xab, 0x54, 0x1a, 0x1d, 0x15, 0xa5,
	0x51, 0x11, 0x55, 0xa9, 0xc8, 0xef, 0xa1, 0x3f, 0x97, 0x4b, 0x2d, 0xd4, 0x4b, 0xa9, 0xac, 0xc3,
	0x23, 0x68, 0xa5, 0x49, 0xc8, 0xc6, 0x6c, 0xd2, 0x16, 0xad, 0x34, 0xc1, 0x21, 0xf4, 0x4a, 0xab,
	0xcc, 0xcd, 0x62, 0xa5, 0xc2, 0xd6, 0x98, 0x4d, 0x0e, 0xc4, 0x77, 0x8c, 0x21, 0xec, 0x4b, 0x9d,
	0x3b, 0x95, 0xbb, 0xb0, 0x4d, 0xa5, 0x4d, 0xc8, 0x3f, 0x18, 0x0c, 0x2a, 0x55, 0x5b, 0xe8, 0xdc,
	0x12, 0x6a, 0x4b, 0x29, 0x95, 0xb5, 0xa4, 0xdd, 0x13, 0x9b, 0x10, 0x63, 0x08, 0x94, 0x31, 0xda,
	0x90, 0x7a, 0x3f, 0x3e, 0x8f, 0x7e, 0x1d, 0x33, 0x9a, 0x7b, 0x46, 0x54, 0x28, 0x4e, 0xa1, 0xe3,
	0x87, 0xa0, 0x57, 0xfb, 0xf1, 0x59, 0x43, 0x8b, 0x47, 0x04, 0x81, 0x7c, 0x06, 0x01, 0x09, 0x20,
	0x42, 0x47, 0xea, 0x44, 0xd1, 0x10, 0x81, 0xa0, 0x6f, 0x3c, 0x81, 0x6e, 0xa2, 0xdc, 0x22, 0xcd,
	0xd6, 0x0b, 0xae, 0x23, 0x2e, 0xaa, 0x57, 0xfe, 0x64, 0xc9, 0x08, 0xc0, 0x2e, 0xde, 0xae, 0x6a,
	0xae, 0x6c, 0x65, 0xf8, 0x08, 0x06, 0xb7, 0x26, 0xcd, 0x5d, 0x83, 0xdd, 0xfc, 0x9d, 0xc1, 0xe1,
	0x1a, 0xf8, 0x17, 0xe7, 0x38, 0x0c, 0x0a, 0x2f, 0x5f, 0x9f, 0xb0, 0x96, 0x8b, 0x3f, 0x19, 0xc0,
	0xb5, 0xef, 0x7f, 0xf4, 0xfd, 0x78, 0x07, 0x1d, 0xff, 0x2b, 0x91, 0x37, 0xe9, 0xff, 0x5c, 0xcf,
	0xf0, 0x62, 0x27, 0x53, 0x6d, 0xc4, 0xf7, 0xf0, 0x01, 0x02, 0x5a, 0x12, 0x9b, 0xf8, 0x6d, 0x8f,
	0x86, 0x97, 0xbb, 0xa1, 0x8d, 0xea, 0x73, 0x97, 0xee, 0x7c, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0xa9, 0x8d, 0x4e, 0x02, 0x03, 0x00, 0x00,
}
