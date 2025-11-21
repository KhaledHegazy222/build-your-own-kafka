package types

import "io"

type KafkaType interface {
	Marshal(w io.Writer) error
	Unmarshal(r io.Reader) error
}
