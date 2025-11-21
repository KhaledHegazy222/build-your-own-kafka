package types

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type String struct {
	Value string
}

func (s *String) Marshal(w io.Writer) error {
	data := []byte(s.Value)
	if len(data) > math.MaxInt16 {
		return fmt.Errorf("STRING too long")
	}

	// Write length
	var lenBuf [2]byte
	binary.BigEndian.PutUint16(lenBuf[:], uint16(len(data)))
	if _, err := w.Write(lenBuf[:]); err != nil {
		return err
	}

	// Write bytes
	_, err := w.Write(data)
	return err
}

func (s *String) Unmarshal(r io.Reader) error {
	var lenBuf [2]byte
	if _, err := r.Read(lenBuf[:]); err != nil {
		return err
	}

	n := int(binary.BigEndian.Uint16(lenBuf[:]))
	if n < 0 {
		return fmt.Errorf("invalid STRING length")
	}

	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return err
	}

	s.Value = string(buf)
	return nil
}
