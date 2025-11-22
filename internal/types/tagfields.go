package types

import "io"

type TagFields Uint8

func (u *TagFields) Marshal(w io.Writer) error {
	_, err := w.Write([]byte{u.Value})
	return err
}

func (u *TagFields) Unmarshal(r io.Reader) error {
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		return err
	}
	u.Value = buf[0]
	return nil
}
