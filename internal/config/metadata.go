package config

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

// 3 record types
// 1. Feature level record type
// 2. Topic record
// 3. Partition Record

const (
	FeatureLevelRecordType = 12
	TopicRecordType        = 2
	PartitionRecordType    = 3
)

type MetaDataTopicRecordBatch struct {
	BaseOffset           types.Uint64
	BatchLength          types.Uint32
	PartitionLeaderEpoch types.Uint32
	MagicByte            types.Uint8
	CRC                  types.Uint32
	Attributes           types.Uint16
	LastOffsetDelta      types.Uint32
	BaseTimestamp        types.Uint64
	MaxTimestamp         types.Uint64
	ProducerID           types.Int64
	ProducerEpoch        types.Int16
	BaseSequence         types.Int32
	Records              types.Array[*MetaDataTopicRecord]
}

type MetaDataTopicRecord struct {
	Length         types.VarInt
	Attributes     types.Uint8
	TimestampDelta types.VarInt
	OffsetDelta    types.VarInt
	KeyLength      types.VarInt
	Key            []byte
	ValueLength    types.VarInt
	Value          []byte
	HeadersCount   types.UVarInt
}

func (mb *MetaDataTopicRecordBatch) Unmarshal(r io.Reader) error {
	err := utils.UnmarshalAll(r, &mb.BaseOffset, &mb.BatchLength, &mb.PartitionLeaderEpoch, &mb.MagicByte, &mb.CRC, &mb.Attributes, &mb.LastOffsetDelta, &mb.BaseTimestamp, &mb.MaxTimestamp, &mb.ProducerID, &mb.ProducerEpoch, &mb.BaseSequence)
	if err != nil {
		return err
	}

	return mb.Records.Unmarshal(r, MetaDataTopicRecordFactory)
}

func (mr *MetaDataTopicRecord) Marshal(w io.Writer) error {
	err := utils.MarshalAll(w, &mr.Length, &mr.Attributes, &mr.TimestampDelta, &mr.OffsetDelta, &mr.KeyLength)
	if err != nil {
		return err
	}

	_, err = w.Write(mr.Key)
	if err != nil {
		return err
	}

	err = mr.ValueLength.Marshal(w)
	if err != nil {
		return err
	}

	_, err = w.Write(mr.Value)
	if err != nil {
		return err
	}

	err = mr.HeadersCount.Marshal(w)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MetaDataTopicRecord) Unmarshal(r io.Reader) error {
	err := utils.UnmarshalAll(r, &mr.Length, &mr.Attributes, &mr.TimestampDelta, &mr.OffsetDelta, &mr.KeyLength)
	if err != nil {
		return err
	}

	if mr.KeyLength.Value > 0 {
		mr.Key = make([]byte, mr.KeyLength.Value)
		_, err = r.Read(mr.Key)
		if err != nil {
			return err
		}
	}

	err = mr.ValueLength.Unmarshal(r)
	if err != nil {
		return err
	}

	if mr.ValueLength.Value > 0 {
		mr.Value = make([]byte, mr.ValueLength.Value)
		_, err = r.Read(mr.Value)
		if err != nil {
			return err
		}
	}

	err = mr.HeadersCount.Unmarshal(r)
	if err != nil {
		return err
	}

	return nil
}

func MetaDataTopicRecordFactory() *MetaDataTopicRecord {
	return &MetaDataTopicRecord{}
}
