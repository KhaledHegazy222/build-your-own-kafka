package processor

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

type KafkaAPIProcessor interface {
	GetRequestAPIKey() uint16
	GenerateResponseBody(*request.Request) ([]byte, error)
}

func GetAPIProcessor(req *request.Request) (KafkaAPIProcessor, error) {
	avP := APIVersionsProcessor{}
	if req.RequestAPIKey == avP.GetRequestAPIKey() {
		return &avP, nil
	}
	return nil, fmt.Errorf("not supported api with value: %d", req.RequestAPIKey)
}
