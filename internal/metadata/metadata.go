package metadata

import (
	"bytes"
	"fmt"
	"os"
)

type Metadata struct {
	Features   []FeatureLevelRecord
	Topics     []TopicRecord
	Partitions []PartitionRecord
}

var appMetadata *Metadata

func GetMetadata() (*Metadata, error) {
	if appMetadata == nil {
		return nil, fmt.Errorf("config is not initialized")
	}

	return appMetadata, nil
}

func LoadMetadata() (*Metadata, error) {
	appMetadata = &Metadata{}

	filePath := "./resources/metadata/cluster-metadata.log"
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("metadata log file (%s) doesn't exist", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file (%s), err: %s", filePath, err)
	}

	buf := bytes.NewBuffer(data)

	for buf.Len() != 0 {
		recBatch := new(MetaDataTopicRecordBatch)
		err := recBatch.Unmarshal(buf)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling metadata log file, err: %s", err)
		}

		err = ParseBatch(recBatch)
		if err != nil {
			return nil, fmt.Errorf("error processing batch, err: %s", err)
		}
	}

	return appMetadata, nil
}
