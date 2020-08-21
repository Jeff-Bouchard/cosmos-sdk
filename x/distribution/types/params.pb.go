// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/distribution/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the set of params for the distribution module.
type Params struct {
	CommunityTax        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=community_tax,json=communityTax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_tax" yaml:"community_tax"`
	BaseProposerReward  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=base_proposer_reward,json=baseProposerReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_proposer_reward" yaml:"base_proposer_reward"`
	BonusProposerReward github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=bonus_proposer_reward,json=bonusProposerReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bonus_proposer_reward" yaml:"bonus_proposer_reward"`
	WithdrawAddrEnabled bool                                   `protobuf:"varint,4,opt,name=withdraw_addr_enabled,json=withdrawAddrEnabled,proto3" json:"withdraw_addr_enabled,omitempty" yaml:"withdraw_addr_enabled"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10c7a20ed4bc694, []int{0}
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

func (m *Params) GetWithdrawAddrEnabled() bool {
	if m != nil {
		return m.WithdrawAddrEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "cosmos.distribution.v1beta1.Params")
}

func init() {
	proto.RegisterFile("cosmos/distribution/v1beta1/params.proto", fileDescriptor_a10c7a20ed4bc694)
}

var fileDescriptor_a10c7a20ed4bc694 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x86, 0xa8, 0xd4, 0x43, 0x56, 0xa9, 0x07, 0x55, 0x29, 0x25,
	0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa7, 0x0f, 0x62, 0x41, 0xb4, 0x28, 0x5d, 0x67, 0xe6, 0x62,
	0x0b, 0x00, 0x9b, 0x21, 0x94, 0xcd, 0xc5, 0x9b, 0x9c, 0x9f, 0x9b, 0x5b, 0x9a, 0x97, 0x59, 0x52,
	0x19, 0x5f, 0x92, 0x58, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xe4, 0x76, 0xe2, 0x9e, 0x3c,
	0xc3, 0xad, 0x7b, 0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x50, 0x17, 0x41, 0x28, 0xdd, 0xe2, 0x94, 0x6c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0x97,
	0xd4, 0xe4, 0x4f, 0xf7, 0xe4, 0x45, 0x2a, 0x13, 0x73, 0x73, 0xac, 0x94, 0x50, 0x0c, 0x53, 0x0a,
	0xe2, 0x81, 0xf3, 0x43, 0x12, 0x2b, 0x84, 0xea, 0xb9, 0x44, 0x92, 0x12, 0x8b, 0x53, 0xe3, 0x0b,
	0x8a, 0xf2, 0x0b, 0xf2, 0x8b, 0x53, 0x8b, 0xe2, 0x8b, 0x52, 0xcb, 0x13, 0x8b, 0x52, 0x24, 0x98,
	0xc0, 0x76, 0xfa, 0x92, 0x6c, 0xa7, 0x34, 0xc4, 0x4e, 0x6c, 0x66, 0x2a, 0x05, 0x09, 0x81, 0x84,
	0x03, 0xa0, 0xa2, 0x41, 0x60, 0x41, 0xa1, 0x26, 0x46, 0x2e, 0xd1, 0xa4, 0xfc, 0xbc, 0xd2, 0x62,
	0x0c, 0x27, 0x30, 0x83, 0x9d, 0xe0, 0x47, 0xb2, 0x13, 0x64, 0xa0, 0x4e, 0xc0, 0x66, 0xa8, 0x52,
	0x90, 0x30, 0x58, 0x1c, 0xcd, 0x11, 0x21, 0x5c, 0xa2, 0xe5, 0x99, 0x25, 0x19, 0x29, 0x45, 0x89,
	0xe5, 0xf1, 0x89, 0x29, 0x29, 0x45, 0xf1, 0xa9, 0x79, 0x89, 0x49, 0x39, 0xa9, 0x29, 0x12, 0x2c,
	0x0a, 0x8c, 0x1a, 0x1c, 0x4e, 0x0a, 0x08, 0x53, 0xb1, 0x2a, 0x53, 0x0a, 0x12, 0x86, 0x89, 0x3b,
	0xa6, 0xa4, 0x14, 0xb9, 0x42, 0x44, 0xad, 0x58, 0x66, 0x2c, 0x90, 0x67, 0x70, 0xf2, 0x3e, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x43, 0xbc, 0x5e, 0xaa, 0x40, 0x4d, 0x68,
	0x60, 0x1f, 0x26, 0xb1, 0x81, 0x53, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x25, 0xf1,
	0x84, 0x8c, 0x02, 0x00, 0x00,
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
	if m.WithdrawAddrEnabled {
		i--
		if m.WithdrawAddrEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.BonusProposerReward.Size()
		i -= size
		if _, err := m.BonusProposerReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BaseProposerReward.Size()
		i -= size
		if _, err := m.BaseProposerReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.CommunityTax.Size()
		i -= size
		if _, err := m.CommunityTax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	l = m.CommunityTax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BaseProposerReward.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BonusProposerReward.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.WithdrawAddrEnabled {
		n += 2
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityTax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityTax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseProposerReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseProposerReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusProposerReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BonusProposerReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddrEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithdrawAddrEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
