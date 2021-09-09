package tdsparser_test

import (
	"testing"

	"github.com/AbdulazizHQ/tdsparser"
)

func TestHeaderStringer(t *testing.T) {
	want := "[ SQL_BATCH | NORMAL | 10 | 20 | 1 | 2 ]"

	header := tdsparser.TDSHeader{
		tdsparser.SQL_BATCH,
		tdsparser.NORMAL,
		[2]byte{0, 0xA},
		[2]byte{0, 0x14},
		1,
		2,
	}

	output := header.String()
	if want != output {
		t.Fatalf("Expected: %s, but was: %s", want, output)
	}
}
