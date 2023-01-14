package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	var startingProgramId uint64 = 1
	genesisState := DefaultGenesisState()
	require.Equal(t, genesisState.StartingProgramId, startingProgramId)

	// TODO add String equal
}
