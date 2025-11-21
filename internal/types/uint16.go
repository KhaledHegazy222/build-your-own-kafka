package types

import (
	"encoding/binary"
	"io"
)

type Uint16 struct {
	Value uint16
}

func (u *Uint16) Marshal(w io.Writer) error {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], u.Value)
	_, err := w.Write(buf[:])
	return err
}

func (u *Uint16) Unmarshal(r io.Reader) error {
	var buf [2]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	u.Value = binary.BigEndian.Uint16(buf[:])
	return nil
}
