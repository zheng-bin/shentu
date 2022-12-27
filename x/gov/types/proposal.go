package types

import (
	"fmt"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	proto "github.com/gogo/protobuf/proto"

	certtypes "github.com/shentufoundation/shentu/v2/x/cert/types"
	shieldtypes "github.com/shentufoundation/shentu/v2/x/shield/types"
)

// Proposal types implement UnpackInterfaceMessages to unpack
// Content fields.
var _ types.UnpackInterfacesMessage = Proposal{}
var _ types.UnpackInterfacesMessage = Proposals{}

// NewProposal creates a new Proposal instance
func NewProposal(content govtypes.Content, id uint64, proposerAddress sdk.AccAddress, submitTime time.Time, depositEndTime time.Time) (Proposal, error) {
	p := Proposal{
		ProposalId:       id,
		Status:           govtypes.StatusDepositPeriod,
		ProposerAddress:  proposerAddress.String(),
		FinalTallyResult: govtypes.EmptyTallyResult(),
		TotalDeposit:     sdk.NewCoins(),
		SubmitTime:       submitTime,
		DepositEndTime:   depositEndTime,
	}

	msg, ok := content.(proto.Message)
	if !ok {
		return Proposal{}, fmt.Errorf("%T does not implement proto.Message", content)
	}

	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return Proposal{}, err
	}

	p.Content = any

	return p, nil
}

// GetContent returns the proposal Content
func (p Proposal) GetContent() govtypes.Content {
	content, ok := p.Content.GetCachedValue().(govtypes.Content)
	if !ok {
		return nil
	}
	return content
}

// String implements stringer interface
func (p Proposal) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content govtypes.Content
	return unpacker.UnpackAny(p.Content, &content)
}

func (p Proposal) GetTitle() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.GetTitle()
}

func (p Proposal) ProposalType() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalType()
}

func (p Proposal) ProposalRoute() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalRoute()
}

// HasSecurityVoting returns true if the proposal needs to go through security
// (certifier) voting before stake (validator) voting.
func (p Proposal) HasSecurityVoting() bool {
	switch p.GetContent().(type) {
	case *upgradetypes.SoftwareUpgradeProposal, *certtypes.CertifierUpdateProposal, shieldtypes.ShieldClaimProposal:
		return true
	default:
		return false
	}
}

// Proposals is an array of proposals.
type Proposals []Proposal

// String implements stringer interface.
func (p Proposals) String() string {
	out := "ID - (Status) [Type] Title\n"
	for _, prop := range p {
		out += fmt.Sprintf("%d - (%s) [%s] %s\n",
			prop.ProposalId, prop.Status,
			prop.ProposalType(), prop.GetTitle())
	}
	return strings.TrimSpace(out)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposals) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, x := range p {
		err := x.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}
