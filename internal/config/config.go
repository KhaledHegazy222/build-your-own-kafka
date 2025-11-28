package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Config struct{}

var appConfig *Config

func GetConfig() (*Config, error) {
	if appConfig == nil {
		return nil, fmt.Errorf("config is not initialized")
	}

	return appConfig, nil
}

func LoadConfig() error {
	appConfig = &Config{}

	// filePath := " ./resources/metadata/cluster-metadata.log"
	filePath := "/home/khaled/Desktop/work/codecrafters-kafka-go/resources/metadata/cluster-metadata.log"
	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("metadata log file (%s) doesn't exist", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error opening file (%s), err: %s", filePath, err)
	}

	buf := bytes.NewBuffer(data)

	for buf.Len() != 0 {
		log.Printf("Available bytes: %d", buf.Available())
		recBatch := new(MetaDataTopicRecordBatch)
		err := recBatch.Unmarshal(buf)
		if err != nil {
			return fmt.Errorf("error unmarshaling metadata log file, err: %s", err)
		}

		log.Printf("record: %+v", *recBatch)
		for _, rec := range recBatch.Records.Items {
			log.Printf("INNER REC: %+v", *rec)
		}
	}

	return nil
}
