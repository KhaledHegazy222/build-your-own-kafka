package types

import (
	"io"
)

type Int8 struct {
	Value uint8
}

func NewInt8() *Int8 {
	return &Int8{}
}

func (u *Int8) Marshal(w io.Writer) error {
	_, err := w.Write([]byte{u.Value})
	return err
}

func (u *Int8) Unmarshal(r io.Reader) error {
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	u.Value = buf[0]
	return nil
}
