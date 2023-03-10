// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobufs/auth.proto

package auth

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

type AuthenticationRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85aef938c2975bdf, []int{0}
}

func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(m, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticationResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TokenValid           bool     `protobuf:"varint,2,opt,name=token_valid,json=tokenValid,proto3" json:"token_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationResponse) Reset()         { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()    {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85aef938c2975bdf, []int{1}
}

func (m *AuthenticationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResponse.Unmarshal(m, b)
}
func (m *AuthenticationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResponse.Merge(m, src)
}
func (m *AuthenticationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResponse.Size(m)
}
func (m *AuthenticationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResponse proto.InternalMessageInfo

func (m *AuthenticationResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AuthenticationResponse) GetTokenValid() bool {
	if m != nil {
		return m.TokenValid
	}
	return false
}

func init() {
	proto.RegisterType((*AuthenticationRequest)(nil), "auth.AuthenticationRequest")
	proto.RegisterType((*AuthenticationResponse)(nil), "auth.AuthenticationResponse")
}

func init() {
	proto.RegisterFile("protobufs/auth.proto", fileDescriptor_85aef938c2975bdf)
}

var fileDescriptor_85aef938c2975bdf = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0x2b, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x03, 0x73, 0x85, 0x58, 0x40,
	0x6c, 0x25, 0x5d, 0x2e, 0x51, 0xc7, 0xd2, 0x92, 0x8c, 0xd4, 0xbc, 0x92, 0xcc, 0xe4, 0xc4, 0x92,
	0xcc, 0xfc, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xfc,
	0xec, 0xd4, 0x3c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x29, 0x88, 0x4b, 0x0c,
	0x5d, 0x79, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x38, 0x17, 0x7b, 0x69, 0x71, 0x6a, 0x51,
	0x7c, 0x66, 0x0a, 0x58, 0x07, 0x73, 0x10, 0x1b, 0x88, 0xeb, 0x99, 0x22, 0x24, 0xcf, 0xc5, 0x0d,
	0xd6, 0x1b, 0x5f, 0x96, 0x98, 0x93, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x05,
	0x16, 0x0a, 0x03, 0x89, 0x18, 0x05, 0x72, 0xb1, 0x80, 0xcc, 0x14, 0xf2, 0xe4, 0xe2, 0x41, 0x32,
	0x3b, 0x55, 0x48, 0x5a, 0x0f, 0xec, 0x5a, 0xac, 0xce, 0x93, 0x92, 0xc1, 0x2e, 0x09, 0x71, 0x8c,
	0x93, 0x60, 0x14, 0x3f, 0x58, 0xba, 0xb4, 0x24, 0x33, 0x07, 0xe2, 0xe9, 0x24, 0x36, 0xb0, 0xaf,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x20, 0x27, 0x54, 0x0d, 0x01, 0x00, 0x00,
}
