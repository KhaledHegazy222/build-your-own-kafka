package types

import (
	"encoding/binary"
	"io"
)

func writeUVarInt(w io.Writer, v uint64) error {
	var buf [10]byte
	n := binary.PutUvarint(buf[:], v)
	_, err := w.Write(buf[:n])
	return err
}

func readUVarInt(r io.ByteReader) (uint64, error) {
	return binary.ReadUvarint(r)
}
