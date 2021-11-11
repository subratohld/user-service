package handler

import (
	"encoding/json"
	"net/http"
)

type jsonResp struct {
	Msg  string      `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Errs []string    `json:"errors,omitempty"`
}

func (res *jsonResp) Message(msg string) *jsonResp {
	res.Msg = msg
	return res
}

func (res *jsonResp) Body(resp interface{}) *jsonResp {
	res.Data = resp
	return res
}

func (res *jsonResp) Errors(errs ...string) *jsonResp {
	res.Errs = errs
	return res
}

func (res *jsonResp) Write(w http.ResponseWriter, code int) {
	bytes, _ := json.Marshal(res)

	w.Header().Add("content-type", "application/json")
	w.Write(bytes)
	w.WriteHeader(code)
}
