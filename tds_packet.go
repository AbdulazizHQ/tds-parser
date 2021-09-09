package tdsparser

import "fmt"

type TDSPacket struct {
	Header  *TDSHeader
	Payload []byte
}

func (packet *TDSPacket) String() string {
	return fmt.Sprintf("Header: %s\nPayload length: %d", packet.Header, len(packet.Payload))
}
