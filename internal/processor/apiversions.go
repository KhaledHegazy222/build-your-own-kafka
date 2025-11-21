package processor

import (
	"bytes"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

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
	request.WriteUint16(buf, resp.errorCode)

	// TODO: Revisit unsigned variant sizes
	// size of api keys array
	request.WriteUint8(buf, uint8(len(resp.apiKeys)+1))

	for _, apiKey := range resp.apiKeys {
		request.WriteUint16(buf, apiKey.apiKey)
		request.WriteUint16(buf, apiKey.minVersion)
		request.WriteUint16(buf, apiKey.maxVersion)
		request.WriteUint8(buf, apiKey.tagBuffer)
	}

	// Throttle in milliseconds
	request.WriteUint32(buf, resp.throttleTimeMs)

	// Tag Buffer (empty)
	request.WriteUint8(buf, resp.tagBuffer)

	return buf.Bytes(), nil
}
