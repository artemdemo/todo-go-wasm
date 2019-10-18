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
    todoListRenderer.AppendTodoEl(toDoItem)
}

func deleteTodo(todoId int64) {
    fmt.Println("deleteTodo", todoId)
    deletedTodo, ok := toDoList.DeleteTodoById(todoId)
    if ok {
        todoListRenderer.DeleteTodoEl(*deletedTodo)
    }
}

func toggleDone(todoId int64) {
    fmt.Println("doneTodo", todoId)
    if todoEl, _, ok := toDoList.GetTodoById(todoId); ok {
        _item := todoEl.Clone()
        _item.Done = !todoEl.Done
        todoEl.UpdateItem(_item)
    }
    fmt.Println(toDoList.GetListJson())
}

func registerCallbacks() {
    formRenderer.OnSubmit(addTodo)
    todoListRenderer.OnDelete(deleteTodo)
    todoListRenderer.OnDone(toggleDone)
}

func renderApp() {
    fmt.Println("-> renderApp()")
    app := htmlrender.NewDocumentEl().GetElementById("app")
    if appEl, ok := app.(*htmlrender.DomEl); ok {
        appEl.AppendChild(
            htmlrender.ElementDef{
                Tag: "div",
                Children: []htmlrender.ElementDef{
                    formRenderer.GetBaseElDef(),
                    todoListRenderer.GetBaseElDef(),
                    loggerRenderer.GetElementDef(),
                },
            },
        )
    } else {
        fmt.Printf("app is not of type htmlrender.DomEl, got %T instead\n", app)
        panic("app is not of type htmlrender.DomEl")
    }
}

func renderForm() {
    fmt.Println("-> renderForm()")
    formRenderer = renderers.NewFormRenderer()
    formRenderer.RenderForm(form)
}

func renderTodoList() {
    fmt.Println("-> renderTodoList()")
    todoListRenderer = renderers.NewTodoListRender()
    todoListRenderer.RenderTodoList(toDoList)
}

func renderLogger() {
    fmt.Println("-> renderLogger()")
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
