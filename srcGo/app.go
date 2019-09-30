package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// ToDo is basic model of todo list
type ToDo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// ToDoList is list of ToDo structs
type ToDoList []ToDo

var toDoList ToDoList

func initToDoList() {
	toDoList = []ToDo{
		ToDo{
			ID:    1,
			Title: "First title",
			Done:  false,
		},
		ToDo{
			ID:    2,
			Title: "Second title",
			Done:  false,
		},
	}
}

func add(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(
		args[0].Int() + args[1].Int(),
	)
}

func getToDoList(this js.Value, args []js.Value) interface{} {
	result, err := json.Marshal(toDoList)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return js.ValueOf(string(result))
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("getToDoList", js.FuncOf(getToDoList))
}

func printToDOM(msg string) {
	document := js.Global().Get("document")
	msgEl := document.Call("createElement", "p")
	msgEl.Set("innerText", msg)
	appLoggerEl := document.Call("getElementsByClassName", "app-logger").Index(0)
	appLoggerEl.Call("appendChild", msgEl)
}

func main() {
	// Creating a channel will turn program into long-running one
	c := make(chan bool)

	printToDOM("WASM Go Initialized")

	initToDoList()
	registerCallbacks()

	c <- true
}
