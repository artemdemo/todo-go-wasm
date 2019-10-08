package main

import (
    "fmt"
    "syscall/js"

    "./htmlrender"
    "./models"
    "./renderers"
)

var toDoList = models.ToDoList{}
var form = models.Form{}

var todoListRenderer *renderers.TodoListRenderer
var formRenderer *renderers.FormRenderer

func initToDoList() {
    toDoList.AddTodoItem("First title", false)
    toDoList.AddTodoItem("Second title", true)
}

func addTodo(this js.Value, args []js.Value) interface{} {
    toDoItem := toDoList.AddTodoItem(
        formRenderer.GetTitle(),
        false,
    )
    formRenderer.ClearTitleInput()
    todoListRenderer.AppendTodoItem(getDocumentEl(), toDoItem)
    return true
}

func deleteTodo(todoId int64) {
    fmt.Println("deleteTodo", todoId)
}

func toggleDone(todoId int64) {
    fmt.Println("doneTodo", todoId)
}

func registerCallbacks() {
    formRenderer.OnSubmitCb(addTodo)
    todoListRenderer.OnDelete(deleteTodo)
    todoListRenderer.OnDone(toggleDone)
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
    formRenderer = renderers.NewFormRenderer(getDocumentEl())
    formRenderer.RenderForm(getDocumentEl(), form)
}

func renderTodoList() {
    todoListRenderer = renderers.NewTodoListRender(getDocumentEl())
    todoListRenderer.RenderTodoList(getDocumentEl(), toDoList)
}

func main() {
    // Creating a channel will turn program into long-running one
    c := make(chan bool)

    initToDoList()
    renderApp()
    renderForm()
    renderTodoList()

    registerCallbacks()

    logToDOM("WASM Go Initialized")

    c <- true
}
