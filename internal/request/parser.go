package request

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Request struct {
	MessageSize       uint32
	RequestAPIKey     uint16
	RequestAPIVersion uint16
	CorrelationID     uint32
	ClientID          string
}

func Parse(buf *bytes.Buffer) (*Request, error) {
	req := new(Request)
	var err error

	req.MessageSize, err = ReadUint32(buf)
	if err != nil {
		return nil, err
	}
	req.RequestAPIKey, err = ReadUint16(buf)
	if err != nil {
		return nil, err
	}

	req.RequestAPIVersion, err = ReadUint16(buf)
	if err != nil {
		return nil, err
	}

	req.CorrelationID, err = ReadUint32(buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func ReadUint8(buf *bytes.Buffer) (uint8, error) {
	b := buf.Next(1) // consumes 2 bytes
	if len(b) < 1 {
		return 0, io.ErrUnexpectedEOF
	}
	return b[0], nil
}

func ReadUint16(buf *bytes.Buffer) (uint16, error) {
	b := buf.Next(2) // consumes 2 bytes
	if len(b) < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint16(b), nil
}

func ReadUint32(buf *bytes.Buffer) (uint32, error) {
	b := buf.Next(4) // consumes 4 bytes
	if len(b) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint32(b), nil
}
