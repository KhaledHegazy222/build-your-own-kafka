package request

import "github.com/codecrafters-io/kafka-starter-go/internal/types"

// APIVersionsRequest V4
type APIVersionsRequest struct {
	BaseRequest
	Body struct {
		ClientSoftwareName    types.CompactString
		ClientSoftwareVersion types.CompactString
	}
}
