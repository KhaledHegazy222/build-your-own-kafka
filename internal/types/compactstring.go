package types

import (
	"bufio"
	"io"
)

type CompactString struct {
	Value string
}

func (s *CompactString) Marshal(w io.Writer) error {
	data := []byte(s.Value)
	n := uint64(len(data))

	if err := writeUVarInt(w, n+1); err != nil {
		return err
	}
	_, err := w.Write(data)
	return err
}

func (s *CompactString) Unmarshal(r io.Reader) error {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	lengthPlus1, err := readUVarInt(br)
	if err != nil {
		return err
	}
	if lengthPlus1 == 0 {
		s.Value = ""
		return nil
	}
	n := lengthPlus1 - 1
	buf := make([]byte, n)
	_, err = io.ReadFull(r, buf)
	s.Value = string(buf)
	return err
}
