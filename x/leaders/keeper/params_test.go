package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tmsdkeys/leaders/testutil/keeper"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LeadersKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
