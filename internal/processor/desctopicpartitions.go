package processor

import (
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
)

const (
	DescribeTopicPartitionsAPIKey = 75
)

type DescribeTopicPartitionsProcessor struct{}

var _ APIProcessor = (*DescribeTopicPartitionsProcessor)(nil)

type DescribeTopicPartitionsResponse struct{}

func (avp *DescribeTopicPartitionsProcessor) GetRequestAPIKey() uint16 {
	return DescribeTopicPartitionsAPIKey
}

func (avp *DescribeTopicPartitionsProcessor) Process(req *request.BaseRequest) (response.Resposne, error) {
	return nil, nil
}
