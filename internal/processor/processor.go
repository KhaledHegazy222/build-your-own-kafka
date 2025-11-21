package processor

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
)

type APIProcessor interface {
	GetRequestAPIKey() uint16
	Process(*request.BaseRequest) (response.Resposne, error)
}

var apiKeyProcessorMapper = make(map[uint16]APIProcessor)

func init() {
	apiKeyProcessorMapper[APIVersionsAPIKey] = &APIVersionsProcessor{}
	// apiKeyProcessorMapper[DescribeTopicPartitionsAPIKey] = &DescribeTopicPartitionsProcessor{}
}

func GetAPIProcessor(req *request.BaseRequest) (APIProcessor, error) {
	processor, exist := apiKeyProcessorMapper[req.Headers.RequestAPIKey.Value]
	if !exist {
		return nil, fmt.Errorf("not supported api with value: %d", req.Headers.RequestAPIKey)
	}
	return processor, nil
}
