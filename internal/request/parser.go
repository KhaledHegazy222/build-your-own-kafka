package request

import "encoding/binary"

type RequestParser struct {
	RequestAPIKey     uint16
	RequestAPIVersion uint16
	CorrelationID     uint32
	ClientID          string
}

func NewRequestParser() RequestParser {
	return RequestParser{}
}

func (rp *RequestParser) Parse(request []byte) error {
	rp.CorrelationID = binary.BigEndian.Uint32(request[8:12])
	return nil
}
