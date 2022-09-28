package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tmsdkeys/leaders/testutil/keeper"
	"github.com/tmsdkeys/leaders/testutil/nullify"
	"github.com/tmsdkeys/leaders/x/leaders/keeper"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

func createTestTopRanked(keeper *keeper.Keeper, ctx sdk.Context) types.TopRanked {
	item := types.TopRanked{}
	keeper.SetTopRanked(ctx, item)
	return item
}

func TestTopRankedGet(t *testing.T) {
	keeper, ctx := keepertest.LeadersKeeper(t)
	item := createTestTopRanked(keeper, ctx)
	rst, found := keeper.GetTopRanked(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTopRankedRemove(t *testing.T) {
	keeper, ctx := keepertest.LeadersKeeper(t)
	createTestTopRanked(keeper, ctx)
	keeper.RemoveTopRanked(ctx)
	_, found := keeper.GetTopRanked(ctx)
	require.False(t, found)
}
