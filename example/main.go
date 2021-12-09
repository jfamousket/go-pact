package main

import (
	_ "embed"
	"log"

	"github.com/jfamousket/go-pact"
	"github.com/wailsapp/wails"
)

const API_HOST = "http://localhost:9001"

var KP pact.KeyPair

func genKeyPair() (err error) {
	KP, err = pact.GenKeyPair("")
	return
}

func getTodos() *pact.LocalResponse {
	pactCode := pact.MakeExpression("todos.read-todos")
	res, err := pact.Local(pact.PrepareExec{
		KeyPairs: []pact.KeyPair{KP},
		PactCode: pactCode,
	}, API_HOST)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func add(title string, uuid string) *pact.SendResponse {
	pactCode := pact.MakeExpression("todos.new-todo", uuid, title)
	res, err := pact.Send([]pact.PrepareCommand{
		pact.PrepareExec{
			KeyPairs: []pact.KeyPair{KP},
			PactCode: pactCode,
		},
	}, API_HOST)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func toggle(id string) *pact.SendResponse {
	pactCode := pact.MakeExpression("todos.toggle-todo-status", id)
	res, err := pact.Send([]pact.PrepareCommand{
		pact.PrepareExec{
			KeyPairs: []pact.KeyPair{KP},
			PactCode: pactCode,
		},
	}, API_HOST)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func toggleAll(ids []string) *pact.SendResponse {
	var cmds []pact.PrepareCommand
	for _, id := range ids {
		pactCode := pact.MakeExpression("todos.toggle-todo-status", id)
		cmds = append(cmds, pact.PrepareExec{
			KeyPairs: []pact.KeyPair{KP},
			PactCode: pactCode,
		})
	}
	res, err := pact.Send(cmds, API_HOST)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func clearCompleted(ids []string) *pact.SendResponse {
	var cmds []pact.PrepareCommand
	for _, id := range ids {
		pactCode := pact.MakeExpression("todos.delete-todo", id)
		cmds = append(cmds, pact.PrepareExec{
			KeyPairs: []pact.KeyPair{KP},
			PactCode: pactCode,
		})
	}
	res, err := pact.Send(cmds, API_HOST)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func destroy(id string) *pact.SendResponse {
	pactCode := pact.MakeExpression("todos.delete-todo", id)
	res, err := pact.Send([]pact.PrepareCommand{
		pact.PrepareExec{
			KeyPairs: []pact.KeyPair{KP},
			PactCode: pactCode,
		},
	}, API_HOST)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func save(id string, text string) *pact.SendResponse {
	pactCode := pact.MakeExpression("todos.edit-todo", id, text)
	res, err := pact.Send([]pact.PrepareCommand{
		pact.PrepareExec{
			KeyPairs: []pact.KeyPair{KP},
			PactCode: pactCode,
		},
	}, API_HOST)
	if err != nil {
		panic(err)
	}
	return res
}

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "example",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(genKeyPair)
	app.Bind(add)
	app.Bind(save)
	app.Bind(toggle)
	app.Bind(toggleAll)
	app.Bind(clearCompleted)
	app.Bind(destroy)
	app.Bind(getTodos)
	app.Run()
}
