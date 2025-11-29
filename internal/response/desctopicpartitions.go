package response

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

type Partition struct {
	ErrorCode              types.Int16
	PartitionIndex         types.Int32
	LeaderID               types.Int32
	LeaderEpoch            types.Int32
	ReplicaNodes           types.CompactArray[*types.Uint32]
	IsrNodes               types.CompactArray[*types.Uint32]
	EligibleLeaderReplicas types.CompactArray[*types.Uint32]
	LastKnownELR           types.CompactArray[*types.Uint32]
	OfflineReplicas        types.CompactArray[*types.Uint32]
	types.TagFields
}

type Topic struct {
	ErrorCode                 types.Int16
	Name                      types.CompactNullableString
	TopicID                   types.UUID
	IsInternal                types.Boolean
	Partitions                types.CompactArray[*Partition]
	TopicAuthorizedOperations types.Int32
	types.TagFields
}

type NextCursor struct {
	TopicName      types.CompactString
	PartitionIndex types.Int32
	types.TagFields
}

// DescribeTopicPartitionsResponse (V0)
type DescribeTopicPartitionsResponse struct {
	ThrottleTimeMs types.Int32
	Topics         types.CompactArray[*Topic]
	// TODO: Revisit this
	NextCursor types.Uint8
	// NextCursor NextCursor

	types.TagFields
}

func (dr *DescribeTopicPartitionsResponse) Marshal(w io.Writer) error {
	err := dr.ThrottleTimeMs.Marshal(w)
	if err != nil {
		return err
	}

	err = dr.Topics.Marshal(w)
	if err != nil {
		return err
	}

	err = dr.NextCursor.Marshal(w)
	if err != nil {
		return err
	}

	err = dr.TagFields.Marshal(w)
	if err != nil {
		return err
	}

	return nil
}

func (p *Partition) Marshal(w io.Writer) error {
	err := p.ErrorCode.Marshal(w)
	if err != nil {
		return err
	}

	err = p.PartitionIndex.Marshal(w)
	if err != nil {
		return err
	}

	err = p.LeaderID.Marshal(w)
	if err != nil {
		return err
	}

	err = p.LeaderEpoch.Marshal(w)
	if err != nil {
		return err
	}

	err = p.ReplicaNodes.Marshal(w)
	if err != nil {
		return err
	}

	err = p.IsrNodes.Marshal(w)
	if err != nil {
		return err
	}

	err = p.EligibleLeaderReplicas.Marshal(w)
	if err != nil {
		return err
	}

	err = p.LastKnownELR.Marshal(w)
	if err != nil {
		return err
	}

	err = p.OfflineReplicas.Marshal(w)
	if err != nil {
		return err
	}

	err = p.TagFields.Marshal(w)
	if err != nil {
		return err
	}

	return nil
}

func (p *Partition) Unmarshal(r io.Reader) error {
	err := p.ErrorCode.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.PartitionIndex.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.LeaderID.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.LeaderEpoch.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.ReplicaNodes.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = p.IsrNodes.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = p.EligibleLeaderReplicas.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = p.LastKnownELR.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = p.OfflineReplicas.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = p.TagFields.Unmarshal(r)
	if err != nil {
		return err
	}

	return nil
}

func (p *Topic) Marshal(w io.Writer) error {
	err := p.ErrorCode.Marshal(w)
	if err != nil {
		return err
	}

	err = p.Name.Marshal(w)
	if err != nil {
		return err
	}

	err = p.TopicID.Marshal(w)
	if err != nil {
		return err
	}

	err = p.IsInternal.Marshal(w)
	if err != nil {
		return err
	}

	err = p.Partitions.Marshal(w)
	if err != nil {
		return err
	}

	err = p.TopicAuthorizedOperations.Marshal(w)
	if err != nil {
		return err
	}

	err = p.TagFields.Marshal(w)
	if err != nil {
		return err
	}

	return nil
}

func (p *Topic) Unmarshal(r io.Reader) error {
	err := p.ErrorCode.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.Name.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.TopicID.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.IsInternal.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.Partitions.Unmarshal(r, paritionFactory)
	if err != nil {
		return err
	}

	err = p.TopicAuthorizedOperations.Unmarshal(r)
	if err != nil {
		return err
	}

	err = p.TagFields.Unmarshal(r)
	if err != nil {
		return err
	}

	return nil
}

func (c *NextCursor) Marshal(w io.Writer) error {
	return utils.MarshalAll(w, &c.TopicName, &c.PartitionIndex, &c.TagFields)
}

func paritionFactory() *Partition {
	return &Partition{}
}
