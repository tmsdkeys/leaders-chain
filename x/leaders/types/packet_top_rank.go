package types

// ValidateBasic is used for validating the packet
func (p TopRankPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p TopRankPacketData) GetBytes() ([]byte, error) {
	var modulePacket LeadersPacketData

	modulePacket.Packet = &LeadersPacketData_TopRankPacket{&p}

	return modulePacket.Marshal()
}
