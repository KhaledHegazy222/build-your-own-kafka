package response

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type APIKey struct {
	APIKey     types.Uint16
	MinVersion types.Uint16
	MaxVersion types.Uint16
	TagBuffer  types.Uint8
}

func (ak *APIKey) Marshal(w io.Writer) error {
	err := ak.APIKey.Marshal(w)
	if err != nil {
		return err
	}

	err = ak.MinVersion.Marshal(w)
	if err != nil {
		return err
	}

	err = ak.MaxVersion.Marshal(w)
	if err != nil {
		return err
	}

	err = ak.TagBuffer.Marshal(w)
	if err != nil {
		return err
	}

	return nil
}

func (ak *APIKey) Unmarshal(r io.Reader) error {
	err := ak.APIKey.Unmarshal(r)
	if err != nil {
		return err
	}

	err = ak.MinVersion.Unmarshal(r)
	if err != nil {
		return err
	}

	err = ak.MaxVersion.Unmarshal(r)
	if err != nil {
		return err
	}

	err = ak.TagBuffer.Unmarshal(r)
	if err != nil {
		return err
	}

	return nil
}

type APIVersionsResponse struct {
	ErrorCode      types.Uint16
	APIKeys        types.CompactArray[*APIKey]
	ThrottleTimeMs types.Uint32
	TagBuffer      types.Uint8
}

func (avr *APIVersionsResponse) Marshal(w io.Writer) error {
	err := avr.ErrorCode.Marshal(w)
	if err != nil {
		return err
	}
	err = avr.APIKeys.Marshal(w)
	if err != nil {
		return err
	}
	err = avr.ThrottleTimeMs.Marshal(w)
	if err != nil {
		return err
	}
	err = avr.TagBuffer.Marshal(w)
	if err != nil {
		return err
	}
	return nil
}
