package pact

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func Local(localCmd PrepareExec, apiHost string) (res *LocalResponse, err error) {
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
	resp, err := localRawCmd(localCmd, apiHost)
	EnforceNoError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	EnforceNoError(err)
	EnforceValid(resp.StatusCode == http.StatusOK, fmt.Errorf("%v", string(body)))
	err = json.Unmarshal(body, &res)
	EnforceNoError(err)
	return
}

func localRawCmd(localCmd PrepareExec, apiHost string) (res *http.Response, err error) {
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
	EnforceType(apiHost, "string", "apiHost")
	EnforceValid(apiHost != "", fmt.Errorf("no api host provided"))
	cmd := PrepareExecCommand(localCmd)
	body, err := MarshalBody(cmd)
	EnforceNoError(err)
	return http.Post(fmt.Sprintf("%s/api/v1/local", apiHost), "application/json", body)
}
