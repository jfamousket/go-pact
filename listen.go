package pact

import (
	"fmt"
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
			err = e.(Error)
		}
	}()
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/listen", apiHost))
	EnforceNoError(err)
	defer resp.Body.Close()
	err = UnMarshalBody(resp, res)
	EnforceNoError(err)
	return
}
