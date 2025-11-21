package types

import (
	"io"
)

type Array[T KafkaType] struct {
	Items []T
}

func (a *Array[T]) Marshal(w io.Writer) error {
	var lenField Int32
	lenField.Value = int32(len(a.Items))
	if err := lenField.Marshal(w); err != nil {
		return err
	}

	for _, item := range a.Items {
		if err := item.Marshal(w); err != nil {
			return err
		}
	}
	return nil
}

func (a *Array[T]) Unmarshal(r io.Reader) error {
	// Read the array length (Int32)
	var lenField Int32
	if err := lenField.Unmarshal(r); err != nil {
		return err
	}

	n := lenField.Value
	if n < 0 {
		// Kafka allows null arrays using -1
		a.Items = nil
		return nil
	}

	a.Items = make([]T, n)

	// Decode each element
	for i := int32(0); i < n; i++ {
		var item T
		if err := item.Unmarshal(r); err != nil { // ← T must implement KafkaType
			return err
		}
		a.Items[i] = item
	}

	return nil
}
