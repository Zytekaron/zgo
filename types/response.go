package types

import (
	"encoding/json"
	"io"
)

type Response struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Code    int32           `json:"code,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (r *Response) Decode(reader io.Reader) (res *Response, err error) {
	return r, json.NewDecoder(reader).Decode(&res)
}
