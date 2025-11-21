package processor

import (
	"bytes"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

const (
	// Error Codes
	NoErrorCode                 = 0
	UnsupportedVersionErrorCode = 35

	// API key
	APIVersionsAPIKey = 18
)

var supportedAPIKeys = []APIKey{
	{
		apiKey:     APIVersionsAPIKey,
		minVersion: 0,
		maxVersion: 4,
		tagBuffer:  0,
	},
	{
		apiKey:     DescribeTopicPartitionsAPIKey,
		minVersion: 0,
		maxVersion: 0,
		tagBuffer:  0,
	},
}

type APIVersionsProcessor struct{}

var _ KafkaAPIProcessor = (*APIVersionsProcessor)(nil)

type APIKey struct {
	apiKey     uint16
	minVersion uint16
	maxVersion uint16
	tagBuffer  uint8
}

type APIVersionsResponse struct {
	errorCode      uint16
	apiKeys        []APIKey
	throttleTimeMs uint32
	tagBuffer      uint8
}

func (avp *APIVersionsProcessor) GetRequestAPIKey() uint16 {
	return APIVersionsAPIKey
}

func (avp *APIVersionsProcessor) GenerateResponseBody(req *request.Request) ([]byte, error) {
	var errorCode uint16
	if req.RequestAPIVersion <= 4 {
		errorCode = NoErrorCode
	} else {
		errorCode = UnsupportedVersionErrorCode
	}
	resp := APIVersionsResponse{
		errorCode:      errorCode,
		apiKeys:        supportedAPIKeys,
		throttleTimeMs: 0,
		tagBuffer:      0,
	}
	buf := new(bytes.Buffer)
	utils.WriteUint16(buf, resp.errorCode)

	// TODO: Revisit unsigned variant sizes
	// size of api keys array
	utils.WriteUint8(buf, uint8(len(resp.apiKeys)+1))

	for _, apiKey := range resp.apiKeys {
		utils.WriteUint16(buf, apiKey.apiKey)
		utils.WriteUint16(buf, apiKey.minVersion)
		utils.WriteUint16(buf, apiKey.maxVersion)
		utils.WriteUint8(buf, apiKey.tagBuffer)
	}

	// Throttle in milliseconds
	utils.WriteUint32(buf, resp.throttleTimeMs)

	// Tag Buffer (empty)
	utils.WriteUint8(buf, resp.tagBuffer)

	return buf.Bytes(), nil
}
