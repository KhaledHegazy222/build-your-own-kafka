package handler

import (
	"bytes"
	"log"

	"github.com/codecrafters-io/kafka-starter-go/internal/constants"
	"github.com/codecrafters-io/kafka-starter-go/internal/manager"
	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type DescribeTopicPartitionsHandler struct{}

var _ APIHandler = (*DescribeTopicPartitionsHandler)(nil)

func (dp *DescribeTopicPartitionsHandler) GetRequestAPIKey() uint16 {
	return constants.DescribeTopicPartitionsAPIKey
}

func (dp *DescribeTopicPartitionsHandler) Handle(req *request.BaseRequest) (response.Resposne, error) {
	payloadReader := bytes.NewReader(req.Payload)
	describeReq := request.DescribeTopicPartitionsRequest{
		BaseRequest: *req,
	}
	err := describeReq.Body.Unmarshal(payloadReader)
	if err != nil {
		return nil, err
	}

	log.Printf("DescribeTopicPartions request: %+v", describeReq)

	return dp.perpareResposne(&describeReq)
}

func (dp *DescribeTopicPartitionsHandler) perpareResposne(req *request.DescribeTopicPartitionsRequest) (response.Resposne, error) {
	m, err := manager.GetManager()
	if err != nil {
		return nil, err
	}

	resp := response.DescribeTopicPartitionsResponse{
		ThrottleTimeMs: types.Int32{Value: 0},
		Topics:         types.CompactArray[*response.Topic]{Items: []*response.Topic{}},
	}

	for _, topic := range req.Body.Topics.Items {
		foundTopic, err := m.GetTopicByName(topic.Name.Value)
		if err != nil {
			resp.Topics.Items = append(resp.Topics.Items, &response.Topic{
				ErrorCode:  types.Int16{Value: constants.UnknownTopicOrPartitionErrorCode},
				Name:       types.CompactNullableString{Value: &topic.Name.Value},
				TopicID:    types.UUID{Value: [16]byte{}}, // 00000000-0000-0000-0000-000000000000
				Partitions: types.CompactArray[*response.Partition]{Items: []*response.Partition{}},
				// Authorized operations
				TopicAuthorizedOperations: types.Int32{Value: 0b0000110111111000},
			})
			continue
		}
		partitionResp := make([]*response.Partition, len(foundTopic.Partitions))
		for idx := range partitionResp {
			foundPartition := foundTopic.Partitions[idx]

			partitionResp[idx] = &response.Partition{
				ErrorCode:              types.Int16{Value: constants.NoErrorCode},
				PartitionIndex:         types.Uint32{Value: foundPartition.PartitionID},
				LeaderID:               types.Uint32{Value: foundPartition.Leader},
				LeaderEpoch:            types.Uint32{Value: foundPartition.LeaderEpoch},
				ReplicaNodes:           types.CompactArray[*types.Uint32]{Items: []*types.Uint32{}},
				IsrNodes:               types.CompactArray[*types.Uint32]{Items: []*types.Uint32{}},
				EligibleLeaderReplicas: types.CompactArray[*types.Uint32]{Items: []*types.Uint32{}},
				LastKnownELR:           types.CompactArray[*types.Uint32]{Items: []*types.Uint32{}},
				OfflineReplicas:        types.CompactArray[*types.Uint32]{Items: []*types.Uint32{}},
			}
			for _, replica := range foundPartition.Replicas {
				partitionResp[idx].ReplicaNodes.Items = append(partitionResp[idx].ReplicaNodes.Items, &types.Uint32{Value: replica})
			}
			for _, replica := range foundPartition.InSyncReplicas {
				partitionResp[idx].IsrNodes.Items = append(partitionResp[idx].IsrNodes.Items, &types.Uint32{Value: replica})
			}

		}

		resp.Topics.Items = append(resp.Topics.Items, &response.Topic{
			ErrorCode:  types.Int16{Value: constants.NoErrorCode},
			Name:       types.CompactNullableString{Value: &topic.Name.Value},
			TopicID:    types.UUID{Value: types.RawUUID(foundTopic.UUID)}, // 00000000-0000-0000-0000-000000000000
			Partitions: types.CompactArray[*response.Partition]{Items: partitionResp},
			// Authorized operations
			TopicAuthorizedOperations: types.Int32{Value: 0b0000110111111000},
		})

	}
	resp.NextCursor.Value = 0xFF // -1 for null

	log.Printf("DescribeTopicPartions resposne: %+v", resp)
	return &resp, nil
}
