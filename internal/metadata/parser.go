package metadata

import (
	"bytes"
	"fmt"
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

func ParseBatch(batch *MetaDataTopicRecordBatch) error {
	records := batch.Records.Items
	for _, record := range records {
		valueBuffer := bytes.NewBuffer(record.Value)
		baseRecord := BaseRecord{}
		err := baseRecord.Unmarshal(valueBuffer)
		if err != nil {
			return err
		}

		var p Processor

		switch baseRecord.Type.Value {
		case FeatureLevelRecordType:
			fr := &FeatureLevelRecord{BaseRecord: baseRecord}
			err = fr.Unmarshal(valueBuffer)
			p = fr

		case TopicRecordType:
			tr := &TopicRecord{BaseRecord: baseRecord}
			err = tr.Unmarshal(valueBuffer)
			p = tr

		case PartitionRecordType:
			pr := &PartitionRecord{BaseRecord: baseRecord}
			err = pr.Unmarshal(valueBuffer)
			p = pr

		default:
			err = fmt.Errorf("not supported type (%d)", baseRecord.Type.Value)
		}

		if err != nil {
			return err
		}

		err = p.Process()
		if err != nil {
			return err
		}

	}

	return nil
}

type BaseRecord struct {
	FrameVersion types.Int8
	Type         types.Int8
}

func (br *BaseRecord) Unmarshal(r io.Reader) error {
	return utils.UnmarshalAll(r, &br.FrameVersion, &br.Type)
}

type FeatureLevelRecord struct {
	BaseRecord
	Version      types.Uint8
	Name         types.CompactString
	FeatureLevel types.Uint16
	types.TagFields
}

func (fr *FeatureLevelRecord) Unmarshal(r io.Reader) error {
	return utils.UnmarshalAll(r, &fr.Version, &fr.Name, &fr.FeatureLevel, &fr.TagFields)
}

type TopicRecord struct {
	BaseRecord
	Version   types.Uint8
	Name      types.CompactString
	TopicUUID types.UUID
	types.TagFields
}

func (tr *TopicRecord) Unmarshal(r io.Reader) error {
	return utils.UnmarshalAll(r, &tr.Version, &tr.Name, &tr.TopicUUID, &tr.TagFields)
}

type PartitionRecord struct {
	BaseRecord
	Version          types.Uint8
	PartitionID      types.Uint32
	TopicUUID        types.UUID
	Replicas         types.CompactArray[*types.Uint32]
	SyncReplicas     types.CompactArray[*types.Uint32]
	RemovingReplicas types.CompactArray[*types.Uint32]
	AddingReplicas   types.CompactArray[*types.Uint32]
	Leader           types.Uint32
	LeaderEpoch      types.Uint32
	PartitionEpoch   types.Uint32
	Directories      types.CompactArray[*types.UUID]
	types.TagFields
}

func (pr *PartitionRecord) Unmarshal(r io.Reader) error {
	err := utils.UnmarshalAll(r, &pr.Version, &pr.PartitionID, &pr.TopicUUID)
	if err != nil {
		return err
	}

	err = pr.Replicas.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = pr.SyncReplicas.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = pr.RemovingReplicas.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = pr.AddingReplicas.Unmarshal(r, types.NewUint32)
	if err != nil {
		return err
	}

	err = utils.UnmarshalAll(r, &pr.Leader, &pr.LeaderEpoch, &pr.PartitionEpoch)
	if err != nil {
		return err
	}

	err = pr.Directories.Unmarshal(r, types.NewUUID)
	if err != nil {
		return err
	}

	return pr.TagFields.Unmarshal(r)
}
