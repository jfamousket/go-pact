package pact

import "testing"

func TestSend(t *testing.T) {
	res, err := Send([]PrepareCommand{
		PrepareExec{
			KeyPairs: []KeyPair{},
			PactCode: "(+ 1 1)",
		},
	}, LOCAL_PACT_URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.RequestKeys) == 0 {
		t.Fatalf("expected RequestKeys, got %v", res.RequestKeys)
	}
}
