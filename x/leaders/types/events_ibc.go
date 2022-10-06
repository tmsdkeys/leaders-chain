package types

// IBC events
const (
	EventTypeTimeout          = "timeout"
	EventTypeTopRankPacket    = "topRank_packet"
	EventTypeGameResultPacket = "gameResult_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
