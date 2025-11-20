package server

import (
	"bytes"
	"fmt"
	"net"
)

func HandleNewConnection(conn net.Conn) error {
	// NOTE: the data hear are using big endian representation
	request := make([]byte, 1024) // 1K buffer
	for {
		size, err := conn.Read(request)
		if err != nil {
			fmt.Println("Error Reading from connection", err.Error())
			return err
		}

		if size > 0 {
			buf := bytes.NewBuffer(make([]byte, 0, 8))

			// message_size (32-bit)
			buf.Write([]byte{0, 0, 0, 0})

			// header correlation_id (32-bit)
			// hard-coded value to 7
			buf.Write([]byte{0, 0, 0, 7})

			conn.Write(buf.Bytes())
			conn.Close()
			return nil
		}
	}
}
