package renderers

import (
    "fmt"
    "strings"
    "syscall/js"

    "../htmlrender"
    "../models"
)

type TodoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl js.Value
    onDeleteCb       func(todoId int)
}

func NewTodoListRender(documentEl js.Value) *TodoListRenderer {
    todoListR := new(TodoListRenderer)
    todoListR.todoListParentEl = htmlrender.GetFirstElementByClass(documentEl, "todo-list")
    return todoListR
}

func (this *TodoListRenderer) clickOnTodoList(_this js.Value, args []js.Value) interface{} {
    target := args[0].Get("target")
    className := target.Get("className").String()
    if strings.Contains(className, "todo-delete") {
        todoId := target.Get("dataset").Get("todoId").String()
        fmt.Println("Delete", todoId)
        this.onDeleteCb(10)
    }
    return ""
}

func (this *TodoListRenderer) OnDelete(cb func(todoId int)) {
    this.onDeleteCb = cb
}

func (this *TodoListRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: "todo-list mb-5",
    }
}

func (this *TodoListRenderer) RenderTodoList(documentEl js.Value,
                                             todoList models.ToDoList) {
    htmlrender.ClearElementContent(this.todoListParentEl)
    htmlrender.RenderElement(
        this.todoListParentEl,
        htmlrender.CreateElement(
            documentEl,
            todoList.GetElementDef(),
        ),
    )
    this.todoListParentEl.Call("addEventListener", "click", js.FuncOf(this.clickOnTodoList))
}

// AppendTodoItem is adding item to the DOM.
// And setting link to the corresponded DOM element.
func (this *TodoListRenderer) AppendTodoItem(documentEl js.Value,
                                             todoItem *models.ToDoItem) {
   itemEl := htmlrender.CreateElement(
       documentEl,
       todoItem.GetElementDef(),
   )
   todoItem.SetEl(itemEl)
   htmlrender.RenderElement(
       this.todoListParentEl,
       itemEl,
   )
}
