package response

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

type APISpec struct {
	APIKey     types.Uint16
	MinVersion types.Uint16
	MaxVersion types.Uint16
	TagBuffer  types.Uint8
}

type APIVersionsResponse struct {
	ErrorCode      types.Uint16
	APIKeys        types.CompactArray[*APISpec]
	ThrottleTimeMs types.Uint32
	TagBuffer      types.Uint8
}

func (ar *APIVersionsResponse) Marshal(w io.Writer) error {
	err := ar.ErrorCode.Marshal(w)
	if err != nil {
		return err
	}
	err = ar.APIKeys.Marshal(w)
	if err != nil {
		return err
	}
	err = ar.ThrottleTimeMs.Marshal(w)
	if err != nil {
		return err
	}
	err = ar.TagBuffer.Marshal(w)
	if err != nil {
		return err
	}
	return nil
}

func (ak *APISpec) Marshal(w io.Writer) error {
	return utils.MarshalAll(w, &ak.APIKey, &ak.MinVersion, &ak.MaxVersion, &ak.TagBuffer)
}

func (ak *APISpec) Unmarshal(r io.Reader) error {
	return utils.UnmarshalAll(r, &ak.APIKey, &ak.MinVersion, &ak.MaxVersion, &ak.TagBuffer)
}
