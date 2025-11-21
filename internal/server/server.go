package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
)

type Server struct {
	port uint16
}

func NewServer(port uint16) Server {
	return Server{
		port: port,
	}
}

func (s *Server) Start() error {
	address := fmt.Sprintf("0.0.0.0:%d", s.port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to bind to port 9092")
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			err = fmt.Errorf("error accepting connection: %s", err.Error())
			return err
		}
		// launch new go routine with every new connection
		go s.handleNewConnection(conn)
	}
}

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

func (s *Server) handleNewConnection(conn net.Conn) error {
	defer conn.Close()

	buf, err := s.readRequest(conn)
	if err != nil {
		return fmt.Errorf("reading request error: %s", err.Error())
	}

	req, err := request.Parse(buf)
	if err != nil {
		return err
	}
	fmt.Printf("REQUEST PARSED: %+v", req)

	return response.WriteResponse(conn, req)
}
