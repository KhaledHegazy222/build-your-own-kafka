package response

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type Resposne interface {
	Marshal(r io.Writer) error
}

// BaseResponse (V1)
type BaseResponse struct {
	CorrelationID types.String
	Payload       []byte
	types.TagFields
}
