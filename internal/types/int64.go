package types

import (
	"encoding/binary"
	"io"
)

type Int64 struct {
	Value int64
}

func NewInt64() *Int64 {
	return &Int64{}
}

func (i *Int64) Marshal(w io.Writer) error {
	var buf [4]byte
	binary.BigEndian.PutUint64(buf[:], uint64(i.Value))
	_, err := w.Write(buf[:])
	return err
}

func (i *Int64) Unmarshal(r io.Reader) error {
	var buf [4]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	i.Value = int64(binary.BigEndian.Uint64(buf[:]))
	return nil
}
