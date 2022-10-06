package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

// OnRecvGameResultPacket processes packet reception
func (k Keeper) OnRecvGameResultPacket(ctx sdk.Context, packet channeltypes.Packet, data types.GameResultPacketData) (packetAck types.GameResultPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic
	topRank, found := k.GetTopRanked(ctx)
	if !found {
		k.NewLeader(ctx, data)
	}
	topScore, err := strconv.ParseUint(topRank.Score, 10, 64)
	if err != nil {
		return types.GameResultPacketAck{GameId: data.GameId}, err
	}

	if data.WinCount > topScore {
		k.NewLeader(ctx, data)
	}

	return packetAck, nil
}

func (k Keeper) NewLeader(ctx sdk.Context, data types.GameResultPacketData) error {
	newTopRank := types.TopRanked{
		Address: data.Address,
		Score:   strconv.FormatUint(data.WinCount, 10),
	}
	k.SetTopRanked(ctx, newTopRank)
	return nil
}
