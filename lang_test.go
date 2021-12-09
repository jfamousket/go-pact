package pact_test

import (
	"testing"

	"github.com/jfamousket/go-pact"
)

func TestMakeExpression(t *testing.T) {
	pactCode := pact.MakeExpression("todos.edit-todo", int(1), "hello")
	expected := "(todos.edit-todo 1 \"hello\")"
	if pactCode != expected {
		t.Fatalf("expected=%v, got=%v", expected, pactCode)
	}
}
