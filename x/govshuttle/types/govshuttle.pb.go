// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canto/govshuttle/v1/govshuttle.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f3a63fcc428040, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// Define this object so that the govshuttle.pb.go file is generate, implements
// govtypes.Content
type LendingMarketProposal struct {
	// title
	Title       string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Metadata    *LendingMarketMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *LendingMarketProposal) Reset()         { *m = LendingMarketProposal{} }
func (m *LendingMarketProposal) String() string { return proto.CompactTextString(m) }
func (*LendingMarketProposal) ProtoMessage()    {}
func (*LendingMarketProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f3a63fcc428040, []int{1}
}
func (m *LendingMarketProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LendingMarketProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LendingMarketProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LendingMarketProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendingMarketProposal.Merge(m, src)
}
func (m *LendingMarketProposal) XXX_Size() int {
	return m.Size()
}
func (m *LendingMarketProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_LendingMarketProposal.DiscardUnknown(m)
}

var xxx_messageInfo_LendingMarketProposal proto.InternalMessageInfo

func (m *LendingMarketProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *LendingMarketProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LendingMarketProposal) GetMetadata() *LendingMarketMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// treasury proposal type,
type TreasuryProposal struct {
	Title       string                    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Metadata    *TreasuryProposalMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *TreasuryProposal) Reset()         { *m = TreasuryProposal{} }
