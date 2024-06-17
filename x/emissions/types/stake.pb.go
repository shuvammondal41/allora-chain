// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/v1/stake.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_allora_network_allora_chain_math "github.com/allora-network/allora-chain/math"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type StakePlacement struct {
	BlockRemovalStarted   int64                 `protobuf:"varint,1,opt,name=block_removal_started,json=blockRemovalStarted,proto3" json:"block_removal_started,omitempty"`
	TopicId               uint64                `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Reputer               string                `protobuf:"bytes,3,opt,name=reputer,proto3" json:"reputer,omitempty"`
	Amount                cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	BlockRemovalCompleted int64                 `protobuf:"varint,5,opt,name=block_removal_completed,json=blockRemovalCompleted,proto3" json:"block_removal_completed,omitempty"`
}

func (m *StakePlacement) Reset()         { *m = StakePlacement{} }
func (m *StakePlacement) String() string { return proto.CompactTextString(m) }
func (*StakePlacement) ProtoMessage()    {}
func (*StakePlacement) Descriptor() ([]byte, []int) {
	return fileDescriptor_f03692b073d250c3, []int{0}
}
func (m *StakePlacement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakePlacement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakePlacement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakePlacement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakePlacement.Merge(m, src)
}
func (m *StakePlacement) XXX_Size() int {
	return m.Size()
}
func (m *StakePlacement) XXX_DiscardUnknown() {
	xxx_messageInfo_StakePlacement.DiscardUnknown(m)
}

var xxx_messageInfo_StakePlacement proto.InternalMessageInfo

func (m *StakePlacement) GetBlockRemovalStarted() int64 {
	if m != nil {
		return m.BlockRemovalStarted
	}
	return 0
}

func (m *StakePlacement) GetTopicId() uint64 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *StakePlacement) GetReputer() string {
	if m != nil {
		return m.Reputer
	}
	return ""
}

func (m *StakePlacement) GetBlockRemovalCompleted() int64 {
	if m != nil {
		return m.BlockRemovalCompleted
	}
	return 0
}

type DelegateStakePlacement struct {
	BlockRemovalStarted   int64                 `protobuf:"varint,1,opt,name=block_removal_started,json=blockRemovalStarted,proto3" json:"block_removal_started,omitempty"`
	TopicId               uint64                `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Reputer               string                `protobuf:"bytes,3,opt,name=reputer,proto3" json:"reputer,omitempty"`
	Delegator             string                `protobuf:"bytes,4,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Amount                cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	BlockRemovalCompleted int64                 `protobuf:"varint,6,opt,name=block_removal_completed,json=blockRemovalCompleted,proto3" json:"block_removal_completed,omitempty"`
}

func (m *DelegateStakePlacement) Reset()         { *m = DelegateStakePlacement{} }
func (m *DelegateStakePlacement) String() string { return proto.CompactTextString(m) }
func (*DelegateStakePlacement) ProtoMessage()    {}
func (*DelegateStakePlacement) Descriptor() ([]byte, []int) {
	return fileDescriptor_f03692b073d250c3, []int{1}
}
func (m *DelegateStakePlacement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegateStakePlacement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegateStakePlacement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegateStakePlacement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateStakePlacement.Merge(m, src)
}
func (m *DelegateStakePlacement) XXX_Size() int {
	return m.Size()
}
func (m *DelegateStakePlacement) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateStakePlacement.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateStakePlacement proto.InternalMessageInfo

func (m *DelegateStakePlacement) GetBlockRemovalStarted() int64 {
	if m != nil {
		return m.BlockRemovalStarted
	}
	return 0
}

func (m *DelegateStakePlacement) GetTopicId() uint64 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *DelegateStakePlacement) GetReputer() string {
	if m != nil {
		return m.Reputer
	}
	return ""
}

func (m *DelegateStakePlacement) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *DelegateStakePlacement) GetBlockRemovalCompleted() int64 {
	if m != nil {
		return m.BlockRemovalCompleted
	}
	return 0
}

type DelegatorInfo struct {
	Amount     github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"amount"`
	RewardDebt github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,2,opt,name=reward_debt,json=rewardDebt,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"reward_debt"`
}

func (m *DelegatorInfo) Reset()         { *m = DelegatorInfo{} }
func (m *DelegatorInfo) String() string { return proto.CompactTextString(m) }
func (*DelegatorInfo) ProtoMessage()    {}
func (*DelegatorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f03692b073d250c3, []int{2}
}
func (m *DelegatorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorInfo.Merge(m, src)
}
func (m *DelegatorInfo) XXX_Size() int {
	return m.Size()
}
func (m *DelegatorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StakePlacement)(nil), "emissions.v1.StakePlacement")
	proto.RegisterType((*DelegateStakePlacement)(nil), "emissions.v1.DelegateStakePlacement")
	proto.RegisterType((*DelegatorInfo)(nil), "emissions.v1.DelegatorInfo")
}

func init() { proto.RegisterFile("emissions/v1/stake.proto", fileDescriptor_f03692b073d250c3) }

var fileDescriptor_f03692b073d250c3 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xbf, 0x6e, 0x13, 0x41,
	0x10, 0xc6, 0xbd, 0xf9, 0xe3, 0x90, 0xe5, 0x8f, 0xc4, 0x81, 0xe1, 0x12, 0xa1, 0x8b, 0x95, 0xca,
	0x42, 0xca, 0x1d, 0x01, 0x09, 0xa8, 0x83, 0x0b, 0x5c, 0x81, 0x2e, 0x0d, 0xa2, 0x39, 0xed, 0xdd,
	0x0d, 0xf6, 0xca, 0xb7, 0x3b, 0xa7, 0xdd, 0xb1, 0x03, 0x6f, 0xc1, 0x2b, 0xd0, 0x51, 0x52, 0x50,
	0xf1, 0x04, 0x29, 0x23, 0x2a, 0x44, 0x11, 0x21, 0xbb, 0xe0, 0x21, 0x68, 0x90, 0x77, 0x6d, 0xc7,
	0x34, 0x08, 0x09, 0x21, 0x9a, 0xd3, 0x7d, 0xf3, 0xcd, 0xcd, 0xfd, 0xe6, 0x93, 0x86, 0x87, 0xa0,
	0xa4, 0xb5, 0x12, 0xb5, 0x4d, 0xc6, 0x87, 0x89, 0x25, 0x31, 0x84, 0xb8, 0x36, 0x48, 0x18, 0x5c,
	0x59, 0x3a, 0xf1, 0xf8, 0x70, 0x77, 0xa7, 0x40, 0xab, 0xd0, 0x66, 0xce, 0x4b, 0xbc, 0xf0, 0x8d,
	0xbb, 0xd7, 0x85, 0x92, 0x1a, 0x13, 0xf7, 0x9c, 0x97, 0x6e, 0xf6, 0xb1, 0x8f, 0xbe, 0x75, 0xf6,
	0xe6, 0xab, 0xfb, 0x3f, 0x18, 0xbf, 0x76, 0x3c, 0xfb, 0xc3, 0xf3, 0x4a, 0x14, 0xa0, 0x40, 0x53,
	0x70, 0x9f, 0xb7, 0xf2, 0x0a, 0x8b, 0x61, 0x66, 0x40, 0xe1, 0x58, 0x54, 0x99, 0x25, 0x61, 0x08,
	0xca, 0x90, 0xb5, 0x59, 0x67, 0x3d, 0xbd, 0xe1, 0xcc, 0xd4, 0x7b, 0xc7, 0xde, 0x0a, 0x76, 0xf8,
	0x25, 0xc2, 0x5a, 0x16, 0x99, 0x2c, 0xc3, 0xb5, 0x36, 0xeb, 0x6c, 0xa4, 0x5b, 0x4e, 0xf7, 0xca,
	0x20, 0xe4, 0x5b, 0x06, 0xea, 0x11, 0x81, 0x09, 0xd7, 0xdb, 0xac, 0xb3, 0x9d, 0x2e, 0x64, 0xf0,
	0x94, 0x37, 0x85, 0xc2, 0x91, 0xa6, 0x70, 0x63, 0x66, 0x1c, 0xdd, 0x3b, 0x3d, 0xdf, 0x6b, 0x7c,
	0x3d, 0xdf, 0x6b, 0xf9, 0x55, 0x6c, 0x39, 0x8c, 0x25, 0x26, 0x4a, 0xd0, 0x20, 0xee, 0x69, 0xfa,
	0xfc, 0xf1, 0x80, 0xcf, 0x77, 0xec, 0x69, 0x7a, 0xff, 0xfd, 0xc3, 0x5d, 0x96, 0xce, 0xbf, 0x0f,
	0x1e, 0xf2, 0xdb, 0xbf, 0x22, 0x17, 0xa8, 0xea, 0x0a, 0x66, 0xd0, 0x9b, 0x0e, 0xba, 0xb5, 0x0a,
	0xfd, 0x64, 0x61, 0xee, 0xbf, 0x5b, 0xe3, 0xb7, 0xba, 0x50, 0x41, 0x5f, 0x10, 0xfc, 0xaf, 0x14,
	0xee, 0xf0, 0xed, 0xd2, 0x23, 0xa0, 0xf1, 0x41, 0xa4, 0x17, 0x85, 0x95, 0x8c, 0x36, 0xff, 0x5d,
	0x46, 0xcd, 0xdf, 0x65, 0xf4, 0x89, 0xf1, 0xab, 0xdd, 0x05, 0x4f, 0x4f, 0xbf, 0xc2, 0xe0, 0xd9,
	0x92, 0x89, 0x39, 0xa6, 0x47, 0x73, 0xa6, 0xa4, 0x2f, 0x69, 0x30, 0xca, 0xe3, 0x02, 0x55, 0x22,
	0xaa, 0x0a, 0x8d, 0x38, 0xd0, 0x40, 0x27, 0x68, 0x86, 0x0b, 0x59, 0x0c, 0x84, 0xd4, 0x9e, 0xb6,
	0x0b, 0xc5, 0x12, 0xed, 0x05, 0xbf, 0x6c, 0xe0, 0x44, 0x98, 0x32, 0x2b, 0x21, 0x27, 0x17, 0xdd,
	0x5f, 0x4c, 0xe5, 0x7e, 0x56, 0x17, 0x72, 0x3a, 0x4a, 0x4f, 0x27, 0x11, 0x3b, 0x9b, 0x44, 0xec,
	0xdb, 0x24, 0x62, 0x6f, 0xa7, 0x51, 0xe3, 0x6c, 0x1a, 0x35, 0xbe, 0x4c, 0xa3, 0xc6, 0xcb, 0xc7,
	0x7f, 0x38, 0xf6, 0x75, 0x72, 0x71, 0x8d, 0xf4, 0xa6, 0x06, 0x9b, 0x37, 0xdd, 0xe5, 0x3c, 0xf8,
	0x19, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x91, 0xb8, 0xe5, 0xa7, 0x03, 0x00, 0x00,
}

func (m *StakePlacement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakePlacement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakePlacement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockRemovalCompleted != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.BlockRemovalCompleted))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStake(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Reputer) > 0 {
		i -= len(m.Reputer)
		copy(dAtA[i:], m.Reputer)
		i = encodeVarintStake(dAtA, i, uint64(len(m.Reputer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TopicId != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.TopicId))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockRemovalStarted != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.BlockRemovalStarted))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DelegateStakePlacement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegateStakePlacement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegateStakePlacement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockRemovalCompleted != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.BlockRemovalCompleted))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStake(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintStake(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Reputer) > 0 {
		i -= len(m.Reputer)
		copy(dAtA[i:], m.Reputer)
		i = encodeVarintStake(dAtA, i, uint64(len(m.Reputer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TopicId != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.TopicId))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockRemovalStarted != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.BlockRemovalStarted))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DelegatorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RewardDebt.Size()
		i -= size
		if _, err := m.RewardDebt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStake(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStake(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintStake(dAtA []byte, offset int, v uint64) int {
	offset -= sovStake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakePlacement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockRemovalStarted != 0 {
		n += 1 + sovStake(uint64(m.BlockRemovalStarted))
	}
	if m.TopicId != 0 {
		n += 1 + sovStake(uint64(m.TopicId))
	}
	l = len(m.Reputer)
	if l > 0 {
		n += 1 + l + sovStake(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovStake(uint64(l))
	if m.BlockRemovalCompleted != 0 {
		n += 1 + sovStake(uint64(m.BlockRemovalCompleted))
	}
	return n
}

func (m *DelegateStakePlacement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockRemovalStarted != 0 {
		n += 1 + sovStake(uint64(m.BlockRemovalStarted))
	}
	if m.TopicId != 0 {
		n += 1 + sovStake(uint64(m.TopicId))
	}
	l = len(m.Reputer)
	if l > 0 {
		n += 1 + l + sovStake(uint64(l))
	}
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovStake(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovStake(uint64(l))
	if m.BlockRemovalCompleted != 0 {
		n += 1 + sovStake(uint64(m.BlockRemovalCompleted))
	}
	return n
}

func (m *DelegatorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovStake(uint64(l))
	l = m.RewardDebt.Size()
	n += 1 + l + sovStake(uint64(l))
	return n
}

func sovStake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStake(x uint64) (n int) {
	return sovStake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakePlacement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStake
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
			return fmt.Errorf("proto: StakePlacement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakePlacement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockRemovalStarted", wireType)
			}
			m.BlockRemovalStarted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockRemovalStarted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicId", wireType)
			}
			m.TopicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reputer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reputer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockRemovalCompleted", wireType)
			}
			m.BlockRemovalCompleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockRemovalCompleted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStake
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
func (m *DelegateStakePlacement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStake
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
			return fmt.Errorf("proto: DelegateStakePlacement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegateStakePlacement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockRemovalStarted", wireType)
			}
			m.BlockRemovalStarted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockRemovalStarted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicId", wireType)
			}
			m.TopicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reputer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reputer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockRemovalCompleted", wireType)
			}
			m.BlockRemovalCompleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockRemovalCompleted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStake
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
func (m *DelegatorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStake
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
			return fmt.Errorf("proto: DelegatorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDebt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStake
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
func skipStake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStake
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
					return 0, ErrIntOverflowStake
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
					return 0, ErrIntOverflowStake
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
				return 0, ErrInvalidLengthStake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStake = fmt.Errorf("proto: unexpected end of group")
)
