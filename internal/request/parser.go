package request

import (
	"bytes"
)

func Parse(buf *bytes.Buffer) (*BaseRequest, error) {
	req := new(BaseRequest)
	err := req.Unmarshal(buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}