func (m *TreasuryProposal) String() string { return proto.CompactTextString(m) }
func (*TreasuryProposal) ProtoMessage()    {}
func (*TreasuryProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f3a63fcc428040, []int{2}
}
func (m *TreasuryProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TreasuryProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TreasuryProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TreasuryProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreasuryProposal.Merge(m, src)
}
func (m *TreasuryProposal) XXX_Size() int {
	return m.Size()
}
func (m *TreasuryProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TreasuryProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TreasuryProposal proto.InternalMessageInfo

func (m *TreasuryProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TreasuryProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TreasuryProposal) GetMetadata() *TreasuryProposalMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type TreasuryProposalMetadata struct {
	PropID    uint64 `protobuf:"varint,1,opt,name=PropID,proto3" json:"PropID,omitempty"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount    uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom     string `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *TreasuryProposalMetadata) Reset()         { *m = TreasuryProposalMetadata{} }
func (m *TreasuryProposalMetadata) String() string { return proto.CompactTextString(m) }
func (*TreasuryProposalMetadata) ProtoMessage()    {}
func (*TreasuryProposalMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f3a63fcc428040, []int{3}
}
func (m *TreasuryProposalMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TreasuryProposalMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TreasuryProposalMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TreasuryProposalMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreasuryProposalMetadata.Merge(m, src)
}
func (m *TreasuryProposalMetadata) XXX_Size() int {
	return m.Size()
}
func (m *TreasuryProposalMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TreasuryProposalMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TreasuryProposalMetadata proto.InternalMessageInfo

func (m *TreasuryProposalMetadata) GetPropID() uint64 {
	if m != nil {
		return m.PropID
	}
	return 0
}

func (m *TreasuryProposalMetadata) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *TreasuryProposalMetadata) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TreasuryProposalMetadata) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type LendingMarketMetadata struct {
	Account    []string `protobuf:"bytes,1,rep,name=Account,proto3" json:"Account,omitempty"`
	PropId     uint64   `protobuf:"varint,2,opt,name=PropId,proto3" json:"PropId,omitempty"`
	Values     []uint64 `protobuf:"varint,3,rep,packed,name=values,proto3" json:"values,omitempty"`
	Calldatas  []string `protobuf:"bytes,4,rep,name=calldatas,proto3" json:"calldatas,omitempty"`
	Signatures []string `protobuf:"bytes,5,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *LendingMarketMetadata) Reset()         { *m = LendingMarketMetadata{} }
func (m *LendingMarketMetadata) String() string { return proto.CompactTextString(m) }
func (*LendingMarketMetadata) ProtoMessage()    {}
func (*LendingMarketMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f3a63fcc428040, []int{4}
}
func (m *LendingMarketMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LendingMarketMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LendingMarketMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LendingMarketMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendingMarketMetadata.Merge(m, src)
}
func (m *LendingMarketMetadata) XXX_Size() int {
	return m.Size()
}
func (m *LendingMarketMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_LendingMarketMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_LendingMarketMetadata proto.InternalMessageInfo

func (m *LendingMarketMetadata) GetAccount() []string {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *LendingMarketMetadata) GetPropId() uint64 {
	if m != nil {
		return m.PropId
	}
	return 0
}

func (m *LendingMarketMetadata) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *LendingMarketMetadata) GetCalldatas() []string {
	if m != nil {
		return m.Calldatas
	}
	return nil
}

func (m *LendingMarketMetadata) GetSignatures() []string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "canto.govshuttle.v1.Params")
	proto.RegisterType((*LendingMarketProposal)(nil), "canto.govshuttle.v1.LendingMarketProposal")
	proto.RegisterType((*TreasuryProposal)(nil), "canto.govshuttle.v1.TreasuryProposal")
	proto.RegisterType((*TreasuryProposalMetadata)(nil), "canto.govshuttle.v1.TreasuryProposalMetadata")
	proto.RegisterType((*LendingMarketMetadata)(nil), "canto.govshuttle.v1.LendingMarketMetadata")
}

func init() {
	proto.RegisterFile("canto/govshuttle/v1/govshuttle.proto", fileDescriptor_39f3a63fcc428040)
}

var fileDescriptor_39f3a63fcc428040 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xcf, 0xb8, 0xe9, 0xea, 0x4e, 0x41, 0x24, 0x56, 0x19, 0x8b, 0xa4, 0x61, 0xf1, 0xb0, 0x08,
	0x9b, 0xb0, 0xfe, 0x01, 0xd9, 0x93, 0x5a, 0x11, 0x0a, 0xb6, 0x94, 0xe0, 0xc9, 0x8b, 0xcc, 0x26,
	0x43, 0x1a, 0x36, 0x99, 0x09, 0x33, 0x2f, 0xd1, 0xde, 0xfd, 0x00, 0x1e, 0xbd, 0xd9, 0x0f, 0xe1,
	0x87, 0xf0, 0x58, 0xbc, 0xa8, 0x37, 0xd9, 0xbd, 0xf8, 0x31, 0x64, 0x66, 0xb2, 0xdb, 0x58, 0xeb,
	0x49, 0x3c, 0x25, 0xbf, 0xdf, 0x9b, 0xf7, 0xde, 0xef, 0xf7, 0xde, 0x0c, 0xbe, 0x93, 0x50, 0x0e,
	0x22, 0xca, 0x44, 0xa3, 0x8e, 0x6a, 0x80, 0x82, 0x45, 0xcd, 0xa4, 0x83, 0xc2, 0x4a, 0x0a, 0x10,
	0xde, 0x75, 0x73, 0x2a, 0xec, 0xf0, 0xcd, 0x64, 0x7b, 0x2b, 0x13, 0x99, 0x30, 0xf1, 0x48, 0xff,
	0xd9, 0xa3, 0xdb, 0xb7, 0x12, 0xa1, 0x4a, 0xa1, 0x5e, 0xdb, 0x80, 0x05, 0x36, 0x34, 0xbc, 0x8a,
	0xfb, 0x87, 0x54, 0xd2, 0x52, 0x4d, 0xdd, 0x0f, 0x27, 0x3b, 0xce, 0xf0, 0x3b, 0xc2, 0x37, 0x5e,
	0x30, 0x9e, 0xe6, 0x3c, 0xdb, 0xa7, 0x72, 0xce, 0xe0, 0x50, 0x8a, 0x4a, 0x28, 0x5a, 0x78, 0x5b,
	0x78, 0x03, 0x72, 0x28, 0x18, 0x41, 0x01, 0x1a, 0x0d, 0x62, 0x0b, 0xbc, 0x00, 0x6f, 0xa6, 0x4c,
	0x25, 0x32, 0xaf, 0x20, 0x17, 0x9c, 0x5c, 0x32, 0xb1, 0x2e, 0xe5, 0x3d, 0xc7, 0x57, 0x4a, 0x06,
	0x34, 0xa5, 0x40, 0x49, 0x2f, 0x40, 0xa3, 0xcd, 0x7b, 0x77, 0xc3, 0x0b, 0xa4, 0x87, 0xbf, 0x75,
	0xdd, 0x6f, 0x33, 0xe2, 0x75, 0xee, 0xf4, 0xf1, 0xcf, 0x93, 0x1d, 0xe7, 0xcb, 0xa7, 0xf1, 0xa3,
	0x2c, 0x87, 0xa3, 0x7a, 0x16, 0x26, 0xa2, 0x6c, 0xad, 0xb4, 0x9f, 0xb1, 0x4a, 0xe7, 0xd1, 0x5b,
	0x3d, 0xa8, 0x08, 0x8e, 0x2b, 0xa6, 0xa2, 0x66, 0x32, 0x63, 0x40, 0x27, 0xe1, 0xae, 0xe0, 0xc0,
	0x38, 0x0c, 0xbf, 0x22, 0x7c, 0xed, 0xa5, 0x64, 0x54, 0xd5, 0xf2, 0xf8, 0x9f, 0x6d, 0xed, 0xfd,
	0x61, 0x6b, 0x7c, 0xa1, 0xad, 0xf3, 0x0d, 0xff, 0x8b, 0xb3, 0x77, 0x08, 0x93, 0xbf, 0x35, 0xf2,
	0x6e, 0xe2, 0xbe, 0xe6, 0xf6, 0x9e, 0x19, 0x8b, 0x6e, 0xdc, 0x22, 0xef, 0x36, 0x1e, 0x48, 0x96,
	0xe4, 0x55, 0xce, 0x38, 0xb4, 0x0e, 0xcf, 0x08, 0x9d, 0x45, 0x4b, 0x51, 0x73, 0x30, 0xee, 0xdc,
	0xb8, 0x45, 0x7a, 0x5e, 0x29, 0xe3, 0xa2, 0x24, 0xae, 0x9d, 0x97, 0x01, 0x53, 0x57, 0x5b, 0x18,
	0x7e, 0x3c, 0x7f, 0x79, 0xd6, 0x1a, 0x08, 0xbe, 0xfc, 0x24, 0x49, 0x4c, 0x39, 0x14, 0xf4, 0x46,
	0x83, 0x78, 0x05, 0xd7, 0xea, 0x52, 0x23, 0x61, 0xa5, 0x2e, 0xd5, 0x7c, 0x43, 0x8b, 0x9a, 0x29,
	0xd2, 0x0b, 0x7a, 0x9a, 0xb7, 0x48, 0xab, 0x4e, 0x68, 0x51, 0xe8, 0xaa, 0x8a, 0xb8, 0xa6, 0xd6,
	0x19, 0xe1, 0xf9, 0x18, 0xab, 0x3c, 0xe3, 0x14, 0x6a, 0xc9, 0x14, 0xd9, 0x30, 0xe1, 0x0e, 0xf3,
	0xf4, 0xe0, 0xf3, 0xc2, 0x47, 0xa7, 0x0b, 0x1f, 0xfd, 0x58, 0xf8, 0xe8, 0xfd, 0xd2, 0x77, 0x4e,
	0x97, 0xbe, 0xf3, 0x6d, 0xe9, 0x3b, 0xaf, 0x1e, 0x74, 0x86, 0xbf, 0xab, 0xf7, 0x38, 0x3e, 0x60,
	0xf0, 0x46, 0xc8, 0xb9, 0x45, 0x51, 0xf3, 0xd0, 0x6e, 0x60, 0xf5, 0x24, 0xcd, 0x22, 0x66, 0x7d,
	0xf3, 0x8a, 0xee, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xc9, 0x8c, 0xa3, 0xb3, 0x03, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *LendingMarketProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LendingMarketProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LendingMarketProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGovshuttle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TreasuryProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TreasuryProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TreasuryProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGovshuttle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TreasuryProposalMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TreasuryProposalMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TreasuryProposalMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if m.Amount != 0 {
		i = encodeVarintGovshuttle(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.PropID != 0 {
		i = encodeVarintGovshuttle(dAtA, i, uint64(m.PropID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LendingMarketMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LendingMarketMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LendingMarketMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Calldatas) > 0 {
		for iNdEx := len(m.Calldatas) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Calldatas[iNdEx])
			copy(dAtA[i:], m.Calldatas[iNdEx])
			i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Calldatas[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Values) > 0 {
		dAtA4 := make([]byte, len(m.Values)*10)
		var j3 int
		for _, num := range m.Values {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGovshuttle(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	if m.PropId != 0 {
		i = encodeVarintGovshuttle(dAtA, i, uint64(m.PropId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Account) > 0 {
		for iNdEx := len(m.Account) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Account[iNdEx])
			copy(dAtA[i:], m.Account[iNdEx])
			i = encodeVarintGovshuttle(dAtA, i, uint64(len(m.Account[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGovshuttle(dAtA []byte, offset int, v uint64) int {
	offset -= sovGovshuttle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *LendingMarketProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	return n
}

func (m *TreasuryProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	return n
}

func (m *TreasuryProposalMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PropID != 0 {
		n += 1 + sovGovshuttle(uint64(m.PropID))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovGovshuttle(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGovshuttle(uint64(l))
	}
	return n
}

func (m *LendingMarketMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Account) > 0 {
		for _, s := range m.Account {
			l = len(s)
			n += 1 + l + sovGovshuttle(uint64(l))
		}
	}
	if m.PropId != 0 {
		n += 1 + sovGovshuttle(uint64(m.PropId))
	}
	if len(m.Values) > 0 {
		l = 0
		for _, e := range m.Values {
			l += sovGovshuttle(uint64(e))
		}
		n += 1 + sovGovshuttle(uint64(l)) + l
	}
	if len(m.Calldatas) > 0 {
		for _, s := range m.Calldatas {
			l = len(s)
			n += 1 + l + sovGovshuttle(uint64(l))
		}
	}
	if len(m.Signatures) > 0 {
		for _, s := range m.Signatures {
			l = len(s)
			n += 1 + l + sovGovshuttle(uint64(l))
		}
	}
	return n
}

func sovGovshuttle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGovshuttle(x uint64) (n int) {
	return sovGovshuttle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovshuttle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGovshuttle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LendingMarketProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovshuttle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LendingMarketProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LendingMarketProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &LendingMarketMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovshuttle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TreasuryProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovshuttle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TreasuryProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TreasuryProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &TreasuryProposalMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovshuttle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TreasuryProposalMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovshuttle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TreasuryProposalMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TreasuryProposalMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropID", wireType)
			}
			m.PropID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PropID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovshuttle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LendingMarketMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovshuttle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LendingMarketMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LendingMarketMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = append(m.Account, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropId", wireType)
			}
			m.PropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGovshuttle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Values = append(m.Values, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGovshuttle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGovshuttle
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGovshuttle
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Values) == 0 {
					m.Values = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGovshuttle
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Values = append(m.Values, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Calldatas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Calldatas = append(m.Calldatas, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGovshuttle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovshuttle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovshuttle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGovshuttle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGovshuttle
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGovshuttle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGovshuttle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGovshuttle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGovshuttle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGovshuttle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGovshuttle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGovshuttle = fmt.Errorf("proto: unexpected end of group")
)
