package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tmsdkeys/leaders/testutil/keeper"
	"github.com/tmsdkeys/leaders/x/leaders/keeper"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LeadersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
