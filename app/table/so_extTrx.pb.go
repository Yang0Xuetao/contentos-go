// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extTrx.proto

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

type SoExtTrx struct {
	TrxId                *prototype.Sha256             `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	BlockHeight          uint64                        `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	TrxWrap              *prototype.TransactionWrapper `protobuf:"bytes,3,opt,name=trx_wrap,json=trxWrap,proto3" json:"trx_wrap,omitempty"`
	BlockTime            *prototype.TimePointSec       `protobuf:"bytes,4,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	TrxCreateOrder       *prototype.UserTrxCreateOrder `protobuf:"bytes,5,opt,name=trx_create_order,json=trxCreateOrder,proto3" json:"trx_create_order,omitempty"`
	BlockId              *prototype.Sha256             `protobuf:"bytes,6,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoExtTrx) Reset()         { *m = SoExtTrx{} }
func (m *SoExtTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtTrx) ProtoMessage()    {}
func (*SoExtTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{0}
}
func (m *SoExtTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtTrx.Unmarshal(m, b)
}
func (m *SoExtTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtTrx.Marshal(b, m, deterministic)
}
func (dst *SoExtTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtTrx.Merge(dst, src)
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

func (m *SoExtTrx) GetBlockId() *prototype.Sha256 {
	if m != nil {
		return m.BlockId
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{1}
}
func (m *SoMemExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoMemExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (dst *SoMemExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByTrxId.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{2}
}
func (m *SoMemExtTrxByBlockHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByBlockHeight.Unmarshal(m, b)
}
func (m *SoMemExtTrxByBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByBlockHeight.Marshal(b, m, deterministic)
}
func (dst *SoMemExtTrxByBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByBlockHeight.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{3}
}
func (m *SoMemExtTrxByTrxWrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByTrxWrap.Unmarshal(m, b)
}
func (m *SoMemExtTrxByTrxWrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByTrxWrap.Marshal(b, m, deterministic)
}
func (dst *SoMemExtTrxByTrxWrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByTrxWrap.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{4}
}
func (m *SoMemExtTrxByBlockTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByBlockTime.Unmarshal(m, b)
}
func (m *SoMemExtTrxByBlockTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByBlockTime.Marshal(b, m, deterministic)
}
func (dst *SoMemExtTrxByBlockTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByBlockTime.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{5}
}
func (m *SoMemExtTrxByTrxCreateOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Unmarshal(m, b)
}
func (m *SoMemExtTrxByTrxCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Marshal(b, m, deterministic)
}
func (dst *SoMemExtTrxByTrxCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByTrxCreateOrder.Merge(dst, src)
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

type SoMemExtTrxByBlockId struct {
	BlockId              *prototype.Sha256 `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoMemExtTrxByBlockId) Reset()         { *m = SoMemExtTrxByBlockId{} }
func (m *SoMemExtTrxByBlockId) String() string { return proto.CompactTextString(m) }
func (*SoMemExtTrxByBlockId) ProtoMessage()    {}
func (*SoMemExtTrxByBlockId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{6}
}
func (m *SoMemExtTrxByBlockId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtTrxByBlockId.Unmarshal(m, b)
}
func (m *SoMemExtTrxByBlockId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtTrxByBlockId.Marshal(b, m, deterministic)
}
func (dst *SoMemExtTrxByBlockId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtTrxByBlockId.Merge(dst, src)
}
func (m *SoMemExtTrxByBlockId) XXX_Size() int {
	return xxx_messageInfo_SoMemExtTrxByBlockId.Size(m)
}
func (m *SoMemExtTrxByBlockId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtTrxByBlockId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtTrxByBlockId proto.InternalMessageInfo

func (m *SoMemExtTrxByBlockId) GetBlockId() *prototype.Sha256 {
	if m != nil {
		return m.BlockId
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{7}
}
func (m *SoListExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoListExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (dst *SoListExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByTrxId.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{8}
}
func (m *SoListExtTrxByBlockHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Unmarshal(m, b)
}
func (m *SoListExtTrxByBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Marshal(b, m, deterministic)
}
func (dst *SoListExtTrxByBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByBlockHeight.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{9}
}
func (m *SoListExtTrxByBlockTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Unmarshal(m, b)
}
func (m *SoListExtTrxByBlockTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Marshal(b, m, deterministic)
}
func (dst *SoListExtTrxByBlockTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByBlockTime.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{10}
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Unmarshal(m, b)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Marshal(b, m, deterministic)
}
func (dst *SoListExtTrxByTrxCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Merge(dst, src)
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
	return fileDescriptor_so_extTrx_ae86d1ba901d74e9, []int{11}
}
func (m *SoUniqueExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (dst *SoUniqueExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtTrxByTrxId.Merge(dst, src)
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
	proto.RegisterType((*SoMemExtTrxByBlockId)(nil), "table.so_mem_extTrx_by_block_id")
	proto.RegisterType((*SoListExtTrxByTrxId)(nil), "table.so_list_extTrx_by_trx_id")
	proto.RegisterType((*SoListExtTrxByBlockHeight)(nil), "table.so_list_extTrx_by_block_height")
	proto.RegisterType((*SoListExtTrxByBlockTime)(nil), "table.so_list_extTrx_by_block_time")
	proto.RegisterType((*SoListExtTrxByTrxCreateOrder)(nil), "table.so_list_extTrx_by_trx_create_order")
	proto.RegisterType((*SoUniqueExtTrxByTrxId)(nil), "table.so_unique_extTrx_by_trx_id")
}

func init() {
	proto.RegisterFile("app/table/so_extTrx.proto", fileDescriptor_so_extTrx_ae86d1ba901d74e9)
}

var fileDescriptor_so_extTrx_ae86d1ba901d74e9 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xd1, 0x6b, 0xd3, 0x50,
	0x14, 0xc6, 0xb9, 0x75, 0xed, 0xb6, 0x33, 0x11, 0x0d, 0x82, 0x69, 0xa7, 0xa3, 0xcb, 0x53, 0x91,
	0xd9, 0xc0, 0x44, 0xd1, 0xd7, 0x4d, 0xc4, 0xfa, 0x22, 0x94, 0xe1, 0xc0, 0x97, 0xcb, 0x4d, 0x72,
	0x68, 0x2e, 0x36, 0xb9, 0xd7, 0x7b, 0x4f, 0x30, 0x7d, 0xf5, 0xcf, 0xf0, 0x1f, 0xf2, 0xdf, 0x92,
	0xdc, 0xcc, 0x9a, 0xad, 0x29, 0x5b, 0x2d, 0x3e, 0xe6, 0xdc, 0xf3, 0x7d, 0xe7, 0xe3, 0xfb, 0x41,
	0xa0, 0x2f, 0xb4, 0x0e, 0x49, 0x44, 0x73, 0x0c, 0xad, 0xe2, 0x58, 0xd2, 0x85, 0x29, 0xc7, 0xda,
	0x28, 0x52, 0x5e, 0xd7, 0x8d, 0x07, 0x87, 0xee, 0x8b, 0x16, 0x1a, 0x43, 0x32, 0x22, 0xb7, 0x22,
	0x26, 0xa9, 0xf2, 0x7a, 0x67, 0xe0, 0xff, 0x7d, 0xcc, 0x8a, 0x39, 0x49, 0x2e, 0x93, 0xab, 0x97,
	0xc7, 0x0d, 0xd9, 0x42, 0x63, 0x3d, 0x0d, 0x7e, 0x75, 0x60, 0x7f, 0x79, 0xc7, 0x1b, 0x41, 0x8f,
	0x4c, 0xc9, 0x65, 0xe2, 0xb3, 0x21, 0x1b, 0x1d, 0x9c, 0x3e, 0x1a, 0x2f, 0x45, 0x63, 0x9b, 0x8a,
	0xd3, 0x57, 0xaf, 0xa7, 0x5d, 0x32, 0xe5, 0x24, 0xf1, 0x8e, 0xe1, 0x7e, 0x34, 0x57, 0xf1, 0x57,
	0x9e, 0xa2, 0x9c, 0xa5, 0xe4, 0x77, 0x86, 0x6c, 0xb4, 0x33, 0x3d, 0x70, 0xb3, 0x0f, 0x6e, 0xe4,
	0xbd, 0x85, 0xbd, 0xca, 0xec, 0xbb, 0x11, 0xda, 0xbf, 0xe7, 0xec, 0x8e, 0x1a, 0x76, 0x8d, 0xe8,
	0x6e, 0x45, 0xa3, 0x99, 0xee, 0x92, 0x29, 0x2f, 0x8d, 0xd0, 0xde, 0x1b, 0x80, 0xda, 0x9d, 0x64,
	0x86, 0xfe, 0x8e, 0x13, 0xf7, 0x9b, 0x62, 0x99, 0x21, 0xd7, 0x4a, 0xe6, 0xc4, 0x2d, 0xc6, 0xd3,
	0x7d, 0xb7, 0x7c, 0x21, 0x33, 0xf4, 0x3e, 0xc2, 0xc3, 0xea, 0x68, 0x6c, 0x50, 0x10, 0x72, 0x65,
	0x12, 0x34, 0x7e, 0xd7, 0xe9, 0x87, 0x0d, 0x7d, 0x61, 0xd1, 0xf0, 0x9b, 0x7b, 0xd3, 0x07, 0x64,
	0xca, 0x73, 0x37, 0xf8, 0x54, 0x7d, 0x7b, 0x27, 0xb0, 0x57, 0xa7, 0x90, 0x89, 0xdf, 0x5b, 0xd7,
	0xc7, 0xae, 0x5b, 0x99, 0x24, 0xc1, 0x39, 0x3c, 0xb1, 0x8a, 0x67, 0x98, 0x5d, 0x95, 0xc9, 0xa3,
	0x05, 0xaf, 0xcb, 0xbc, 0x7b, 0xad, 0xc1, 0x19, 0x3c, 0x5b, 0x31, 0x69, 0xf6, 0xbc, 0xd2, 0x3b,
	0x5b, 0xe9, 0x3d, 0xf8, 0x0c, 0xfd, 0xd6, 0x20, 0x55, 0xcb, 0xd7, 0xa0, 0xb0, 0x8d, 0xa0, 0x04,
	0x97, 0x70, 0xb8, 0x26, 0x5b, 0x85, 0xe3, 0x06, 0x33, 0x76, 0x77, 0x66, 0x81, 0x82, 0xe3, 0xd6,
	0xc0, 0x4d, 0x38, 0xad, 0x60, 0xd9, 0xbf, 0x81, 0x0d, 0x26, 0x2d, 0x0d, 0xfd, 0x21, 0x7d, 0x8d,
	0x3a, 0xbb, 0x95, 0xfa, 0x3b, 0xf0, 0xad, 0xe2, 0x73, 0x69, 0x69, 0x1b, 0xec, 0x19, 0x1c, 0xad,
	0xba, 0x6c, 0xc8, 0xbd, 0x71, 0xae, 0x73, 0xcb, 0xb9, 0x1f, 0x0c, 0x9e, 0xae, 0xbb, 0xb7, 0x1d,
	0xcb, 0x0d, 0x42, 0xfc, 0x64, 0x10, 0xb4, 0x57, 0xf7, 0xbf, 0xb8, 0x6f, 0x10, 0xee, 0x3d, 0x0c,
	0xac, 0xe2, 0x45, 0x2e, 0xbf, 0x15, 0xb8, 0x05, 0xd8, 0xb3, 0x93, 0x2f, 0xcf, 0x67, 0x92, 0xd2,
	0x22, 0x1a, 0xc7, 0x2a, 0x0b, 0x63, 0x65, 0xe3, 0x54, 0xc8, 0x3c, 0x8c, 0x55, 0x4e, 0x98, 0x93,
	0xb2, 0x2f, 0x66, 0x2a, 0x5c, 0xfe, 0xf0, 0xa3, 0x9e, 0xb3, 0x79, 0xf9, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0xab, 0x60, 0xf6, 0x04, 0x06, 0x00, 0x00,
}
