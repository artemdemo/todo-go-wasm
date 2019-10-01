package main

import (
    "encoding/json"
    "fmt"
    "strconv"
    "syscall/js"

    "./htmlrender"
)

func initToDoList() {
    model_addToDo("First title", false)
    model_addToDo("Second title", true)
}

func addToDo(this js.Value, args []js.Value) interface{} {
    title := getTitleInputEl().Get("value").String()
    done := false
    id := model_addToDo(title, done)
    htmlrender.RenderElement(
        getTodoListEL(),
        htmlrender.CreateElement(
            getDocumentEl(),
            todoItemElementDef(id, title, done),
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
                {
                    Tag: "div",
                    ClassName: "mb-4",
                    Children: []htmlrender.ElementDef{
                        {
                            Tag: "input",
                            ID: "todo-title",
                            ClassName: "bg-white focus:outline-none focus:shadow-outline border border-gray-300 rounded-lg py-2 px-4 block w-full appearance-none leading-normal mb-4",
                            Attributes: []htmlrender.ElementAttr{
                                {
                                    Name: "placeholder",
                                    Content: "Title",
                                },
                            },
                        },
                        {
                            Tag: "button",
                            ID: "submit-todo",
                            ClassName: "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded",
                            InnerText: "Add ToDo",
                        },
                    },
                },
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

func todoItemElementDef(id int, title string, done bool) htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: fmt.Sprintf(
            "todo-item todo-item-%d p-2 border-b-2 border-gray-200 flex justify-between",
            id,
        ),
        Children: []htmlrender.ElementDef{
            {
                Tag: "span",
                InnerText: title,
            },
            {
                Tag: "button",
                ClassName: "todo-delete bg-gray-500 hover:bg-gray-600 text-xs text-white py-1 px-2 rounded",
                InnerText: "Delete",
                Attributes: []htmlrender.ElementAttr{
                    {
                        Name: "data-todo-id",
                        Content: strconv.Itoa(id),
                    },
                },
            },
        },
    }
}

func renderTodoList() {
    var todoListEls []htmlrender.ElementDef
    for i := 0; i < len(toDoList); i++ {
        todoListEls = append(
            todoListEls,
            todoItemElementDef(
                toDoList[i].ID,
                toDoList[i].Title,
                toDoList[i].Done,
            ),
        )
    }
    htmlrender.ClearElementContent(getTodoListEL())
    htmlrender.RenderElement(
        getTodoListEL(),
        htmlrender.CreateElement(
            getDocumentEl(),
            htmlrender.ElementDef{
                Tag: "div",
                Children: todoListEls,
            },
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
