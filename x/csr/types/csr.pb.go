// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canto/csr/v1/csr.proto

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

// The CSR struct is a wrapper to all of the metadata associated with a given Contract Secured Revenue registration. It maintains the pool
// where all of the fees are being sent to, the deployer that is responsible for all deployments, and the set of dApps (smart contracts)
// that are registered with this pool.
type CSR struct {
	// The owner keeps track of the user that deploys all of the dApps, same user that initially owns the NFT
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Contracts is the list of all EVM address that are registered to this pool (EVM addresses)
	Contracts []string `protobuf:"bytes,2,rep,name=contracts,proto3" json:"contracts,omitempty"`
	// The NFT id which this CSR corresponds to
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// The account which will be accumulating rewards for this CSR (bech32 Canto address)
	Account string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *CSR) Reset()         { *m = CSR{} }
func (m *CSR) String() string { return proto.CompactTextString(m) }
func (*CSR) ProtoMessage()    {}
func (*CSR) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c53cea3d443afa, []int{0}
}
func (m *CSR) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CSR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CSR.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CSR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSR.Merge(m, src)
}
func (m *CSR) XXX_Size() int {
	return m.Size()
}
func (m *CSR) XXX_DiscardUnknown() {
	xxx_messageInfo_CSR.DiscardUnknown(m)
}

var xxx_messageInfo_CSR proto.InternalMessageInfo

func (m *CSR) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CSR) GetContracts() []string {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *CSR) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CSR) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// Keeps track of all of the CSRs, primarily used to (un)marshal data
type CSRs struct {
	Csrs []*CSR `protobuf:"bytes,1,rep,name=csrs,proto3" json:"csrs,omitempty"`
}

func (m *CSRs) Reset()         { *m = CSRs{} }
func (m *CSRs) String() string { return proto.CompactTextString(m) }
func (*CSRs) ProtoMessage()    {}
func (*CSRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c53cea3d443afa, []int{1}
}
func (m *CSRs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CSRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CSRs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CSRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSRs.Merge(m, src)
}
func (m *CSRs) XXX_Size() int {
	return m.Size()
}
func (m *CSRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CSRs.DiscardUnknown(m)
}

var xxx_messageInfo_CSRs proto.InternalMessageInfo

func (m *CSRs) GetCsrs() []*CSR {
	if m != nil {
		return m.Csrs
	}
	return nil
}

func init() {
	proto.RegisterType((*CSR)(nil), "canto.csr.v1.CSR")
	proto.RegisterType((*CSRs)(nil), "canto.csr.v1.CSRs")
}

func init() { proto.RegisterFile("canto/csr/v1/csr.proto", fileDescriptor_57c53cea3d443afa) }

var fileDescriptor_57c53cea3d443afa = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4e, 0xcc, 0x2b,
	0xc9, 0xd7, 0x4f, 0x2e, 0x2e, 0xd2, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x3c, 0x60, 0x71, 0x3d, 0x90, 0x40, 0x99, 0xa1, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x42, 0x1f, 0xc4, 0x82, 0xa8, 0x91, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x87, 0x48,
	0x40, 0x38, 0x10, 0x29, 0xa5, 0x4c, 0x2e, 0x66, 0xe7, 0xe0, 0x20, 0x21, 0x11, 0x2e, 0xd6, 0xfc,
	0xf2, 0xbc, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x86, 0x8b,
	0x33, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x31, 0xb9, 0xa4, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33,
	0x08, 0x21, 0x20, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x12, 0xc4,
	0x94, 0x99, 0x22, 0x24, 0xc1, 0xc5, 0x9e, 0x98, 0x9c, 0x9c, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0x02,
	0x36, 0x05, 0xc6, 0xb5, 0x62, 0x79, 0xb1, 0x40, 0x9e, 0x51, 0x49, 0x97, 0x8b, 0xc5, 0x39, 0x38,
	0xa8, 0x58, 0x48, 0x95, 0x8b, 0x25, 0xb9, 0xb8, 0xa8, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb,
	0x48, 0x50, 0x0f, 0xd9, 0x03, 0x7a, 0xce, 0xc1, 0x41, 0x41, 0x60, 0x69, 0x27, 0xf7, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x06, 0x69, 0xd6, 0xf5, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x86,
	0xf0, 0xf4, 0xcb, 0x8c, 0xf4, 0x2b, 0xc0, 0x01, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0xf6, 0xa9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x93, 0xff, 0xef, 0x0e, 0x42, 0x01, 0x00, 0x00,
}

func (this *CSR) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CSR)
	if !ok {
		that2, ok := that.(CSR)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if len(this.Contracts) != len(that1.Contracts) {
		return false
	}
	for i := range this.Contracts {
		if this.Contracts[i] != that1.Contracts[i] {
			return false
		}
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Account != that1.Account {
		return false
	}
	return true
}
func (m *CSR) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSR) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CSR) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintCsr(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x22
	}
	if m.Id != 0 {
		i = encodeVarintCsr(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintCsr(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCsr(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CSRs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSRs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CSRs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Csrs) > 0 {
		for iNdEx := len(m.Csrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Csrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCsr(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCsr(dAtA []byte, offset int, v uint64) int {
	offset -= sovCsr(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CSR) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCsr(uint64(l))
	}
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovCsr(uint64(l))
		}
	}
	if m.Id != 0 {
		n += 1 + sovCsr(uint64(m.Id))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovCsr(uint64(l))
	}
	return n
}

func (m *CSRs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Csrs) > 0 {
		for _, e := range m.Csrs {
			l = e.Size()
			n += 1 + l + sovCsr(uint64(l))
		}
	}
	return n
}

func sovCsr(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCsr(x uint64) (n int) {
	return sovCsr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CSR) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsr
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
			return fmt.Errorf("proto: CSR: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSR: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCsr
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
func (m *CSRs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsr
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
			return fmt.Errorf("proto: CSRs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSRs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csrs = append(m.Csrs, &CSR{})
			if err := m.Csrs[len(m.Csrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCsr
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
func skipCsr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCsr
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
					return 0, ErrIntOverflowCsr
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
					return 0, ErrIntOverflowCsr
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
				return 0, ErrInvalidLengthCsr
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCsr
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCsr
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCsr        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCsr          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCsr = fmt.Errorf("proto: unexpected end of group")
)
