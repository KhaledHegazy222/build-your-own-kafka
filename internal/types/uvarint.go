package types

import (
	"bufio"
	"encoding/binary"
	"io"
)

type UVarInt struct {
	Value uint64
}

func NewUVarInt() *UVarInt {
	return &UVarInt{}
}

func (u *UVarInt) Marshal(w io.Writer) error {
	return writeUVarInt(w, u.Value)
}

func (u *UVarInt) Unmarshal(r io.Reader) error {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	val, err := readUVarInt(br)
	if err != nil {
		return err
	}

	u.Value = val
	return nil
}

func writeUVarInt(w io.Writer, v uint64) error {
	var buf [10]byte
	n := binary.PutUvarint(buf[:], v)
	_, err := w.Write(buf[:n])
	return err
}

func readUVarInt(r io.ByteReader) (uint64, error) {
	return binary.ReadUvarint(r)
}
