package leaders_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tmsdkeys/leaders/testutil/keeper"
	"github.com/tmsdkeys/leaders/testutil/nullify"
	"github.com/tmsdkeys/leaders/x/leaders"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		TopRanked: &types.TopRanked{
			ClientId: "62",
			Address:  "74",
			Score:    "2",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LeadersKeeper(t)
	leaders.InitGenesis(ctx, *k, genesisState)
	got := leaders.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.Equal(t, genesisState.TopRanked, got.TopRanked)
	// this line is used by starport scaffolding # genesis/test/assert
}
