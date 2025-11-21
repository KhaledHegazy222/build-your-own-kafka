package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/internal/processor"
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
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

func (s *Server) handleNewConnection(conn net.Conn) error {
	defer conn.Close()

	for {
		buf, err := s.readRequest(conn)
		if err != nil {
			return fmt.Errorf("reading request error: %s", err.Error())
		}

		req, err := request.Parse(buf)
		if err != nil {
			return fmt.Errorf("parsing request error: %s", err.Error())
		}

		proc, err := processor.GetAPIProcessor(req)
		if err != nil {
			return fmt.Errorf("processor lookup error: %s", err.Error())
		}
		resp, err := proc.Process(req)
		if err != nil {
			return fmt.Errorf("processing request error: %s", err.Error())
		}

		err = s.writeResponse(conn, req, resp)
		if err != nil {
			return fmt.Errorf("writing response error: %s", err.Error())
		}

	}
}
