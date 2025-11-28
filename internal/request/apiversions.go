package request

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

// APIVersionsRequest V4
type APIVersionsRequest struct {
	BaseRequest
	Body APIVersionsRequestBody
}

type APIVersionsRequestBody struct {
	ClientSoftwareName    types.CompactString
	ClientSoftwareVersion types.CompactString
}

func (ar *APIVersionsRequestBody) Unmarshal(r io.Reader) error {
	return utils.UnmarshalAll(r, &ar.ClientSoftwareName, &ar.ClientSoftwareVersion)
}
