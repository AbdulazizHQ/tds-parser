package tdsparser_test

import (
	"testing"

	"github.com/AbdulazizHQ/tdsparser"
)

func TestPacketStringer(t *testing.T) {
	want := "Header: [ SQL_BATCH | NORMAL | 10 | 20 | 1 | 2 ]\nPayload length: 5"

	header := tdsparser.TDSHeader{
		tdsparser.SQL_BATCH,
		tdsparser.NORMAL,
		[2]byte{0, 0xA},
		[2]byte{0, 0x14},
		1,
		2,
	}
	packet := tdsparser.TDSPacket{
		&header,
		[]byte{1, 2, 3, 4, 5},
	}

	output := packet.String()
	if want != output {
		t.Fatalf("Expected: %s, but was: %s", want, output)
	}
}
