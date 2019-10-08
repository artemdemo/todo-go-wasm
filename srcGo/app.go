package main

import (
    "./htmlrender"
    "./models"
    "./renderers"
    "fmt"
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

func addTodo(title string) {
    toDoItem := toDoList.AddTodoItem(
        title,
        false,
    )
    todoListRenderer.AppendTodoItem(toDoItem)
}

func deleteTodo(todoId int64) {
    fmt.Println("deleteTodo", todoId)
}

func toggleDone(todoId int64) {
    fmt.Println("doneTodo", todoId)
}

func registerCallbacks() {
    formRenderer.OnSubmit(addTodo)
    todoListRenderer.OnDelete(deleteTodo)
    todoListRenderer.OnDone(toggleDone)
}

func renderApp() {
    htmlrender.RenderElement(
        htmlrender.GetElementById(
            htmlrender.GetDocumentEl(),
            "app",
        ),
        htmlrender.ElementDef{
            Tag: "div",
            Children: []htmlrender.ElementDef{
                formRenderer.GetBaseElDef(),
                todoListRenderer.GetBaseElDef(),
                loggerRenderer.GetElementDef(),
            },
        },
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
