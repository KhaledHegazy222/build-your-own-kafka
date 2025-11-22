package types

import (
	"io"
)

type Uint8 struct {
	Value uint8
}

func NewUint8() *Uint8 {
	return &Uint8{}
}

func (u *Uint8) Marshal(w io.Writer) error {
	_, err := w.Write([]byte{u.Value})
	return err
}

func (u *Uint8) Unmarshal(r io.Reader) error {
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	u.Value = buf[0]
	return nil
}
