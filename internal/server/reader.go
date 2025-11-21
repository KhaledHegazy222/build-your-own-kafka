package server

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

func (s *Server) readRequest(conn net.Conn) (*bytes.Buffer, error) {
	// Read the 4-byte size prefix
	sizeBuf := make([]byte, 4)
	_, err := io.ReadFull(conn, sizeBuf)
	if err != nil {
		return nil, err
	}

	size := binary.BigEndian.Uint32(sizeBuf)

	// Create a single buffer with correct capacity
	buf := bytes.NewBuffer(make([]byte, 0, 4+size))

	// Write size prefix
	_, err = buf.Write(sizeBuf)
	if err != nil {
		return nil, err
	}

	// Allocate payload
	payload := make([]byte, size)

	// Read payload directly
	_, err = io.ReadFull(conn, payload)
	if err != nil {
		return nil, err
	}

	// Write payload to final buffer
	_, err = buf.Write(payload)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
