package types

import (
	"encoding/binary"
	"io"
)

type Uint32 struct {
	Value uint32
}

func (u *Uint32) Marshal(w io.Writer) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], u.Value)
	_, err := w.Write(buf[:])
	return err
}

func (u *Uint32) Unmarshal(r io.Reader) error {
	var buf [4]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	u.Value = binary.BigEndian.Uint32(buf[:])
	return nil
}
