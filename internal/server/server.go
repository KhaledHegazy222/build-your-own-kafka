package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
)

func HandleNewConnection(conn net.Conn) error {
	defer conn.Close()

	// NOTE: the data hear are using big endian representation
	payload := make([]byte, 1024) // 1K buffer
	for {
		size, err := conn.Read(payload)
		if err != nil {
			fmt.Println("Error Reading from connection", err.Error())
			return err
		}

		if size > 0 {
			req, _ := request.Parse(payload)
			response.WriteResponse(conn, req)
			return nil
		}
	}
}
