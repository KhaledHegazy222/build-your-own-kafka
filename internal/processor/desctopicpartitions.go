package processor

import (
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

const (
	DescribeTopicPartitionsAPIKey = 75
)

type DescribeTopicPartitionsProcessor struct{}

var _ KafkaAPIProcessor = (*DescribeTopicPartitionsProcessor)(nil)

type DescribeTopicPartitionsResponse struct{}

func (avp *DescribeTopicPartitionsProcessor) GetRequestAPIKey() uint16 {
	return DescribeTopicPartitionsAPIKey
}

func (avp *DescribeTopicPartitionsProcessor) GenerateResponseBody(req *request.Request) ([]byte, error) {
	return nil, nil
}
