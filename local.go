package pact

import (
	"fmt"
	"net/http"
)

type LocalResponse struct {
	Gas          uint64      `json:"gas,omitempty"`
	Result       Result      `json:"result,omitempty"`
	ReqKey       string      `json:"reqKey,omitempty"`
	Logs         string      `json:"logs,omitempty"`
	MetaData     interface{} `json:"metaData,omitempty"`
	Continuation interface{} `json:"continuation,omitempty"`
	TxId         string      `json:"txId,omitempty"`
}

func Local(localCmd PrepareCommand, apiHost string) (res *LocalResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(Error)
		}
	}()
	resp, err := localRawCmd(localCmd, apiHost)
	EnforceNoError(err)
	defer resp.Body.Close()
	err = UnMarshalBody(resp, res)
	EnforceNoError(err)
	return
}

func localRawCmd(localCmd PrepareCommand, apiHost string) (res *http.Response, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(Error)
		}
	}()
	EnforceType(apiHost, "string", "apiHost")
	EnforceValid(apiHost != "", fmt.Errorf("No api host provided"))
	cmd := PrepareExecCommand(localCmd)
	body, err := MarshalBody(cmd)
	EnforceNoError(err)
	return http.Post(fmt.Sprintf("%s/api/v1/local", apiHost), "application/json", body)
}
