package types

import (
	"encoding/binary"
	"io"
)

type Uint64 struct {
	Value uint64
}

func NewUint64() *Uint64 {
	return &Uint64{}
}

func (u *Uint64) Marshal(w io.Writer) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], u.Value)
	_, err := w.Write(buf[:])
	return err
}

func (u *Uint64) Unmarshal(r io.Reader) error {
	var buf [8]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	u.Value = binary.BigEndian.Uint64(buf[:])
	return nil
}
