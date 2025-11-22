package utils

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

func MarshalAll(w io.Writer, items ...types.KafkaType) error {
	for _, item := range items {
		err := item.Marshal(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnmarshalAll(r io.Reader, items ...types.KafkaType) error {
	for _, item := range items {
		err := item.Unmarshal(r)
		if err != nil {
			return err
		}
	}
	return nil
}
