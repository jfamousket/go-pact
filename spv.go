package pact

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SPV(spvCmd SPVCommand, apiHost string) (res interface{}, err error) {
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
	EnforceType(spvCmd.TargetChainId, "string", "targetChainId")
	EnforceType(spvCmd.RequestKey, "string", "requestKey")
	req, err := MarshalBody(spvCmd)
	EnforceNoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/spv", apiHost), "application/json", req)
	EnforceNoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	EnforceNoError(err)
	EnforceValid(resp.StatusCode == http.StatusOK, fmt.Errorf("%v", string(body)))
	err = json.Unmarshal(body, &res)
	EnforceNoError(err)
	return
}
