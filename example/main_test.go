package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenKeyPair(t *testing.T) {
	err := genKeyPair()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTodos(t *testing.T) {
	genKeyPair()
	todos := getTodos()
	t.Log(todos)
}

func TestAdd(t *testing.T) {
	uid := uuid.New()
	res := add("test", uid.String())
	if len(res.RequestKeys) == 0 {
		t.Fatal("no transaction created")
	}
}
