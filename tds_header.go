package tdsparser

import (
	"encoding/binary"
	"fmt"
)

const HEADER_SIZE = 8 // bytes

type TDSHeader struct {
	Type     byte
	Status   byte
	Length   [2]byte
	SPID     [2]byte
	PacketID byte
	Window   byte
}

func (header TDSHeader) String() string {
	var headerType, status string

	switch header.Type {
	case SQL_BATCH:
		headerType = "SQL_BATCH"
	case PRE_TDS7_LOGIN:
		headerType = "PRE_TDS7_LOGIN"
	case RPC:
		headerType = "RPC"
	case TABULAR_RESULT:
		headerType = "TABLUAR_RESULT"
	case ATTENTION_SIGNAL:
		headerType = "ATTENTION_SIGNAL"
	case BULK_LOAD_DATA:
		headerType = "BULK_LOAD_DATA"
	case FEDERATED_AUTHENTICATION_TOKEN:
		headerType = "FEDERATED_AUTHENTICATION_DATA"
	case TRANSACTION_MANAGER_REQUEST:
		headerType = "TRANSACTION_MANAGER_REQUEST"
	case TDS7_LOGIN:
		headerType = "TDS7_LOGIN"
	case SSPI:
		headerType = "SSPI"
	case PRE_LOGIN:
		headerType = "PRE_LOGIN"
	default:
		panic(fmt.Sprintf("Header: %d is undefined", header.Type))
	}

	switch header.Status {
	case NORMAL:
		status = "NORMAL"
	case EOM:
		status = "EOM"
	case IGNORE:
		status = "IGNORE"
	case RESET_CONNECTION:
		status = "RESET_CONNECTION"
	case RESET_CONNECTION_SKIP_TRAN:
		status = "RESET_CONNECTION_SKIP_TRAN"
	default:
		panic(fmt.Sprintf("Status: %d is undefined", header.Status))
	}

	return fmt.Sprintf("[ %s | %s | %d | %d | %d | %d ]",
		headerType, status, binary.BigEndian.Uint16(header.Length[:]),
		binary.BigEndian.Uint16(header.SPID[:]), header.PacketID, header.Window)
}
