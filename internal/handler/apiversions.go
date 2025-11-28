package handler

import (
	"bytes"
	"log"

	"github.com/codecrafters-io/kafka-starter-go/internal/constants"
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type APIVersionsHandler struct{}

var _ APIHandler = (*APIVersionsHandler)(nil)

func (ap *APIVersionsHandler) GetRequestAPIKey() uint16 {
	return constants.APIVersionsAPIKey
}

func (ap *APIVersionsHandler) Handle(req *request.BaseRequest) (response.Resposne, error) {
	payloadReader := bytes.NewReader(req.Payload)
	apiReq := request.APIVersionsRequest{
		BaseRequest: *req,
	}
	err := apiReq.Body.Unmarshal(payloadReader)
	if err != nil {
		return nil, err
	}

	log.Printf("APIVersions request: %+v", apiReq)
	return ap.perpareResposne(&apiReq)
}

func (ap *APIVersionsHandler) perpareResposne(req *request.APIVersionsRequest) (response.Resposne, error) {
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

	log.Printf("APIVersions resposne: %+v", resp)
	return &resp, nil
}

func getSupportedAPIs() []*response.APISpec {
	supportedAPIs := make([]*response.APISpec, 0)
	for _, v := range apiHandlerMapper {
		supportedAPIs = append(supportedAPIs, &v.specs)
	}
	return supportedAPIs
}
