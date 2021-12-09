package pact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	RequestKeys string `json:"requestKeys,omitempty"`
}

type PollResponse struct {
	Gas          int64       `json:"gas,omitempty"`
	ReqKey       string      `json:"reqKey,omitempty"`
	TxId         string      `json:"txId,omitempty"`
	Logs         string      `json:"logs,omitempty"`
	MetaData     interface{} `json:"metaData,omitempty"`
	Continuation interface{} `json:"continuation,omitempty"`
	Events       PactEvents  `json:"events,omitempty"`
}

type RequestKeys struct {
	RequestKeys []string `json:"requestKeys,omitempty"`
}

func Poll(requestKeys []string, apiHost string) (res *PollResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			switch er := e.(type) {
			case Error:
				err = er
			case error:
				err = er
			}
		}
	}()
	postBody, err := json.Marshal(RequestKeys{RequestKeys: requestKeys})
	EnforceNoError(err)
	req := bytes.NewBuffer(postBody)
	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/poll", apiHost),
		"application/json",
		req,
	)
	EnforceNoError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	EnforceNoError(err)
	EnforceValid(resp.StatusCode == http.StatusOK, fmt.Errorf("%v", string(body)))
	err = json.Unmarshal(body, &res)
	EnforceNoError(err)
	return
}
