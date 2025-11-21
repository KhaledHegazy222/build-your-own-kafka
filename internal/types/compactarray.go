package types

import (
	"bufio"
	"io"
)

type CompactArray[T KafkaType] struct {
	Items []T // nil means null array
}

func (a *CompactArray[T]) Marshal(w io.Writer) error {
	// Null array
	if a.Items == nil {
		return writeUVarInt(w, 0)
	}

	// Write N+1
	if err := writeUVarInt(w, uint64(len(a.Items))+1); err != nil {
		return err
	}

	// Write items
	for _, item := range a.Items {
		if err := item.Marshal(w); err != nil {
			return err
		}
	}

	return nil
}

func (a *CompactArray[T]) Unmarshal(r io.Reader, newT func() T) error {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	lengthPlus1, err := readUVarInt(br)
	if err != nil {
		return err
	}

	// Null array
	if lengthPlus1 == 0 {
		a.Items = nil
		return nil
	}

	n := lengthPlus1 - 1

	arr := make([]T, 0, n)
	for i := 0; i < int(n); i++ {
		elem := newT()
		if err := elem.Unmarshal(r); err != nil {
			return err
		}
		arr = append(arr, elem)
	}

	a.Items = arr
	return nil
}
