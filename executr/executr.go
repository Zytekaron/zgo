package executr

import (
	"encoding/json"
	"github.com/zytekaron/zgo/types"
	"github.com/zytekaron/zgo/zgo"
	"net/http"
)

const BaseURL = "https://executr.zyte.dev/"

type Request struct {
	ID   string        `json:"id,omitempty"`
	Code string        `json:"code,omitempty"`
	Args []interface{} `json:"args,omitempty"`
}

type Executor struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type Executr struct {
	Token string
}

func New(token string) *Executr {
	return &Executr{
		Token: token,
	}
}

func (e *Executr) Run(code string, args ...interface{}) (data interface{}, err error) {
	var res *types.Response
	res, err = zgo.Request(http.MethodPost, BaseURL+"run", e.Token, &Request{Code: code, Args: args})
	if err != nil {
		return nil, err
	}

	return data, json.Unmarshal(res.Data, &data)
}

func (e *Executr) Call(id string, args ...interface{}) (data interface{}, err error) {
	var res *types.Response
	res, err = zgo.Request(http.MethodPost, BaseURL+id, e.Token, &Request{ID: id, Args: args})
	if err != nil {
		return nil, err
	}

	return data, json.Unmarshal(res.Data, &data)
}

func (e *Executr) Get(id string) (data *Executor, err error) {
	var res *types.Response
	res, err = zgo.Request(http.MethodGet, BaseURL+id, e.Token, &Request{ID: id})
	if err != nil {
		return nil, err
	}

	var executor *Executor
	return executor, json.Unmarshal(res.Data, &executor)
}

func (e *Executr) Delete(id string) (err error) {
	_, err = zgo.Request(http.MethodDelete, BaseURL+id, e.Token, &Request{ID: id})
	return err
}

func (e *Executr) Patch(ex *Executor) (err error) {
	_, err = zgo.Request(http.MethodPatch, BaseURL+ex.ID, e.Token, ex)
	return err
}

func (e *Executr) Post(ex *Executor) (err error) {
	_, err = zgo.Request(http.MethodPost, BaseURL, e.Token, ex)
	return err
}
