package main

import (
    "./htmlrender"
    "./models"
    "./renderers"
    "fmt"
)

var todoList = models.TodoList{}
var form = models.Form{}

var todoListRenderer *renderers.TodoListRenderer
var formRenderer *renderers.FormRenderer
var loggerRenderer *renderers.LoggerRenderer

func initTodoList() {
    todoList.AddTodoItem("First title", false)
    todoList.AddTodoItem("Second title", true)
}

func addTodo(title string) {
    todoItem := todoList.AddTodoItem(
        title,
        false,
    )
    loggerRenderer.AppendLogMsg(fmt.Sprintf(
        "addTodo id=%d \"%s\"",
        todoItem.GetId(),
        todoItem.GetTitle(),
    ))
    todoListRenderer.AppendTodoEl(todoItem)
}

func deleteTodo(todoId int64) {
    deletedTodo, ok := todoList.DeleteTodoById(todoId)
    loggerRenderer.AppendLogMsg(fmt.Sprintf(
        "deleteTodo id=%d \"%s\"",
        deletedTodo.GetId(),
        deletedTodo.GetTitle(),
    ))
    if ok {
        todoListRenderer.DeleteTodoEl(*deletedTodo)
    }
}

func toggleDone(todoId int64) {
    if todoEl, _, ok := todoList.GetTodoById(todoId); ok {
        todoEl.SetDone(!todoEl.GetDone())
        todoListRenderer.UpdateTodo(*todoEl)
        loggerRenderer.AppendLogMsg(fmt.Sprintf(
            "doneTodo id=%d \"%s\"",
            todoEl.GetId(),
            todoEl.GetTitle(),
        ))
    }
    fmt.Println(todoList.GetListJson())
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
    loggerRenderer.AppendLogMsg("-> renderForm()")
    formRenderer = renderers.NewFormRenderer()
    formRenderer.RenderForm(form)
}

func renderTodoList() {
    loggerRenderer.AppendLogMsg("-> renderTodoList()")
    todoListRenderer = renderers.NewTodoListRender()
    todoListRenderer.RenderTodoList(todoList)
}

func renderLogger() {
    fmt.Println("-> renderLogger()")
    loggerRenderer = renderers.NewLoggerRenderer()
}

func main() {
    // Creating a channel will turn program into long-running one
    c := make(chan bool)

    initTodoList()
    renderApp()
    renderLogger()
    renderForm()
    renderTodoList()

    registerCallbacks()

    loggerRenderer.AppendLogMsg("WASM Go Initialized")

    c <- true
}
