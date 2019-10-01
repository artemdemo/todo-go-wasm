package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func initToDoList() {
	model_addToDo("First title", false)
	model_addToDo("Second title", false)
}

func addToDo(this js.Value, args []js.Value) interface{} {
	var title = args[0].String()
	var done = args[1].Truthy()
	model_addToDo(title, done)
	return true
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
	js.Global().Set("addToDo", js.FuncOf(addToDo))
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
