package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/tmsdkeys/leaders/testutil/keeper"
	"github.com/tmsdkeys/leaders/testutil/nullify"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

func TestTopRankedQuery(t *testing.T) {
	keeper, ctx := keepertest.LeadersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTopRanked(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTopRankedRequest
		response *types.QueryGetTopRankedResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTopRankedRequest{},
			response: &types.QueryGetTopRankedResponse{TopRanked: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TopRanked(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
