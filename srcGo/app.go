package main

import (
    "encoding/json"
    "fmt"
    "syscall/js"

    "./htmlrender"
)

func initToDoList() {
    model_addToDo("First title", false)
    model_addToDo("Second title", false)
}

func addToDo(this js.Value, args []js.Value) interface{} {
    title := args[0].String()
    done := args[1].Truthy()
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

func logToDOM(msg string) {
    htmlrender.RenderElement(
        getAppLoggerEl(),
        htmlrender.CreateElement(
            getDocumentEl(),
            htmlrender.ElementDef{
                Tag: "p",
                Children: []htmlrender.ElementDef{
                    { InnerText: msg },
                },
            },
        ),
    )
}

func renderApp() {
    document := getDocumentEl()
    btnEl := htmlrender.CreateElement(
        document,
        htmlrender.ElementDef{
            Tag: "div",
            Children: []htmlrender.ElementDef{
                {
                    Tag: "input",
                    ID: "todo-title",
                    ClassName: "bg-white focus:outline-none focus:shadow-outline border border-gray-300 rounded-lg py-2 px-4 block w-full appearance-none leading-normal mb-4",
                },
                {
                    Tag: "button",
                    ID: "submit-todo",
                    ClassName: "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded",
                    InnerText: "Add ToDo",
                },
                {
                    Tag: "div",
                    ClassName: "todo-list",
                },
                {
                    Tag: "div",
                    ClassName: "app-logger my-4",
                },
            },
        },
    )
    htmlrender.RenderElement(
        htmlrender.GetElementById(document, "app"),
        btnEl,
    )
}

func main() {
    // Creating a channel will turn program into long-running one
    c := make(chan bool)

    initToDoList()
    registerCallbacks()
    renderApp()

    logToDOM("WASM Go Initialized")

    c <- true
}
