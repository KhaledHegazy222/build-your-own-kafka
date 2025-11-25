package types

import (
	"bufio"
	"io"
)

type VarInt struct {
	Value int32
}

func NewVarInt() *VarInt {
	return &VarInt{}
}

func zigZagEncode32(v int32) uint64 {
	return uint64((v << 1) ^ (v >> 31))
}

func zigZagDecode32(v uint64) int32 {
	return int32((v >> 1) ^ uint64((int32(v&1)<<31)>>31))
}

func (v *VarInt) Marshal(w io.Writer) error {
	encoded := zigZagEncode32(v.Value)
	return writeUVarInt(w, encoded)
}

func (v *VarInt) Unmarshal(r io.Reader) error {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	n, err := readUVarInt(br)
	if err != nil {
		return err
	}

	v.Value = zigZagDecode32(n)
	return nil
}
