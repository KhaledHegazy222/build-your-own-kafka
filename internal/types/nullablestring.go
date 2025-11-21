package types

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type NullableString struct {
	Value *string // nil means "null"
}

func (s *NullableString) Marshal(w io.Writer) error {
	if s.Value == nil {
		// Write -1 as INT16
		var lenBuf [2]byte
		binary.BigEndian.PutUint16(lenBuf[:], uint16(0xFFFF))
		_, err := w.Write(lenBuf[:])
		return err
	}

	data := []byte(*s.Value)
	if len(data) > math.MaxInt16 {
		return fmt.Errorf("NULLABLE_STRING too long")
	}

	var lenBuf [2]byte
	binary.BigEndian.PutUint16(lenBuf[:], uint16(len(data)))

	if _, err := w.Write(lenBuf[:]); err != nil {
		return err
	}
	_, err := w.Write(data)
	return err
}

func (s *NullableString) Unmarshal(r io.Reader) error {
	var lenBuf [2]byte
	if _, err := io.ReadFull(r, lenBuf[:]); err != nil {
		return err
	}
	n := int16(binary.BigEndian.Uint16(lenBuf[:]))

	// Null?
	if n == -1 {
		s.Value = nil
		return nil
	}

	if n < 0 {
		return fmt.Errorf("invalid nullable string length %d", n)
	}

	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}

	str := string(buf)
	s.Value = &str
	return nil
}
