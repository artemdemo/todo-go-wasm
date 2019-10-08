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
var loggerRenderer *renderers.LoggerRenderer

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
    todoListRenderer.AppendTodoItem(toDoItem)
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

func renderApp() {
    documentEl := htmlrender.GetDocumentEl()
    baseAppEl := htmlrender.CreateElement(
        htmlrender.ElementDef{
            Tag: "div",
            Children: []htmlrender.ElementDef{
                formRenderer.GetBaseElDef(),
                todoListRenderer.GetBaseElDef(),
                loggerRenderer.GetElementDef(),
            },
        },
    )
    htmlrender.RenderElement(
        htmlrender.GetElementById(documentEl, "app"),
        baseAppEl,
    )
}

func renderForm() {
    formRenderer = renderers.NewFormRenderer()
    formRenderer.RenderForm(form)
}

func renderTodoList() {
    todoListRenderer = renderers.NewTodoListRender()
    todoListRenderer.RenderTodoList(toDoList)
}

func renderLogger() {
    loggerRenderer = renderers.NewLoggerRenderer()
}

func main() {
    // Creating a channel will turn program into long-running one
    c := make(chan bool)

    initToDoList()
    renderApp()
    renderForm()
    renderTodoList()
    renderLogger()

    registerCallbacks()

    loggerRenderer.AppendLogMsg("WASM Go Initialized")

    c <- true
}
