package types

// ValidateBasic is used for validating the packet
func (p GameResultPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p GameResultPacketData) GetBytes() ([]byte, error) {
	var modulePacket LeadersPacketData

	modulePacket.Packet = &LeadersPacketData_GameResultPacket{&p}

	return modulePacket.Marshal()
}
