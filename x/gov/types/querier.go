package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// QueryProposalsParams defines data structure for querying 'custom/gov/proposals'.
type QueryProposalsParams struct {
	Page           int
	Limit          int
	Voter          sdk.AccAddress
	Depositor      sdk.AccAddress
	ProposalStatus govtypes.ProposalStatus
}

// NewQueryProposalsParams creates a new instance of QueryProposalsParams.
func NewQueryProposalsParams(page, limit int, status govtypes.ProposalStatus, voter, depositor sdk.AccAddress) QueryProposalsParams {
	return QueryProposalsParams{
		Page:           page,
		Limit:          limit,
		Voter:          voter,
		Depositor:      depositor,
		ProposalStatus: status,
	}
}
