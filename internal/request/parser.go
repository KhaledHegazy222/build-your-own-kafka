package request

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Request struct {
	MessageSize uint32
	// Headers
	RequestAPIKey     uint16
	RequestAPIVersion uint16
	CorrelationID     uint32
	ClientID          string
	// Body
	Body []byte
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

	req.ClientID, err = ReadString(buf, false)
	if err != nil {
		return nil, err
	}

	req.Body = make([]byte, buf.Available())
	n, err := buf.Read(req.Body)
	if err != nil {
		return nil, err
	}
	// didn't read the whole message
	if len(req.Body) != n {
		return nil, io.ErrUnexpectedEOF
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

func ReadString(buf *bytes.Buffer, isCompact bool) (string, error) {
	strSize, err := ReadUint16(buf)
	if err != nil {
		return "", err
	}

	if isCompact {
		strSize -= 1
	}

	str, err := ReadStringWithSize(buf, int(strSize))
	if int16(strSize) == -1 {
		return "", err
	}

	return str, nil
}

func ReadNullableString(buf *bytes.Buffer, isCompact bool) (string, bool, error) {
	strSize, err := ReadUint16(buf)
	if err != nil {
		return "", true, err
	}

	if isCompact {
		strSize -= 1
	}

	// string is null
	if int16(strSize) == -1 {
		return "", true, nil
	}

	str, err := ReadStringWithSize(buf, int(strSize))
	if int16(strSize) == -1 {
		return "", true, err
	}

	return str, false, nil
}

func ReadStringWithSize(buf *bytes.Buffer, size int) (string, error) {
	b := buf.Next(size)
	if len(b) < 4 {
		return "", io.ErrUnexpectedEOF
	}
	return string(b), nil
}

// WriteUint8 writes a single byte to the buffer
func WriteUint8(buf *bytes.Buffer, value uint8) error {
	return buf.WriteByte(value)
}

// WriteUint16 writes a uint16 in big-endian order
func WriteUint16(buf *bytes.Buffer, value uint16) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, value)
	_, err := buf.Write(tmp)
	return err
}

// WriteUint32 writes a uint32 in big-endian order
func WriteUint32(buf *bytes.Buffer, value uint32) error {
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, value)
	_, err := buf.Write(tmp)
	return err
}
