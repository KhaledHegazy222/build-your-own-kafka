package processor

import (
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

const (
	// Error Codes
	NoErrorCode                 = 0
	UnsupportedVersionErrorCode = 35

	// API key
	APIVersionsAPIKey = 18
)

var supportedAPIKeys = []*response.APIKey{
	{
		APIKey:     types.Uint16{Value: APIVersionsAPIKey},
		MinVersion: types.Uint16{Value: 0},
		MaxVersion: types.Uint16{Value: 4},
		TagBuffer:  types.Uint8{Value: 0},
	},
	{
		APIKey:     types.Uint16{Value: DescribeTopicPartitionsAPIKey},
		MinVersion: types.Uint16{Value: 0},
		MaxVersion: types.Uint16{Value: 0},
		TagBuffer:  types.Uint8{Value: 0},
	},
}

type APIVersionsProcessor struct{}

var _ APIProcessor = (*APIVersionsProcessor)(nil)

func (avp *APIVersionsProcessor) GetRequestAPIKey() uint16 {
	return APIVersionsAPIKey
}

func (avp *APIVersionsProcessor) Process(req *request.BaseRequest) (response.Resposne, error) {
	var errorCode uint16
	if req.Headers.RequestAPIVersion.Value <= 4 {
		errorCode = NoErrorCode
	} else {
		errorCode = UnsupportedVersionErrorCode
	}
	resp := response.APIVersionsResponse{
		ErrorCode:      types.Uint16{Value: errorCode},
		APIKeys:        types.CompactArray[*response.APIKey]{Items: supportedAPIKeys},
		ThrottleTimeMs: types.Uint32{Value: 0},
		TagBuffer:      types.Uint8{Value: 0},
	}
	return &resp, nil
}
