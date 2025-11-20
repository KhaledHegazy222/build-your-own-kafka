package request

import (
	"encoding/binary"
)

type Request struct {
	MessageSize       uint32
	RequestAPIKey     uint16
	RequestAPIVersion uint16
	CorrelationID     uint32
	ClientID          string
}

func Parse(payload []byte) (Request, error) {
	req := Request{}

	req.MessageSize = binary.BigEndian.Uint32(payload[:4])
	req.RequestAPIKey = binary.BigEndian.Uint16(payload[4:6])
	req.RequestAPIVersion = binary.BigEndian.Uint16(payload[6:8])
	req.CorrelationID = binary.BigEndian.Uint32(payload[8:12])
	return req, nil
}
