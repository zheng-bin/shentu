// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1alpha1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/gov/types"
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
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty" yaml:"starting_proposal_id"`
	// deposits defines all the deposits present at genesis.
	Deposits []types.Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits"`
	// votes defines all the votes present at genesis.
	Votes []types.Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes"`
	// proposals defines all the proposals present at genesis.
	Proposals []types.Proposal `protobuf:"bytes,4,rep,name=proposals,proto3" json:"proposals"`
	// params defines all the parameters of related to deposit.
	DepositParams types.DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params" yaml:"deposit_params"`
	// params defines all the parameters of related to voting.
	VotingParams types.VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params" yaml:"voting_params"`
	// params defines all the parameters of related to tally.
	TallyParams types.TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params" yaml:"tally_params"`
	// params defines all the parameters of related to custom.
	CustomParams CustomParams `protobuf:"bytes,8,opt,name=custom_params,json=customParams,proto3" json:"custom_params" yaml:"custom_params"`
	// proposals that require and have passed cert votes.
	CertVotedProposalIds []uint64 `protobuf:"varint,9,rep,packed,name=cert_voted_proposal_ids,json=certVotedProposalIds,proto3" json:"cert_voted_proposal_ids,omitempty" yaml:"cert_voted_proposal_ids"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d2eda471d71961d, []int{0}
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

func (m *GenesisState) GetDeposits() []types.Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() []types.Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() []types.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() types.DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return types.DepositParams{}
}

func (m *GenesisState) GetVotingParams() types.VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return types.VotingParams{}
}

func (m *GenesisState) GetTallyParams() types.TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return types.TallyParams{}
}

func (m *GenesisState) GetCustomParams() CustomParams {
	if m != nil {
		return m.CustomParams
	}
	return CustomParams{}
}

func (m *GenesisState) GetCertVotedProposalIds() []uint64 {
	if m != nil {
		return m.CertVotedProposalIds
	}
	return nil
}

type CustomParams struct {
	CertifierUpdateSecurityVoteTally *types.TallyParams `protobuf:"bytes,1,opt,name=certifier_update_security_vote_tally,json=certifierUpdateSecurityVoteTally,proto3" json:"certifier_update_security_vote_tally,omitempty"`
	CertifierUpdateStakeVoteTally    *types.TallyParams `protobuf:"bytes,2,opt,name=certifier_update_stake_vote_tally,json=certifierUpdateStakeVoteTally,proto3" json:"certifier_update_stake_vote_tally,omitempty"`
}

func (m *CustomParams) Reset()         { *m = CustomParams{} }
func (m *CustomParams) String() string { return proto.CompactTextString(m) }
func (*CustomParams) ProtoMessage()    {}
func (*CustomParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d2eda471d71961d, []int{1}
}
func (m *CustomParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustomParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustomParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustomParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomParams.Merge(m, src)
}
func (m *CustomParams) XXX_Size() int {
	return m.Size()
}
func (m *CustomParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomParams.DiscardUnknown(m)
}

var xxx_messageInfo_CustomParams proto.InternalMessageInfo

func (m *CustomParams) GetCertifierUpdateSecurityVoteTally() *types.TallyParams {
	if m != nil {
		return m.CertifierUpdateSecurityVoteTally
	}
	return nil
}

func (m *CustomParams) GetCertifierUpdateStakeVoteTally() *types.TallyParams {
	if m != nil {
		return m.CertifierUpdateStakeVoteTally
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "shentu.gov.v1alpha1.GenesisState")
	proto.RegisterType((*CustomParams)(nil), "shentu.gov.v1alpha1.CustomParams")
}

func init() { proto.RegisterFile("shentu/gov/v1alpha1/genesis.proto", fileDescriptor_1d2eda471d71961d) }

var fileDescriptor_1d2eda471d71961d = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x8f, 0xd2, 0x40,
	0x1c, 0xc5, 0xe9, 0xc2, 0xe2, 0xee, 0x00, 0x1e, 0x66, 0x31, 0x36, 0x0b, 0xdb, 0x96, 0xc6, 0x03,
	0xa7, 0x36, 0xbb, 0x7a, 0x32, 0x31, 0x31, 0x68, 0xa2, 0xde, 0xd6, 0xae, 0x9a, 0xe8, 0xa5, 0x19,
	0xda, 0xd9, 0xd2, 0x08, 0x4c, 0xd3, 0x99, 0x36, 0xf2, 0x1f, 0x78, 0xf4, 0xcf, 0xda, 0xe3, 0x1e,
	0x3d, 0x11, 0x03, 0x37, 0x8f, 0xfc, 0x05, 0x66, 0x7e, 0x14, 0x8a, 0x56, 0xdd, 0x1b, 0x99, 0x79,
	0xef, 0x7d, 0xde, 0x7c, 0x3b, 0x0c, 0x18, 0xd0, 0x09, 0x9e, 0xb3, 0xcc, 0x8d, 0x48, 0xee, 0xe6,
	0xe7, 0x68, 0x9a, 0x4c, 0xd0, 0xb9, 0x1b, 0xe1, 0x39, 0xa6, 0x31, 0x75, 0x92, 0x94, 0x30, 0x02,
	0x4f, 0xa4, 0xc4, 0x89, 0x48, 0xee, 0x14, 0x92, 0xd3, 0x6e, 0x44, 0x22, 0x22, 0xf6, 0x5d, 0xfe,
	0x4b, 0x4a, 0x4f, 0xfb, 0x01, 0xa1, 0x33, 0x42, 0x55, 0xda, 0x18, 0x33, 0x1e, 0x46, 0x72, 0xb9,
	0x6b, 0x7f, 0x6d, 0x82, 0xf6, 0x2b, 0x19, 0x7d, 0xc5, 0x10, 0xc3, 0xf0, 0x2d, 0xe8, 0x52, 0x86,
	0x52, 0x16, 0xcf, 0x23, 0x3f, 0x49, 0x49, 0x42, 0x28, 0x9a, 0xfa, 0x71, 0xa8, 0x6b, 0x96, 0x36,
	0x6c, 0x8c, 0xcc, 0xcd, 0xd2, 0xec, 0x2d, 0xd0, 0x6c, 0xfa, 0xd4, 0xae, 0x52, 0xd9, 0x1e, 0x2c,
	0x96, 0x2f, 0xd5, 0xea, 0x9b, 0x10, 0x3e, 0x03, 0x47, 0x21, 0x4e, 0x08, 0x8d, 0x19, 0xd5, 0x0f,
	0xac, 0xfa, 0xb0, 0x75, 0xd1, 0x73, 0x64, 0x29, 0xd5, 0x5f, 0x94, 0x72, 0x5e, 0x4a, 0xcd, 0xa8,
	0x71, 0xb3, 0x34, 0x6b, 0xde, 0xd6, 0x02, 0x9f, 0x80, 0xc3, 0x9c, 0x30, 0x4c, 0xf5, 0xba, 0xf0,
	0xea, 0x55, 0xde, 0x0f, 0x84, 0x61, 0x65, 0x94, 0x62, 0xf8, 0x1c, 0x1c, 0x17, 0xc5, 0xa8, 0xde,
	0x10, 0xce, 0x7e, 0x95, 0xb3, 0xe8, 0xa9, 0xdc, 0x3b, 0x13, 0x8c, 0xc0, 0x7d, 0xd5, 0xc1, 0x4f,
	0x50, 0x8a, 0x66, 0x54, 0x3f, 0xb4, 0xb4, 0x61, 0xeb, 0x62, 0xf0, 0x8f, 0xf2, 0x97, 0x42, 0x38,
	0x3a, 0xe3, 0x59, 0x9b, 0xa5, 0xf9, 0x40, 0x8e, 0x6a, 0x3f, 0xc6, 0xf6, 0x3a, 0x61, 0x59, 0x0d,
	0x03, 0xd0, 0xc9, 0x89, 0x1c, 0xa5, 0xe4, 0x34, 0x05, 0xc7, 0xfa, 0xcb, 0x41, 0xf9, 0x70, 0x25,
	0xa6, 0xaf, 0x30, 0x5d, 0x89, 0xd9, 0x0b, 0xb1, 0xbd, 0x76, 0x5e, 0xd2, 0x42, 0x1f, 0xb4, 0x19,
	0x9a, 0x4e, 0x17, 0x05, 0xe3, 0x9e, 0x60, 0x98, 0x55, 0x8c, 0x77, 0x5c, 0xa7, 0x10, 0x3d, 0x85,
	0x38, 0x91, 0x88, 0x72, 0x84, 0xed, 0xb5, 0xd8, 0x4e, 0x09, 0x43, 0xd0, 0x09, 0x32, 0xca, 0xc8,
	0xac, 0x20, 0x1c, 0xa9, 0x69, 0x55, 0x5c, 0x55, 0xe7, 0x85, 0x50, 0x56, 0x1f, 0x63, 0x2f, 0xc5,
	0xf6, 0xda, 0x41, 0x49, 0x0b, 0x3f, 0x82, 0x87, 0x01, 0x4e, 0x99, 0xcf, 0x3f, 0x72, 0x58, 0xbe,
	0x7a, 0x54, 0x3f, 0xb6, 0xea, 0xc3, 0xc6, 0xc8, 0xde, 0x2c, 0x4d, 0x43, 0x05, 0x55, 0x0b, 0x6d,
	0xaf, 0xcb, 0x77, 0xf8, 0x95, 0x09, 0x77, 0xb7, 0x94, 0xda, 0x3f, 0x35, 0xd0, 0x2e, 0xf7, 0x82,
	0x04, 0x3c, 0xe2, 0xc2, 0xf8, 0x3a, 0xc6, 0xa9, 0x9f, 0x25, 0x21, 0x62, 0xd8, 0xa7, 0x38, 0xc8,
	0xd2, 0x98, 0x2d, 0x44, 0xae, 0x2f, 0xce, 0x2f, 0xfe, 0x1a, 0xff, 0x1f, 0xa5, 0x67, 0x6d, 0xc3,
	0xde, 0x8b, 0xac, 0x2b, 0x15, 0xc5, 0x8b, 0x08, 0x1d, 0x8c, 0xc1, 0xe0, 0x4f, 0x20, 0x43, 0x9f,
	0x71, 0x99, 0x76, 0x70, 0x37, 0xda, 0xd9, 0xef, 0x34, 0x9e, 0xb3, 0x45, 0x8d, 0x5e, 0xdf, 0xac,
	0x0c, 0xed, 0x76, 0x65, 0x68, 0x3f, 0x56, 0x86, 0xf6, 0x6d, 0x6d, 0xd4, 0x6e, 0xd7, 0x46, 0xed,
	0xfb, 0xda, 0xa8, 0x7d, 0x72, 0xa2, 0x98, 0x4d, 0xb2, 0xb1, 0x13, 0x90, 0x99, 0x2b, 0x3f, 0xdd,
	0x35, 0xc9, 0xe6, 0x21, 0x62, 0x31, 0x99, 0xab, 0x05, 0xf7, 0x8b, 0x78, 0x4d, 0xd8, 0x22, 0xc1,
	0x74, 0xdc, 0x14, 0x0f, 0xc9, 0xe3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xde, 0x65, 0xfc,
	0xb6, 0x04, 0x00, 0x00,
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
	if len(m.CertVotedProposalIds) > 0 {
		dAtA2 := make([]byte, len(m.CertVotedProposalIds)*10)
		var j1 int
		for _, num := range m.CertVotedProposalIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesis(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x4a
	}
	{
		size, err := m.CustomParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
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

func (m *CustomParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustomParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertifierUpdateStakeVoteTally != nil {
		{
			size, err := m.CertifierUpdateStakeVoteTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CertifierUpdateSecurityVoteTally != nil {
		{
			size, err := m.CertifierUpdateSecurityVoteTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
	l = m.CustomParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.CertVotedProposalIds) > 0 {
		l = 0
		for _, e := range m.CertVotedProposalIds {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	return n
}

func (m *CustomParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CertifierUpdateSecurityVoteTally != nil {
		l = m.CertifierUpdateSecurityVoteTally.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.CertifierUpdateStakeVoteTally != nil {
		l = m.CertifierUpdateStakeVoteTally.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
			m.Deposits = append(m.Deposits, types.Deposit{})
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
			m.Votes = append(m.Votes, types.Vote{})
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
			m.Proposals = append(m.Proposals, types.Proposal{})
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
				return fmt.Errorf("proto: wrong wireType = %d for field CustomParams", wireType)
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
			if err := m.CustomParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				m.CertVotedProposalIds = append(m.CertVotedProposalIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
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
				if elementCount != 0 && len(m.CertVotedProposalIds) == 0 {
					m.CertVotedProposalIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
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
					m.CertVotedProposalIds = append(m.CertVotedProposalIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CertVotedProposalIds", wireType)
			}
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
func (m *CustomParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CustomParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifierUpdateSecurityVoteTally", wireType)
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
			if m.CertifierUpdateSecurityVoteTally == nil {
				m.CertifierUpdateSecurityVoteTally = &types.TallyParams{}
			}
			if err := m.CertifierUpdateSecurityVoteTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifierUpdateStakeVoteTally", wireType)
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
			if m.CertifierUpdateStakeVoteTally == nil {
				m.CertifierUpdateStakeVoteTally = &types.TallyParams{}
			}
			if err := m.CertifierUpdateStakeVoteTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
