package request

import (
	"bytes"
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
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

func parseHeaders(buf *bytes.Buffer, req *Request) error {
	var err error
	req.MessageSize, err = utils.ReadUint32(buf)
	if err != nil {
		return err
	}
	req.RequestAPIKey, err = utils.ReadUint16(buf)
	if err != nil {
		return err
	}

	req.RequestAPIVersion, err = utils.ReadUint16(buf)
	if err != nil {
		return err
	}

	req.CorrelationID, err = utils.ReadUint32(buf)
	if err != nil {
		return err
	}

	req.ClientID, err = utils.ReadString(buf, false)
	if err != nil {
		return err
	}

	return nil
}

func parseBody(buf *bytes.Buffer, req *Request) error {
	req.Body = make([]byte, buf.Available())
	n, err := buf.Read(req.Body)
	if err != nil {
		return err
	}
	// didn't read the whole message
	if len(req.Body) != n {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func Parse(buf *bytes.Buffer) (*Request, error) {
	req := new(Request)

	err := parseHeaders(buf, req)
	if err != nil {
		return nil, err
	}

	err = parseBody(buf, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
