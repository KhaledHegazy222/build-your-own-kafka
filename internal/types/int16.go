package types

import (
	"encoding/binary"
	"io"
)

type Int16 struct {
	Value int16
}

func (i *Int16) Marshal(w io.Writer) error {
	var buf [2]byte
	binary.BigEndian.PutUint32(buf[:], uint32(i.Value))
	_, err := w.Write(buf[:])
	return err
}

func (i *Int16) Unmarshal(r io.Reader) error {
	var buf [2]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	i.Value = int16(binary.BigEndian.Uint16(buf[:]))
	return nil
}
