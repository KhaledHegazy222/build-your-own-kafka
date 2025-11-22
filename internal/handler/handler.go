package handler

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/internal/constants"
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type APIHandler interface {
	GetRequestAPIKey() uint16
	Handle(*request.BaseRequest) (response.Resposne, error)
}

type apiEntry struct {
	handler APIHandler
	specs   response.APISpec
}

var apiHandlerMapper = make(map[uint16]apiEntry)

func init() {
	apiHandlerMapper[constants.APIVersionsAPIKey] = apiEntry{
		handler: &APIVersionsHandler{},
		specs: response.APISpec{
			APIKey:     types.Uint16{Value: constants.APIVersionsAPIKey},
			MinVersion: types.Uint16{Value: 0},
			MaxVersion: types.Uint16{Value: 4},
			TagBuffer:  types.Uint8{Value: 0},
		},
	}

	apiHandlerMapper[constants.DescribeTopicPartitionsAPIKey] = apiEntry{
		handler: &DescribeTopicPartitionsHandler{},
		specs: response.APISpec{
			APIKey:     types.Uint16{Value: constants.DescribeTopicPartitionsAPIKey},
			MinVersion: types.Uint16{Value: 0},
			MaxVersion: types.Uint16{Value: 0},
			TagBuffer:  types.Uint8{Value: 0},
		},
	}
}

func GetAPIHandler(req *request.BaseRequest) (APIHandler, error) {
	entry, exist := apiHandlerMapper[req.Headers.RequestAPIKey.Value]
	if !exist {
		return nil, fmt.Errorf("not supported api with value: %d", req.Headers.RequestAPIKey)
	}
	return entry.handler, nil
}
