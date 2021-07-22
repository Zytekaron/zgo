package zgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zytekaron/zgo/types"
	"net/http"
)

func Request(method, url, token string, body interface{}) (data *types.Response, err error) {
	var b []byte
	b, err = json.Marshal(body)
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	return data, json.NewDecoder(res.Body).Decode(&data)
}
