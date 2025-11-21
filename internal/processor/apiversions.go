package processor

import (
	"encoding/binary"

	"github.com/codecrafters-io/kafka-starter-go/internal/request"
)

type APIVersionsProcessor struct{}

var _ KafkaAPIProcessor = (*APIVersionsProcessor)(nil)

const (
	NoError            = 0
	UnsupportedVersion = 35
)

func (avkp *APIVersionsProcessor) GetRequestAPIKey() uint16 {
	return 18
}

func (avkp *APIVersionsProcessor) GenerateResponseBody(request *request.Request) ([]byte, error) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, UnsupportedVersion)
	return buf, nil
}
