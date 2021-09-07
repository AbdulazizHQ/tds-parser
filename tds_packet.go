package tdsparser

import "fmt"

type TDSPacket struct {
	header  TDSHeader
	payload []byte
}

func (packet TDSPacket) String() string {
	return fmt.Sprintf("Header: %s\nPayload length: %d", packet.header, len(packet.payload))
}
