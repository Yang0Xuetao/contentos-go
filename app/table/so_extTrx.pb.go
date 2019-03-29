// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extTrx.proto

package table

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

type SoExtTrx struct {
	TrxId                *prototype.Sha256             `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	BlockHeight          uint64                        `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	TrxWrap              *prototype.TransactionWrapper `protobuf:"bytes,3,opt,name=trx_wrap,json=trxWrap,proto3" json:"trx_wrap,omitempty"`
	BlockTime            *prototype.TimePointSec       `protobuf:"bytes,4,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	TrxCreateOrder       *prototype.UserTrxCreateOrder `protobuf:"bytes,5,opt,name=trx_create_order,json=trxCreateOrder,proto3" json:"trx_create_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoExtTrx) Reset()         { *m = SoExtTrx{} }
func (m *SoExtTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtTrx) ProtoMessage()    {}
func (*SoExtTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{0}
}

func (m *SoExtTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtTrx.Unmarshal(m, b)
}
func (m *SoExtTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtTrx.Marshal(b, m, deterministic)
}
func (m *SoExtTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtTrx.Merge(m, src)
}
func (m *SoExtTrx) XXX_Size() int {
	return xxx_messageInfo_SoExtTrx.Size(m)
}
func (m *SoExtTrx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtTrx.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtTrx proto.InternalMessageInfo

func (m *SoExtTrx) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func (m *SoExtTrx) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SoExtTrx) GetTrxWrap() *prototype.TransactionWrapper {
	if m != nil {
		return m.TrxWrap
	}
	return nil
}

func (m *SoExtTrx) GetBlockTime() *prototype.TimePointSec {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

func (m *SoExtTrx) GetTrxCreateOrder() *prototype.UserTrxCreateOrder {
	if m != nil {
		return m.TrxCreateOrder
	}
	return nil
}

type SoMemExtTrxByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoMemExtTrxByTrxId) Reset()         { *m = SoMemExtTrxByTrxId{} }
func (m *SoMemExtTrxByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoMemExtTrxByTrxId) ProtoMessage()    {}
func (*SoMemExtTrxByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{1}
}

func (m *SoMemExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoMemExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (m *SoMemExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByTrxId.Merge(m, src)
}
func (m *SoMemExtTrxByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoMemExtTrxByTrxId.Size(m)
}
func (m *SoMemExtTrxByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtTrxByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtTrxByTrxId proto.InternalMessageInfo

func (m *SoMemExtTrxByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoMemExtTrxByBlockHeight struct {
	BlockHeight          uint64   `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtTrxByBlockHeight) Reset()         { *m = SoMemExtTrxByBlockHeight{} }
func (m *SoMemExtTrxByBlockHeight) String() string { return proto.CompactTextString(m) }
func (*SoMemExtTrxByBlockHeight) ProtoMessage()    {}
func (*SoMemExtTrxByBlockHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{2}
}

func (m *SoMemExtTrxByBlockHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByBlockHeight.Unmarshal(m, b)
}
func (m *SoMemExtTrxByBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByBlockHeight.Marshal(b, m, deterministic)
}
func (m *SoMemExtTrxByBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByBlockHeight.Merge(m, src)
}
func (m *SoMemExtTrxByBlockHeight) XXX_Size() int {
	return xxx_messageInfo_SoMemExtTrxByBlockHeight.Size(m)
}
func (m *SoMemExtTrxByBlockHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtTrxByBlockHeight.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtTrxByBlockHeight proto.InternalMessageInfo

func (m *SoMemExtTrxByBlockHeight) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type SoMemExtTrxByTrxWrap struct {
	TrxWrap              *prototype.TransactionWrapper `protobuf:"bytes,1,opt,name=trx_wrap,json=trxWrap,proto3" json:"trx_wrap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoMemExtTrxByTrxWrap) Reset()         { *m = SoMemExtTrxByTrxWrap{} }
func (m *SoMemExtTrxByTrxWrap) String() string { return proto.CompactTextString(m) }
func (*SoMemExtTrxByTrxWrap) ProtoMessage()    {}
func (*SoMemExtTrxByTrxWrap) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{3}
}

func (m *SoMemExtTrxByTrxWrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByTrxWrap.Unmarshal(m, b)
}
func (m *SoMemExtTrxByTrxWrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByTrxWrap.Marshal(b, m, deterministic)
}
func (m *SoMemExtTrxByTrxWrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByTrxWrap.Merge(m, src)
}
func (m *SoMemExtTrxByTrxWrap) XXX_Size() int {
	return xxx_messageInfo_SoMemExtTrxByTrxWrap.Size(m)
}
func (m *SoMemExtTrxByTrxWrap) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtTrxByTrxWrap.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtTrxByTrxWrap proto.InternalMessageInfo

func (m *SoMemExtTrxByTrxWrap) GetTrxWrap() *prototype.TransactionWrapper {
	if m != nil {
		return m.TrxWrap
	}
	return nil
}

type SoMemExtTrxByBlockTime struct {
	BlockTime            *prototype.TimePointSec `protobuf:"bytes,1,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemExtTrxByBlockTime) Reset()         { *m = SoMemExtTrxByBlockTime{} }
func (m *SoMemExtTrxByBlockTime) String() string { return proto.CompactTextString(m) }
func (*SoMemExtTrxByBlockTime) ProtoMessage()    {}
func (*SoMemExtTrxByBlockTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{4}
}

func (m *SoMemExtTrxByBlockTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByBlockTime.Unmarshal(m, b)
}
func (m *SoMemExtTrxByBlockTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByBlockTime.Marshal(b, m, deterministic)
}
func (m *SoMemExtTrxByBlockTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByBlockTime.Merge(m, src)
}
func (m *SoMemExtTrxByBlockTime) XXX_Size() int {
	return xxx_messageInfo_SoMemExtTrxByBlockTime.Size(m)
}
func (m *SoMemExtTrxByBlockTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtTrxByBlockTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtTrxByBlockTime proto.InternalMessageInfo

func (m *SoMemExtTrxByBlockTime) GetBlockTime() *prototype.TimePointSec {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

type SoMemExtTrxByTrxCreateOrder struct {
	TrxCreateOrder       *prototype.UserTrxCreateOrder `protobuf:"bytes,1,opt,name=trx_create_order,json=trxCreateOrder,proto3" json:"trx_create_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoMemExtTrxByTrxCreateOrder) Reset()         { *m = SoMemExtTrxByTrxCreateOrder{} }
func (m *SoMemExtTrxByTrxCreateOrder) String() string { return proto.CompactTextString(m) }
func (*SoMemExtTrxByTrxCreateOrder) ProtoMessage()    {}
func (*SoMemExtTrxByTrxCreateOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{5}
}

func (m *SoMemExtTrxByTrxCreateOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Unmarshal(m, b)
}
func (m *SoMemExtTrxByTrxCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Marshal(b, m, deterministic)
}
func (m *SoMemExtTrxByTrxCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Merge(m, src)
}
func (m *SoMemExtTrxByTrxCreateOrder) XXX_Size() int {
	return xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Size(m)
}
func (m *SoMemExtTrxByTrxCreateOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtTrxByTrxCreateOrder proto.InternalMessageInfo

func (m *SoMemExtTrxByTrxCreateOrder) GetTrxCreateOrder() *prototype.UserTrxCreateOrder {
	if m != nil {
		return m.TrxCreateOrder
	}
	return nil
}

type SoListExtTrxByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoListExtTrxByTrxId) Reset()         { *m = SoListExtTrxByTrxId{} }
func (m *SoListExtTrxByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByTrxId) ProtoMessage()    {}
func (*SoListExtTrxByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{6}
}

func (m *SoListExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoListExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByTrxId.Merge(m, src)
}
func (m *SoListExtTrxByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByTrxId.Size(m)
}
func (m *SoListExtTrxByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByTrxId proto.InternalMessageInfo

func (m *SoListExtTrxByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoListExtTrxByBlockHeight struct {
	BlockHeight          uint64            `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	TrxId                *prototype.Sha256 `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoListExtTrxByBlockHeight) Reset()         { *m = SoListExtTrxByBlockHeight{} }
func (m *SoListExtTrxByBlockHeight) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByBlockHeight) ProtoMessage()    {}
func (*SoListExtTrxByBlockHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{7}
}

func (m *SoListExtTrxByBlockHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Unmarshal(m, b)
}
func (m *SoListExtTrxByBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByBlockHeight.Merge(m, src)
}
func (m *SoListExtTrxByBlockHeight) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Size(m)
}
func (m *SoListExtTrxByBlockHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByBlockHeight.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByBlockHeight proto.InternalMessageInfo

func (m *SoListExtTrxByBlockHeight) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SoListExtTrxByBlockHeight) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoListExtTrxByBlockTime struct {
	BlockTime            *prototype.TimePointSec `protobuf:"bytes,1,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	TrxId                *prototype.Sha256       `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtTrxByBlockTime) Reset()         { *m = SoListExtTrxByBlockTime{} }
func (m *SoListExtTrxByBlockTime) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByBlockTime) ProtoMessage()    {}
func (*SoListExtTrxByBlockTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{8}
}

func (m *SoListExtTrxByBlockTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Unmarshal(m, b)
}
func (m *SoListExtTrxByBlockTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByBlockTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByBlockTime.Merge(m, src)
}
func (m *SoListExtTrxByBlockTime) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Size(m)
}
func (m *SoListExtTrxByBlockTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByBlockTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByBlockTime proto.InternalMessageInfo

func (m *SoListExtTrxByBlockTime) GetBlockTime() *prototype.TimePointSec {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

func (m *SoListExtTrxByBlockTime) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoListExtTrxByTrxCreateOrder struct {
	TrxCreateOrder       *prototype.UserTrxCreateOrder `protobuf:"bytes,1,opt,name=trx_create_order,json=trxCreateOrder,proto3" json:"trx_create_order,omitempty"`
	TrxId                *prototype.Sha256             `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoListExtTrxByTrxCreateOrder) Reset()         { *m = SoListExtTrxByTrxCreateOrder{} }
func (m *SoListExtTrxByTrxCreateOrder) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByTrxCreateOrder) ProtoMessage()    {}
func (*SoListExtTrxByTrxCreateOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{9}
}

func (m *SoListExtTrxByTrxCreateOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Unmarshal(m, b)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Merge(m, src)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Size(m)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByTrxCreateOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByTrxCreateOrder proto.InternalMessageInfo

func (m *SoListExtTrxByTrxCreateOrder) GetTrxCreateOrder() *prototype.UserTrxCreateOrder {
	if m != nil {
		return m.TrxCreateOrder
	}
	return nil
}

func (m *SoListExtTrxByTrxCreateOrder) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoUniqueExtTrxByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoUniqueExtTrxByTrxId) Reset()         { *m = SoUniqueExtTrxByTrxId{} }
func (m *SoUniqueExtTrxByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtTrxByTrxId) ProtoMessage()    {}
func (*SoUniqueExtTrxByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{10}
}

func (m *SoUniqueExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtTrxByTrxId.Merge(m, src)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Size(m)
}
func (m *SoUniqueExtTrxByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtTrxByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtTrxByTrxId proto.InternalMessageInfo

func (m *SoUniqueExtTrxByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtTrx)(nil), "table.so_extTrx")
	proto.RegisterType((*SoMemExtTrxByTrxId)(nil), "table.so_mem_extTrx_by_trx_id")
	proto.RegisterType((*SoMemExtTrxByBlockHeight)(nil), "table.so_mem_extTrx_by_block_height")
	proto.RegisterType((*SoMemExtTrxByTrxWrap)(nil), "table.so_mem_extTrx_by_trx_wrap")
	proto.RegisterType((*SoMemExtTrxByBlockTime)(nil), "table.so_mem_extTrx_by_block_time")
	proto.RegisterType((*SoMemExtTrxByTrxCreateOrder)(nil), "table.so_mem_extTrx_by_trx_create_order")
	proto.RegisterType((*SoListExtTrxByTrxId)(nil), "table.so_list_extTrx_by_trx_id")
	proto.RegisterType((*SoListExtTrxByBlockHeight)(nil), "table.so_list_extTrx_by_block_height")
	proto.RegisterType((*SoListExtTrxByBlockTime)(nil), "table.so_list_extTrx_by_block_time")
	proto.RegisterType((*SoListExtTrxByTrxCreateOrder)(nil), "table.so_list_extTrx_by_trx_create_order")
	proto.RegisterType((*SoUniqueExtTrxByTrxId)(nil), "table.so_unique_extTrx_by_trx_id")
}

func init() { proto.RegisterFile("app/table/so_extTrx.proto", fileDescriptor_76957eaae1f8a1bc) }

var fileDescriptor_76957eaae1f8a1bc = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0x99, 0xb5, 0x5b, 0xed, 0xad, 0x88, 0x06, 0xc1, 0xec, 0x56, 0xcb, 0x36, 0x4f, 0x8b,
	0xe8, 0x06, 0x2a, 0x8a, 0xbe, 0xb6, 0x22, 0xea, 0x8b, 0x10, 0x8a, 0x05, 0x5f, 0x86, 0x49, 0xf6,
	0xb2, 0x19, 0xdc, 0xe4, 0xc6, 0x99, 0x1b, 0x4c, 0x5f, 0xfd, 0x16, 0xfa, 0x69, 0x25, 0xb3, 0x65,
	0x1d, 0xbb, 0x5b, 0xdc, 0xb8, 0xf4, 0x25, 0x30, 0xf7, 0xcf, 0xb9, 0x87, 0xdf, 0x81, 0xc0, 0x40,
	0x55, 0x55, 0xcc, 0x2a, 0x9d, 0x63, 0x6c, 0x49, 0x62, 0xc3, 0x67, 0xa6, 0x99, 0x54, 0x86, 0x98,
	0x82, 0xbe, 0x2b, 0x0f, 0x1f, 0xba, 0x17, 0x5f, 0x54, 0x18, 0xb7, 0x9f, 0x45, 0x73, 0x78, 0xe0,
	0x55, 0x8d, 0x2a, 0xad, 0xca, 0x58, 0x53, 0x79, 0xd9, 0x0c, 0xff, 0x34, 0x8b, 0x7a, 0xce, 0x5a,
	0xea, 0xe9, 0xa2, 0x13, 0xfd, 0xec, 0xc1, 0xde, 0xf2, 0x4e, 0x30, 0x86, 0x5d, 0x36, 0x8d, 0xd4,
	0xd3, 0x50, 0x8c, 0xc4, 0x78, 0xff, 0xf8, 0xc1, 0x64, 0xb9, 0x38, 0xb1, 0xb9, 0x3a, 0x7e, 0xf9,
	0x2a, 0xe9, 0xb3, 0x69, 0x3e, 0x4c, 0x83, 0x23, 0xb8, 0x9b, 0xce, 0x29, 0xfb, 0x2a, 0x73, 0xd4,
	0xb3, 0x9c, 0xc3, 0xde, 0x48, 0x8c, 0x77, 0x92, 0x7d, 0x57, 0x7b, 0xef, 0x4a, 0xc1, 0x1b, 0xb8,
	0xd3, 0x8a, 0x7d, 0x37, 0xaa, 0x0a, 0x6f, 0x39, 0xb9, 0x43, 0x4f, 0xce, 0x33, 0xe9, 0x46, 0x2a,
	0x34, 0xc9, 0x6d, 0x36, 0xcd, 0xb9, 0x51, 0x55, 0xf0, 0x1a, 0x60, 0xa1, 0xce, 0xba, 0xc0, 0x70,
	0xc7, 0x2d, 0x0f, 0xfc, 0x65, 0x5d, 0xa0, 0xac, 0x48, 0x97, 0x2c, 0x2d, 0x66, 0xc9, 0x9e, 0x1b,
	0x3e, 0xd3, 0x05, 0x06, 0x1f, 0xe1, 0x7e, 0x7b, 0x34, 0x33, 0xa8, 0x18, 0x25, 0x99, 0x29, 0x9a,
	0xb0, 0xef, 0xf6, 0x47, 0xde, 0x7e, 0x6d, 0xd1, 0xc8, 0xab, 0x73, 0xc9, 0x3d, 0x36, 0xcd, 0xa9,
	0x2b, 0x7c, 0x6a, 0xdf, 0xd1, 0x29, 0x3c, 0xb2, 0x24, 0x0b, 0x2c, 0x2e, 0xf1, 0xc8, 0xf4, 0x42,
	0x2e, 0xf0, 0x6c, 0x0e, 0x2a, 0x3a, 0x81, 0x27, 0x2b, 0x22, 0x3e, 0xb9, 0x15, 0x92, 0x62, 0x85,
	0x64, 0xf4, 0x19, 0x06, 0x6b, 0x8d, 0xb4, 0xdc, 0xfe, 0xc2, 0x2c, 0x3a, 0x61, 0x8e, 0xce, 0xe1,
	0xe0, 0x1a, 0x6f, 0x2d, 0xe0, 0x2b, 0x29, 0x88, 0xcd, 0x53, 0x88, 0x08, 0x8e, 0xd6, 0x1a, 0xf6,
	0x71, 0xaf, 0x8d, 0x4a, 0xfc, 0x67, 0x54, 0x6f, 0x21, 0xb4, 0x24, 0xe7, 0xda, 0xf2, 0x36, 0x59,
	0x15, 0x70, 0xb8, 0xaa, 0xd2, 0x31, 0x2c, 0xef, 0x5c, 0xef, 0x1f, 0xe7, 0x7e, 0x08, 0x78, 0x7c,
	0xdd, 0xbd, 0xed, 0x02, 0xe8, 0x60, 0xe2, 0x97, 0x80, 0x68, 0x3d, 0xba, 0x9b, 0x0a, 0xab, 0x83,
	0xb9, 0x77, 0x30, 0xb4, 0x24, 0xeb, 0x52, 0x7f, 0xab, 0x71, 0x8b, 0x60, 0x4f, 0x9e, 0x7d, 0x79,
	0x3a, 0xd3, 0x9c, 0xd7, 0xe9, 0x24, 0xa3, 0x22, 0xce, 0xc8, 0x66, 0xb9, 0xd2, 0x65, 0x9c, 0x51,
	0xc9, 0x58, 0x32, 0xd9, 0xe7, 0x33, 0x8a, 0x97, 0xff, 0xdd, 0x74, 0xd7, 0xc9, 0xbc, 0xf8, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x78, 0xa1, 0xd2, 0x1a, 0x8b, 0x05, 0x00, 0x00,
}
