package pact

import (
	"testing"
)

func TestLocal(t *testing.T) {
	res, err := Local(PrepareExec{
		KeyPairs: []KeyPair{},
		PactCode: "(+ 1 1)",
	}, LOCAL_PACT_URL)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
