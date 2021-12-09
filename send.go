package pact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SendResponse is the response for a /send api call
type SendResponse struct {
	RequestKeys []string `json:"requestKeys,omitempty"`
}

// Send sends a Pact command to a running Pact server and retrieves
// the transaction result
func Send(sendCmd []PrepareCommand, apiHost string) (res *SendResponse, err error) {
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
	resp, err := sendRawCmds(sendCmd, apiHost)
	EnforceNoError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	EnforceNoError(err)
	EnforceValid(resp.StatusCode == http.StatusOK, fmt.Errorf("%v", string(body)))
	err = json.Unmarshal(body, &res)
	EnforceNoError(err)
	return
}

func sendRawCmds(sendCmds []PrepareCommand, apiHost string) (*http.Response, error) {

	if apiHost == "" {
		return nil, fmt.Errorf("apiHost shouldn't be empty")
	}

	cmds := []Command{}
	for _, cmd := range sendCmds {
		switch c := cmd.(type) {
		case PrepareCont:
			cmds = append(cmds, PrepareContCmd(c))
		case PrepareExec:
			cmds = append(cmds, PrepareExecCommand(c))
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
