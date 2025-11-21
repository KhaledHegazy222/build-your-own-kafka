package types

import (
	"bufio"
	"io"
)

type CompactNullableString struct {
	Value *string
}

func (s *CompactNullableString) Marshal(w io.Writer) error {
	if s.Value == nil {
		// Null → encoded as 0
		return writeUVarInt(w, 0)
	}

	data := []byte(*s.Value)
	if err := writeUVarInt(w, uint64(len(data))+1); err != nil {
		return err
	}

	_, err := w.Write(data)
	return err
}

func (s *CompactNullableString) Unmarshal(r io.Reader) error {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	lengthPlus1, err := readUVarInt(br)
	if err != nil {
		return err
	}

	// Null?
	if lengthPlus1 == 0 {
		s.Value = nil
		return nil
	}

	n := lengthPlus1 - 1
	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}

	str := string(buf)
	s.Value = &str
	return nil
}
