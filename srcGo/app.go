package main

import (
    "encoding/json"
    "fmt"
    "syscall/js"

    "./htmlrender"
    "./models"
)

var toDoList = models.ToDoList{}

var form = models.Form{}

func initToDoList() {
    toDoList.AddTodoItem("First title", false)
    toDoList.AddTodoItem("Second title", true)
}

func addToDo(this js.Value, args []js.Value) interface{} {
    title := getTitleInputEl().Get("value").String()
    getTitleInputEl().Set("value", "")
    toDoItem := toDoList.AddTodoItem(title, false)
    htmlrender.RenderElement(
        getTodoListEL(),
        htmlrender.CreateElement(
            getDocumentEl(),
            toDoItem.GetElementDef(),
        ),
    )
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
                form.GetElementDef(),
                {
                    Tag: "div",
                    ClassName: "todo-list mb-5",
                },
                {
                    Tag: "div",
                    ClassName: "app-logger rounded bg-gray-100 p-3 text-gray-500",
                },
            },
        },
    )
    htmlrender.RenderElement(
        htmlrender.GetElementById(document, "app"),
        btnEl,
    )
}

func renderTodoList() {
    htmlrender.ClearElementContent(getTodoListEL())
    htmlrender.RenderElement(
        getTodoListEL(),
        htmlrender.CreateElement(
            getDocumentEl(),
            toDoList.GetElementDef(),
        ),
    )
}

func main() {
    // Creating a channel will turn program into long-running one
    c := make(chan bool)

    initToDoList()
    registerCallbacks()
    renderApp()
    renderTodoList()

    logToDOM("WASM Go Initialized")

    c <- true
}
