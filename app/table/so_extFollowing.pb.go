// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extFollowing.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoExtFollowing struct {
	FollowingInfo        *prototype.FollowingCreatedOrder `protobuf:"bytes,1,opt,name=following_info,json=followingInfo,proto3" json:"following_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SoExtFollowing) Reset()         { *m = SoExtFollowing{} }
func (m *SoExtFollowing) String() string { return proto.CompactTextString(m) }
func (*SoExtFollowing) ProtoMessage()    {}
func (*SoExtFollowing) Descriptor() ([]byte, []int) {
	return fileDescriptor_fad976f27ab4b609, []int{0}
}

func (m *SoExtFollowing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtFollowing.Unmarshal(m, b)
}
func (m *SoExtFollowing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtFollowing.Marshal(b, m, deterministic)
}
func (m *SoExtFollowing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtFollowing.Merge(m, src)
}
func (m *SoExtFollowing) XXX_Size() int {
	return xxx_messageInfo_SoExtFollowing.Size(m)
}
func (m *SoExtFollowing) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtFollowing.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtFollowing proto.InternalMessageInfo

func (m *SoExtFollowing) GetFollowingInfo() *prototype.FollowingCreatedOrder {
	if m != nil {
		return m.FollowingInfo
	}
	return nil
}

type SoListExtFollowingByFollowingInfo struct {
	FollowingInfo        *prototype.FollowingCreatedOrder `protobuf:"bytes,1,opt,name=following_info,json=followingInfo,proto3" json:"following_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SoListExtFollowingByFollowingInfo) Reset()         { *m = SoListExtFollowingByFollowingInfo{} }
func (m *SoListExtFollowingByFollowingInfo) String() string { return proto.CompactTextString(m) }
func (*SoListExtFollowingByFollowingInfo) ProtoMessage()    {}
func (*SoListExtFollowingByFollowingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fad976f27ab4b609, []int{1}
}

func (m *SoListExtFollowingByFollowingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtFollowingByFollowingInfo.Unmarshal(m, b)
}
func (m *SoListExtFollowingByFollowingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtFollowingByFollowingInfo.Marshal(b, m, deterministic)
}
func (m *SoListExtFollowingByFollowingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtFollowingByFollowingInfo.Merge(m, src)
}
func (m *SoListExtFollowingByFollowingInfo) XXX_Size() int {
	return xxx_messageInfo_SoListExtFollowingByFollowingInfo.Size(m)
}
func (m *SoListExtFollowingByFollowingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtFollowingByFollowingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtFollowingByFollowingInfo proto.InternalMessageInfo

func (m *SoListExtFollowingByFollowingInfo) GetFollowingInfo() *prototype.FollowingCreatedOrder {
	if m != nil {
		return m.FollowingInfo
	}
	return nil
}

type SoUniqueExtFollowingByFollowingInfo struct {
	FollowingInfo        *prototype.FollowingCreatedOrder `protobuf:"bytes,1,opt,name=following_info,json=followingInfo,proto3" json:"following_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SoUniqueExtFollowingByFollowingInfo) Reset()         { *m = SoUniqueExtFollowingByFollowingInfo{} }
func (m *SoUniqueExtFollowingByFollowingInfo) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtFollowingByFollowingInfo) ProtoMessage()    {}
func (*SoUniqueExtFollowingByFollowingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fad976f27ab4b609, []int{2}
}

func (m *SoUniqueExtFollowingByFollowingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtFollowingByFollowingInfo.Unmarshal(m, b)
}
func (m *SoUniqueExtFollowingByFollowingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtFollowingByFollowingInfo.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtFollowingByFollowingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtFollowingByFollowingInfo.Merge(m, src)
}
func (m *SoUniqueExtFollowingByFollowingInfo) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtFollowingByFollowingInfo.Size(m)
}
func (m *SoUniqueExtFollowingByFollowingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtFollowingByFollowingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtFollowingByFollowingInfo proto.InternalMessageInfo

func (m *SoUniqueExtFollowingByFollowingInfo) GetFollowingInfo() *prototype.FollowingCreatedOrder {
	if m != nil {
		return m.FollowingInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtFollowing)(nil), "table.so_extFollowing")
	proto.RegisterType((*SoListExtFollowingByFollowingInfo)(nil), "table.so_list_extFollowing_by_following_info")
	proto.RegisterType((*SoUniqueExtFollowingByFollowingInfo)(nil), "table.so_unique_extFollowing_by_following_info")
}

func init() { proto.RegisterFile("app/table/so_extFollowing.proto", fileDescriptor_fad976f27ab4b609) }

var fileDescriptor_fad976f27ab4b609 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0x28, 0xd0,
	0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0xce, 0x8f, 0x4f, 0xad, 0x28, 0x71, 0xcb, 0xcf, 0xc9,
	0xc9, 0x2f, 0xcf, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x4b, 0x4a,
	0x89, 0x80, 0x79, 0x25, 0x95, 0x05, 0xa9, 0xfa, 0x20, 0x02, 0x22, 0xa9, 0x14, 0xc3, 0xc5, 0x8f,
	0xa6, 0x4b, 0xc8, 0x93, 0x8b, 0x2f, 0x0d, 0xc6, 0x89, 0xcf, 0xcc, 0x4b, 0xcb, 0x97, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd2, 0x83, 0x9b, 0xa0, 0x87, 0x50, 0x90, 0x5c, 0x94, 0x9a, 0x58,
	0x92, 0x9a, 0x12, 0x9f, 0x5f, 0x94, 0x92, 0x5a, 0x14, 0xc4, 0x0b, 0x97, 0xf0, 0xcc, 0x4b, 0xcb,
	0x57, 0x2a, 0xe6, 0x52, 0x2b, 0xce, 0x8f, 0xcf, 0xc9, 0x2c, 0x2e, 0x41, 0xb1, 0x22, 0x3e, 0xa9,
	0x32, 0x1e, 0xd5, 0x0a, 0x6a, 0x5a, 0x5a, 0xca, 0xa5, 0x51, 0x9c, 0x1f, 0x5f, 0x9a, 0x97, 0x59,
	0x58, 0x9a, 0x4a, 0x3f, 0x6b, 0x9d, 0x34, 0xa2, 0xd4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x93, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0x93, 0xf3, 0xf3,
	0x4a, 0x52, 0xf3, 0x4a, 0xf2, 0x8b, 0x75, 0xd3, 0xf3, 0x21, 0xd1, 0x94, 0xc4, 0x06, 0x36, 0xdb,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x87, 0x6b, 0xf4, 0xba, 0x01, 0x00, 0x00,
}