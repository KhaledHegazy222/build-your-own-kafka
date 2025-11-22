package request

import (
	"io"

	"github.com/codecrafters-io/kafka-starter-go/internal/types"
	"github.com/codecrafters-io/kafka-starter-go/internal/utils"
)

type Request interface {
	Unmarshal(io.Reader) error
}

// RequestHeaders (V2)
type RequestHeaders struct {
	RequestAPIKey     types.Uint16
	RequestAPIVersion types.Uint16
	CorrelationID     types.Uint32
	ClientID          types.String
	types.TagFields
}

type BaseRequest struct {
	MessageSize types.Uint32
	Headers     RequestHeaders
	Payload     []byte
}

func (rh *RequestHeaders) Unmarshal(r io.Reader) error {
	return utils.UnmarshalAll(r, &rh.RequestAPIKey, &rh.RequestAPIVersion, &rh.CorrelationID, &rh.ClientID, &rh.TagFields)
}

func (br *BaseRequest) Unmarshal(r io.Reader) error {
	err := br.MessageSize.Unmarshal(r)
	if err != nil {
		return err
	}
	err = br.Headers.Unmarshal(r)
	if err != nil {
		return err
	}

	err = br.unmarshalRequestPayload(r)
	if err != nil {
		return err
	}

	return nil
}

func (br *BaseRequest) unmarshalRequestPayload(r io.Reader) error {
	// Read all data from the reader
	payload, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	br.Payload = payload
	return nil
}
