package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

func (k msgServer) SendTopRank(goCtx context.Context, msg *types.MsgSendTopRank) (*types.MsgSendTopRankResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.TopRankPacketData

	packet.Address = msg.Address
	packet.Score = msg.Score

	// Transmit the packet
	err := k.TransmitTopRankPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendTopRankResponse{}, nil
}
