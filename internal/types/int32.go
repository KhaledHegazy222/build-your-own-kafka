package types

import (
	"encoding/binary"
	"io"
)

type Int32 struct {
	Value int32
}

func (i *Int32) Marshal(w io.Writer) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(i.Value))
	_, err := w.Write(buf[:])
	return err
}

func (i *Int32) Unmarshal(r io.Reader) error {
	var buf [4]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	i.Value = int32(binary.BigEndian.Uint32(buf[:]))
	return nil
}
