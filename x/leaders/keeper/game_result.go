package keeper

import (
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

	return packetAck, nil
}
