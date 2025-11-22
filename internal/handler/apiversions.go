package handler

import (
	"github.com/codecrafters-io/kafka-starter-go/internal/constants"
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type APIVersionsProcessor struct{}

var _ APIHandler = (*APIVersionsProcessor)(nil)

func (ap *APIVersionsProcessor) GetRequestAPIKey() uint16 {
	return constants.APIVersionsAPIKey
}

func (ap *APIVersionsProcessor) Process(req *request.BaseRequest) (response.Resposne, error) {
	var errorCode uint16
	if req.Headers.RequestAPIVersion.Value <= 4 {
		errorCode = constants.NoErrorCode
	} else {
		errorCode = constants.UnsupportedVersionErrorCode
	}
	resp := response.APIVersionsResponse{
		ErrorCode:      types.Uint16{Value: errorCode},
		APIKeys:        types.CompactArray[*response.APISpec]{Items: getSupportedAPIs()},
		ThrottleTimeMs: types.Uint32{Value: 0},
		TagBuffer:      types.Uint8{Value: 0},
	}
	return &resp, nil
}

func getSupportedAPIs() []*response.APISpec {
	supportedAPIs := make([]*response.APISpec, 0)
	for _, v := range apiHandlerMapper {
		supportedAPIs = append(supportedAPIs, &v.specs)
	}
	return supportedAPIs
}
