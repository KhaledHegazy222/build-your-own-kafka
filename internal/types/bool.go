package types

import (
	"io"
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Marshal(w io.Writer) error {
	var v byte = 0
	if b.Value {
		v = 1
	}
	_, err := w.Write([]byte{v})
	return err
}

func (b *Boolean) Unmarshal(r io.Reader) error {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return err
	}
	b.Value = buf[0] != 0
	return nil
}
