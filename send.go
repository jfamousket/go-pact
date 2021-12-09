package pact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendResponse is the response for a /send api call
type SendResponse struct {
	ReqKey string `json:"reqKey,omitempty"`
}

// Send sends a Pact command to a running Pact server and retrieves
// the transaction result
func Send(sendCmd []PrepareCommand, apiHost string) (res *SendResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(Error)
		}
	}()
	resp, err := sendRawCmds(sendCmd, apiHost)
	EnforceNoError(err)
	defer resp.Body.Close()
	err = UnMarshalBody(resp, res)
	EnforceNoError(err)
	return
}

func sendRawCmds(sendCmds []PrepareCommand, apiHost string) (*http.Response, error) {

	if apiHost == "" {
		return nil, fmt.Errorf("apiHost shouldn't be empty")
	}

	cmds := []Command{}
	for _, cmd := range sendCmds {
		if cmd.CmdType == CONT {
			cmds = append(cmds, PrepareContCmd(cmd))
		} else if cmd.CmdType == EXEC {
			cmds = append(cmds, PrepareExecCommand(cmd))
		}
	}

	body, err := json.Marshal(SendCommand{
		Cmds: cmds,
	})
	fmt.Println(string(body))
	EnforceNoError(err)

	bodyBytes := bytes.NewBuffer(body)
	return http.Post(fmt.Sprintf("%s/api/v1/send", apiHost), "application/json", bodyBytes)
}
