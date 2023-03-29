// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gov/v1beta1/genesis.proto

package v1beta1

import (
	fmt "fmt"
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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// deposits defines all the deposits present at genesis.
	Deposits Deposits `protobuf:"bytes,2,rep,name=deposits,proto3,castrepeated=Deposits" json:"deposits"`
	// sealeddeposits defines all the sealed deposits present at genesis.
	SealedDeposits SealedDeposits `protobuf:"bytes,8,rep,name=sealed_deposits,json=sealedDeposits,proto3,castrepeated=SealedDeposits" json:"sealed_deposits"`
	// votes defines all the votes present at genesis.
	Votes Votes `protobuf:"bytes,3,rep,name=votes,proto3,castrepeated=Votes" json:"votes"`
	// proposals defines all the proposals present at genesis.
	Proposals Proposals `protobuf:"bytes,4,rep,name=proposals,proto3,castrepeated=Proposals" json:"proposals"`
	// sealedproposals defines all the sealed proposals present at genesis.
	SealedProposals SealedProposals `protobuf:"bytes,9,rep,name=sealed_proposals,json=sealedProposals,proto3,castrepeated=SealedProposals" json:"sealed_proposals"`
	// params defines all the paramaters of related to deposit.
	DepositParams DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params"`
	// params defines all the paramaters of related to voting.
	VotingParams VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params"`
	// params defines all the paramaters of related to tally.
	TallyParams TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43cd825e0fa7a627, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetDeposits() Deposits {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetSealedDeposits() SealedDeposits {
	if m != nil {
		return m.SealedDeposits
	}
	return nil
}

func (m *GenesisState) GetVotes() Votes {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() Proposals {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetSealedProposals() SealedProposals {
	if m != nil {
		return m.SealedProposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return DepositParams{}
}

func (m *GenesisState) GetVotingParams() VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return VotingParams{}
}

func (m *GenesisState) GetTallyParams() TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return TallyParams{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.gov.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("cosmos/gov/v1beta1/genesis.proto", fileDescriptor_43cd825e0fa7a627) }

var fileDescriptor_43cd825e0fa7a627 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0x94, 0x40,
	0x1c, 0xc6, 0x17, 0xbb, 0xd4, 0xdd, 0xe9, 0xee, 0xb6, 0x4e, 0x1a, 0x25, 0xb5, 0x61, 0xb1, 0xa7,
	0x5e, 0x84, 0x56, 0xcf, 0x5e, 0x88, 0x89, 0x56, 0xa3, 0x69, 0xa8, 0xf1, 0xe0, 0x85, 0x0c, 0x65,
	0x44, 0x22, 0xdb, 0x3f, 0xe1, 0x3f, 0x12, 0xfb, 0x16, 0x3e, 0x87, 0x4f, 0xd2, 0x83, 0x87, 0x1e,
	0x3d, 0xa9, 0xd9, 0x7d, 0x11, 0xc3, 0xcc, 0xc0, 0x42, 0xca, 0x9e, 0x60, 0xbe, 0xff, 0x37, 0x3f,
	0x98, 0x2f, 0xf3, 0x11, 0xe7, 0x12, 0x70, 0x01, 0xe8, 0x25, 0x50, 0x7a, 0xe5, 0x69, 0xc4, 0x05,
	0x3b, 0xf5, 0x12, 0x7e, 0xc5, 0x31, 0x45, 0x37, 0x2f, 0x40, 0x00, 0xa5, 0xca, 0xe1, 0x26, 0x50,
	0xba, 0xda, 0x71, 0xb0, 0x9f, 0x40, 0x02, 0x72, 0xec, 0x55, 0x6f, 0xca, 0x79, 0x70, 0xd8, 0xc7,
	0x82, 0x52, 0x4d, 0x8f, 0x7e, 0x99, 0x64, 0xf2, 0x4a, 0x91, 0x2f, 0x04, 0x13, 0x9c, 0x9e, 0x90,
	0x7d, 0x14, 0xac, 0x10, 0xe9, 0x55, 0x12, 0xe6, 0x05, 0xe4, 0x80, 0x2c, 0x0b, 0xd3, 0xd8, 0x32,
	0x1c, 0xe3, 0x78, 0x18, 0xd0, 0x7a, 0x76, 0xae, 0x47, 0x67, 0x31, 0x3d, 0x23, 0xa3, 0x98, 0xe7,
	0x80, 0xa9, 0x40, 0xeb, 0x9e, 0xb3, 0x75, 0xbc, 0xf3, 0xec, 0xb1, 0x7b, 0xf7, 0xef, 0xdc, 0x97,
	0xca, 0xe3, 0xef, 0xdd, 0xfc, 0x99, 0x0f, 0x7e, 0xfe, 0x9d, 0x8f, 0xb4, 0x80, 0x41, 0xb3, 0x9d,
	0x46, 0x64, 0x17, 0x39, 0xcb, 0x78, 0x1c, 0x36, 0xc4, 0x91, 0x24, 0x3e, 0xe9, 0x23, 0x5e, 0x48,
	0x6b, 0xcd, 0x7d, 0xa8, 0xb9, 0xb3, 0x8e, 0x8c, 0xc1, 0x0c, 0x3b, 0x6b, 0xfa, 0x82, 0x98, 0x25,
	0x08, 0x8e, 0xd6, 0x96, 0x24, 0x5b, 0x7d, 0xe4, 0x8f, 0x20, 0xb8, 0x3f, 0xd5, 0x40, 0xb3, 0x5a,
	0x61, 0xa0, 0x76, 0xd1, 0x77, 0x64, 0x5c, 0xc7, 0x82, 0xd6, 0x50, 0x22, 0x0e, 0xfb, 0x10, 0x75,
	0x40, 0xfe, 0x03, 0x8d, 0x19, 0xd7, 0x0a, 0x06, 0x6b, 0x02, 0xfd, 0x4c, 0xf6, 0xf4, 0x89, 0xd7,
	0xd4, 0xb1, 0xa4, 0x1e, 0x6d, 0x3e, 0x72, 0xc3, 0x7e, 0xa4, 0xd9, 0xbb, 0x5d, 0x1d, 0x03, 0x1d,
	0x63, 0x23, 0xd0, 0xf7, 0x64, 0xa6, 0x23, 0x0d, 0x73, 0x56, 0xb0, 0x05, 0x5a, 0xa6, 0x63, 0x6c,
	0x0a, 0x56, 0x67, 0x75, 0x2e, 0x8d, 0xfe, 0xb0, 0xfa, 0x48, 0x30, 0x8d, 0xdb, 0x22, 0x7d, 0x4b,
	0xa6, 0x25, 0xa8, 0x4b, 0xa2, 0x70, 0xdb, 0x12, 0xe7, 0x6c, 0x48, 0xb3, 0xba, 0x31, 0x6d, 0xda,
	0xa4, 0x6c, 0x69, 0xf4, 0x35, 0x99, 0x08, 0x96, 0x65, 0xd7, 0x35, 0xeb, 0xbe, 0x64, 0xcd, 0xfb,
	0x58, 0x1f, 0x2a, 0x5f, 0x07, 0xb5, 0x23, 0x5a, 0xd2, 0x9b, 0x9b, 0xa5, 0x6d, 0xdc, 0x2e, 0x6d,
	0xe3, 0xdf, 0xd2, 0x36, 0x7e, 0xac, 0xec, 0xc1, 0xed, 0xca, 0x1e, 0xfc, 0x5e, 0xd9, 0x83, 0x4f,
	0x27, 0x49, 0x2a, 0xbe, 0x7c, 0x8b, 0xdc, 0x4b, 0x58, 0x78, 0xba, 0x11, 0xea, 0xf1, 0x14, 0xe3,
	0xaf, 0xde, 0x77, 0x59, 0x0f, 0x71, 0x9d, 0x73, 0xac, 0x4b, 0x12, 0x6d, 0xcb, 0x86, 0x3c, 0xff,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x65, 0x1a, 0x3b, 0xc4, 0x8d, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SealedProposals) > 0 {
		for iNdEx := len(m.SealedProposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SealedProposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.SealedDeposits) > 0 {
		for iNdEx := len(m.SealedDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SealedDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.DepositParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.VotingParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SealedDeposits) > 0 {
		for _, e := range m.SealedDeposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SealedProposals) > 0 {
		for _, e := range m.SealedProposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SealedDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SealedDeposits = append(m.SealedDeposits, SealedDeposit{})
			if err := m.SealedDeposits[len(m.SealedDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SealedProposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SealedProposals = append(m.SealedProposals, SealedProposal{})
			if err := m.SealedProposals[len(m.SealedProposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
