package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

func HandleNewConnection(conn net.Conn) error {
	// NOTE: the data hear are using big endian representation
	payload := make([]byte, 1024) // 1K buffer
	for {
		size, err := conn.Read(payload)
		if err != nil {
			fmt.Println("Error Reading from connection", err.Error())
			return err
		}

		if size > 0 {
			parser := request.NewRequestParser()
			parser.Parse(payload)
			buf := bytes.NewBuffer(make([]byte, 0, 8))

			// message_size (32-bit)
			binary.Write(buf, binary.BigEndian, uint32(0))

			// header correlation_id (32-bit)
			binary.Write(buf, binary.BigEndian, parser.CorrelationID)

			conn.Write(buf.Bytes())
			conn.Close()
			return nil
		}
	}
}
