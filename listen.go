package pact

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ListenResponse struct {
	Gas          uint64      `json:"gas"`
	Result       Result      `json:"result"`
	ReqKey       string      `json:"reqKey"`
	Logs         string      `json:"logs"`
	MetaData     interface{} `json:"metaData"`
	Continuation interface{} `json:"continuation"`
	TxId         interface{} `json:"txId"`
}

func Listen(requestKey string, apiHost string) (res *ListenResponse, err error) {
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
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/listen", apiHost))
	EnforceNoError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	EnforceNoError(err)
	EnforceValid(resp.StatusCode == http.StatusOK, fmt.Errorf("%v", string(body)))
	err = json.Unmarshal(body, &res)
	EnforceNoError(err)
	return
}
