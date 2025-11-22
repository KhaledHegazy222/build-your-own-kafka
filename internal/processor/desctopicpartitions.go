package processor

import (
	"bytes"
	"log"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

const (
	DescribeTopicPartitionsAPIKey = 75

	// ErrorCodes
	UnknownTopicOrPartitionErrorCode = 3
)

type DescribeTopicPartitionsProcessor struct{}

var _ APIProcessor = (*DescribeTopicPartitionsProcessor)(nil)

func (avp *DescribeTopicPartitionsProcessor) GetRequestAPIKey() uint16 {
	return DescribeTopicPartitionsAPIKey
}

func (avp *DescribeTopicPartitionsProcessor) Process(req *request.BaseRequest) (response.Resposne, error) {
	payloadReader := bytes.NewReader(req.Payload)
	describeReq := request.DescribeTopicPartitionsRequest{
		BaseRequest: *req,
	}
	err := describeReq.Body.Unmarshal(payloadReader)
	if err != nil {
		return nil, err
	}

	return avp.perpareResposne(&describeReq)
}

func (avp *DescribeTopicPartitionsProcessor) perpareResposne(req *request.DescribeTopicPartitionsRequest) (response.Resposne, error) {
	resp := response.DescribeTopicPartitionsResponse{
		ThrottleTimeMs: types.Int32{Value: 0},
		Topics:         types.CompactArray[*response.Topic]{Items: []*response.Topic{}},
	}

	for _, topic := range req.Body.Topics.Items {
		resp.Topics.Items = append(resp.Topics.Items, &response.Topic{
			ErrorCode:  types.Int16{Value: UnknownTopicOrPartitionErrorCode},
			Name:       types.CompactNullableString{Value: &topic.Name.Value},
			TopicID:    types.UUID{Value: [16]byte{}}, // 00000000-0000-0000-0000-000000000000
			Partitions: types.CompactArray[*response.Partition]{Items: []*response.Partition{}},
			// Authorized operations
			TopicAuthorizedOperations: types.Int32{Value: 0b0000110111111000},
		})
	}
	resp.NextCursor.Value = 0xFF // -1 for null

	log.Printf("Response: %+v\n", resp.Topics.Items[0])

	return &resp, nil
}
