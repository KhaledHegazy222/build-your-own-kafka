package processor

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

type KafkaAPIProcessor interface {
	GetRequestAPIKey() uint16
	GenerateResponseBody(*request.Request) ([]byte, error)
}

var apiKeyProcessorMapper = make(map[uint16]KafkaAPIProcessor)

func GetAPIProcessor(req *request.Request) (KafkaAPIProcessor, error) {
	processor, exist := apiKeyProcessorMapper[req.RequestAPIKey]
	if !exist {
		return nil, fmt.Errorf("not supported api with value: %d", req.RequestAPIKey)
	}
	return processor, nil
}
