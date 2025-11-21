package response

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/internal/processor"
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

// write response header (v0)
func writeResponseHeaders(buf *bytes.Buffer, req *request.Request) error {
	// header correlation_id (32-bit)
	err := binary.Write(buf, binary.BigEndian, req.CorrelationID)

	return err
}

func writeResponseBody(buf *bytes.Buffer, req *request.Request) error {
	ap, err := processor.GetAPIProcessor(req)
	if err != nil {
		return err
	}

	body, err := ap.GenerateResponseBody(req)
	if err != nil {
		return err
	}

	err = binary.Write(buf, binary.BigEndian, body)
	return err
}

func WriteResponse(conn net.Conn, req *request.Request) error {
	buf := new(bytes.Buffer)

	// write response headers
	err := writeResponseHeaders(buf, req)
	if err != nil {
		return err
	}

	// write response body
	err = writeResponseBody(buf, req)
	if err != nil {
		return err
	}

	// write message size (32-bit)
	err = binary.Write(conn, binary.BigEndian, uint32(buf.Len()))
	if err != nil {
		return err
	}
	// write both headers & body
	_, err = conn.Write(buf.Bytes())

	return err
}
