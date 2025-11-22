package server

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

// write response header (V1)
func (s *Server) writeResponseHeaders(buf *bytes.Buffer, req *request.BaseRequest) error {
	err := req.Headers.CorrelationID.Marshal(buf)
	if err != nil {
		return err
	}

	tagFields := types.Uint8{Value: 0}
	return tagFields.Marshal(buf)
}

func (s *Server) writeResponse(conn net.Conn, req *request.BaseRequest, resp response.Resposne) error {
	buf := new(bytes.Buffer)

	// write response headers
	err := s.writeResponseHeaders(buf, req)
	if err != nil {
		return err
	}

	// write response body
	err = resp.Marshal(buf)
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
