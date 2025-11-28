package types

import "io"

type UUID struct {
	Value [16]byte
}

func NewUUID() *UUID {
	return &UUID{}
}

func (u *UUID) Marshal(w io.Writer) error {
	// Must write all 16 bytes
	_, err := w.Write(u.Value[:])
	return err
}

func (u *UUID) Unmarshal(r io.Reader) error {
	_, err := io.ReadFull(r, u.Value[:])
	return err
}
