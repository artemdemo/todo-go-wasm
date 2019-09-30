package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(
		args[0].Int() + args[1].Int(),
	)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
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

	registerCallbacks()

	c <- true
}
