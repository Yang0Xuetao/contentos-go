// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_transactionObject.proto

package table // import "github.com/coschain/contentos-go/app/table"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototype "github.com/coschain/contentos-go/prototype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoTransactionObject struct {
	TrxId                *prototype.Sha256       `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	Expiration           *prototype.TimePointSec `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoTransactionObject) Reset()         { *m = SoTransactionObject{} }
func (m *SoTransactionObject) String() string { return proto.CompactTextString(m) }
func (*SoTransactionObject) ProtoMessage()    {}
func (*SoTransactionObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_transactionObject_edaab3d12ceec50d, []int{0}
}
func (m *SoTransactionObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoTransactionObject.Unmarshal(m, b)
}
func (m *SoTransactionObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoTransactionObject.Marshal(b, m, deterministic)
}
func (dst *SoTransactionObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoTransactionObject.Merge(dst, src)
}
func (m *SoTransactionObject) XXX_Size() int {
	return xxx_messageInfo_SoTransactionObject.Size(m)
}
func (m *SoTransactionObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SoTransactionObject.DiscardUnknown(m)
}

var xxx_messageInfo_SoTransactionObject proto.InternalMessageInfo

func (m *SoTransactionObject) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func (m *SoTransactionObject) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type SoMemTransactionObjectByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoMemTransactionObjectByTrxId) Reset()         { *m = SoMemTransactionObjectByTrxId{} }
func (m *SoMemTransactionObjectByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoMemTransactionObjectByTrxId) ProtoMessage()    {}
func (*SoMemTransactionObjectByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_transactionObject_edaab3d12ceec50d, []int{1}
}
func (m *SoMemTransactionObjectByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemTransactionObjectByTrxId.Unmarshal(m, b)
}
func (m *SoMemTransactionObjectByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemTransactionObjectByTrxId.Marshal(b, m, deterministic)
}
func (dst *SoMemTransactionObjectByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemTransactionObjectByTrxId.Merge(dst, src)
}
func (m *SoMemTransactionObjectByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoMemTransactionObjectByTrxId.Size(m)
}
func (m *SoMemTransactionObjectByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemTransactionObjectByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemTransactionObjectByTrxId proto.InternalMessageInfo

func (m *SoMemTransactionObjectByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoMemTransactionObjectByExpiration struct {
	Expiration           *prototype.TimePointSec `protobuf:"bytes,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemTransactionObjectByExpiration) Reset()         { *m = SoMemTransactionObjectByExpiration{} }
func (m *SoMemTransactionObjectByExpiration) String() string { return proto.CompactTextString(m) }
func (*SoMemTransactionObjectByExpiration) ProtoMessage()    {}
func (*SoMemTransactionObjectByExpiration) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_transactionObject_edaab3d12ceec50d, []int{2}
}
func (m *SoMemTransactionObjectByExpiration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemTransactionObjectByExpiration.Unmarshal(m, b)
}
func (m *SoMemTransactionObjectByExpiration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemTransactionObjectByExpiration.Marshal(b, m, deterministic)
}
func (dst *SoMemTransactionObjectByExpiration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemTransactionObjectByExpiration.Merge(dst, src)
}
func (m *SoMemTransactionObjectByExpiration) XXX_Size() int {
	return xxx_messageInfo_SoMemTransactionObjectByExpiration.Size(m)
}
func (m *SoMemTransactionObjectByExpiration) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemTransactionObjectByExpiration.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemTransactionObjectByExpiration proto.InternalMessageInfo

func (m *SoMemTransactionObjectByExpiration) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type SoListTransactionObjectByExpiration struct {
	Expiration           *prototype.TimePointSec `protobuf:"bytes,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
	TrxId                *prototype.Sha256       `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListTransactionObjectByExpiration) Reset()         { *m = SoListTransactionObjectByExpiration{} }
func (m *SoListTransactionObjectByExpiration) String() string { return proto.CompactTextString(m) }
func (*SoListTransactionObjectByExpiration) ProtoMessage()    {}
func (*SoListTransactionObjectByExpiration) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_transactionObject_edaab3d12ceec50d, []int{3}
}
func (m *SoListTransactionObjectByExpiration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Unmarshal(m, b)
}
func (m *SoListTransactionObjectByExpiration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Marshal(b, m, deterministic)
}
func (dst *SoListTransactionObjectByExpiration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListTransactionObjectByExpiration.Merge(dst, src)
}
func (m *SoListTransactionObjectByExpiration) XXX_Size() int {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Size(m)
}
func (m *SoListTransactionObjectByExpiration) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListTransactionObjectByExpiration.DiscardUnknown(m)
}

var xxx_messageInfo_SoListTransactionObjectByExpiration proto.InternalMessageInfo

func (m *SoListTransactionObjectByExpiration) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *SoListTransactionObjectByExpiration) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoUniqueTransactionObjectByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoUniqueTransactionObjectByTrxId) Reset()         { *m = SoUniqueTransactionObjectByTrxId{} }
func (m *SoUniqueTransactionObjectByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueTransactionObjectByTrxId) ProtoMessage()    {}
func (*SoUniqueTransactionObjectByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_transactionObject_edaab3d12ceec50d, []int{4}
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Unmarshal(m, b)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Marshal(b, m, deterministic)
}
func (dst *SoUniqueTransactionObjectByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Merge(dst, src)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Size(m)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueTransactionObjectByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueTransactionObjectByTrxId proto.InternalMessageInfo

func (m *SoUniqueTransactionObjectByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func init() {
	proto.RegisterType((*SoTransactionObject)(nil), "table.so_transactionObject")
	proto.RegisterType((*SoMemTransactionObjectByTrxId)(nil), "table.so_mem_transactionObject_by_trx_id")
	proto.RegisterType((*SoMemTransactionObjectByExpiration)(nil), "table.so_mem_transactionObject_by_expiration")
	proto.RegisterType((*SoListTransactionObjectByExpiration)(nil), "table.so_list_transactionObject_by_expiration")
	proto.RegisterType((*SoUniqueTransactionObjectByTrxId)(nil), "table.so_unique_transactionObject_by_trx_id")
}

func init() {
	proto.RegisterFile("app/table/so_transactionObject.proto", fileDescriptor_so_transactionObject_edaab3d12ceec50d)
}

var fileDescriptor_so_transactionObject_edaab3d12ceec50d = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x42, 0x7b, 0x88, 0x27, 0x97, 0x1e, 0xaa, 0x27, 0x59, 0xfc, 0x28, 0xa2, 0x1b,
	0xa8, 0x28, 0x78, 0xf5, 0xe6, 0x45, 0xb1, 0x47, 0x2f, 0x21, 0x9b, 0x0e, 0xdd, 0x91, 0x6e, 0x26,
	0x66, 0x66, 0x61, 0x8b, 0xff, 0xc1, 0xdf, 0x2c, 0xdd, 0x42, 0x5d, 0x54, 0xfc, 0x00, 0x7b, 0x09,
	0x84, 0x3c, 0xef, 0xcc, 0x33, 0x61, 0xd4, 0xa1, 0x0d, 0x41, 0x8b, 0x2d, 0x16, 0xa0, 0x99, 0x8c,
	0x44, 0xeb, 0xd9, 0x3a, 0x41, 0xf2, 0xf7, 0xc5, 0x13, 0x38, 0xc9, 0x43, 0x24, 0xa1, 0xb4, 0xdf,
	0x12, 0xfb, 0xc3, 0xf6, 0x26, 0xcb, 0x00, 0x7a, 0x75, 0xac, 0x1f, 0xb3, 0x17, 0x35, 0xfc, 0x2a,
	0x9a, 0x8e, 0xd5, 0x40, 0x62, 0x63, 0x70, 0x36, 0x4a, 0x0e, 0x92, 0xf1, 0xce, 0x64, 0x37, 0xdf,
	0xc4, 0x73, 0x2e, 0xed, 0xe4, 0xf2, 0x6a, 0xda, 0x97, 0xd8, 0xdc, 0xce, 0xd2, 0x6b, 0xa5, 0xa0,
	0x09, 0x18, 0xed, 0x2a, 0x3d, 0xea, 0xb5, 0xf4, 0x5e, 0x87, 0x16, 0xac, 0xc0, 0x04, 0x42, 0x2f,
	0x86, 0xc1, 0x4d, 0x3b, 0x70, 0x76, 0xa7, 0x32, 0x26, 0x53, 0x41, 0xf5, 0x59, 0xc0, 0x14, 0x4b,
	0xb3, 0x16, 0xf8, 0xbd, 0x4a, 0xe6, 0xd4, 0xf1, 0x77, 0xf5, 0xde, 0x3b, 0x7f, 0x90, 0x4e, 0xfe,
	0x22, 0xfd, 0x9a, 0xa8, 0x13, 0x26, 0xb3, 0x40, 0x96, 0x2d, 0xb6, 0xe9, 0x4c, 0xdd, 0xfb, 0x61,
	0xea, 0x07, 0x75, 0xc4, 0x64, 0x6a, 0x8f, 0xcf, 0x35, 0xfc, 0xcf, 0x47, 0xde, 0x9c, 0x3d, 0x9e,
	0xce, 0x51, 0xca, 0xba, 0xc8, 0x1d, 0x55, 0xda, 0x11, 0xbb, 0xd2, 0xa2, 0xd7, 0x8e, 0xbc, 0x80,
	0x17, 0xe2, 0xf3, 0x39, 0xe9, 0xcd, 0xee, 0x15, 0x83, 0xb6, 0xcc, 0xc5, 0x5b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0x40, 0x9e, 0x41, 0x8f, 0x02, 0x00, 0x00,
}
