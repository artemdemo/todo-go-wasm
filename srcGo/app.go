package main

import (
    "syscall/js"

    "./htmlrender"
    "./models"
    "./renderers"
)

var toDoList = models.ToDoList{}
var form = models.Form{}

var todoListRenderer = renderers.TodoListRenderer{}
var formRenderer = renderers.FormRenderer{}

func initToDoList() {
    toDoList.AddTodoItem("First title", false)
    toDoList.AddTodoItem("Second title", true)
}

func addToDo(this js.Value, args []js.Value) interface{} {
    title := getTitleInputEl().Get("value").String()
    getTitleInputEl().Set("value", "")
    toDoItem := toDoList.AddTodoItem(title, false)
    todoListRenderer.AppendTodoItem(getDocumentEl(), toDoItem)
    return true
}

func registerCallbacks() {
    js.Global().Set("addToDo", js.FuncOf(addToDo))
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
    baseAppEl := htmlrender.CreateElement(
        document,
        htmlrender.ElementDef{
            Tag: "div",
            Children: []htmlrender.ElementDef{
                formRenderer.GetBaseElDef(),
                todoListRenderer.GetBaseElDef(),
                {
                    Tag: "div",
                    ClassName: "app-logger rounded bg-gray-100 p-3 text-gray-500",
                },
            },
        },
    )
    htmlrender.RenderElement(
        htmlrender.GetElementById(document, "app"),
        baseAppEl,
    )
}

func renderForm() {
    formRenderer.RenderForm(getDocumentEl(), form)
}

func renderTodoList() {
    todoListRenderer.RenderTodoList(getDocumentEl(), toDoList)
}

func main() {
    // Creating a channel will turn program into long-running one
    c := make(chan bool)

    initToDoList()
    registerCallbacks()
    renderApp()
    renderForm()
    renderTodoList()

    logToDOM("WASM Go Initialized")

    c <- true
}
