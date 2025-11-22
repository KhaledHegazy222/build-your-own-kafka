package request

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
)

type Topic struct {
	Name types.CompactString
	types.TagFields
}

type Cursor struct {
	TopicName      types.CompactString
	PartitionIndex types.Int32
	TagBuffer      types.Uint8
}

// DescribeTopicPartitionsRequest (V0)
type DescribeTopicPartitionsRequest struct {
	BaseRequest
	Body DescribeTopicPartitionsRequestBody
}
type DescribeTopicPartitionsRequestBody struct {
	Topics                types.CompactArray[*Topic]
	ResponsePartitionTime types.Int32
	Cursor                types.Uint8
	// Cursor               Cursor

	types.TagFields
}

func (d *DescribeTopicPartitionsRequestBody) Unmarshal(r io.Reader) error {
	err := d.Topics.Unmarshal(r, topicFactory)
	if err != nil {
		return err
	}

	err = d.ResponsePartitionTime.Unmarshal(r)
	if err != nil {
		return err
	}

	err = d.Cursor.Unmarshal(r)
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Marshal(w io.Writer) error {
	err := t.Name.Marshal(w)
	if err != nil {
		return err
	}

	err = t.TagFields.Marshal(w)
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Unmarshal(r io.Reader) error {
	err := t.Name.Unmarshal(r)
	if err != nil {
		return err
	}

	err = t.TagFields.Unmarshal(r)
	if err != nil {
		return err
	}

	return nil
}

func (t *Cursor) Unmarshal(r io.Reader) error {
	return nil
}

func topicFactory() *Topic {
	return &Topic{}
}
