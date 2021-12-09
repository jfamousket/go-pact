package pact

import (
	"fmt"
	"net/http"
)

func SPV(spvCmd SPVCommand, apiHost string) (res interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(Error)
		}
	}()
	EnforceType(spvCmd.TargetChainId, "string", "targetChainId")
	EnforceType(spvCmd.RequestKey, "string", "requestKey")
	req, err := MarshalBody(spvCmd)
	EnforceNoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/spv", apiHost), "application/json", req)
	EnforceNoError(err)
	err = UnMarshalBody(resp, res)
	EnforceNoError(err)
	return
}
