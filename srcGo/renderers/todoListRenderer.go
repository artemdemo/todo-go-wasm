package renderers

import (
    "strconv"
    "syscall/js"

    "../htmlrender"
    "../models"
    "../services"
)

type TodoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl js.Value
    onDeleteCb       func(todoId int64)
}

const (
    todoListClassname = "todo-list"
)

func NewTodoListRender(documentEl js.Value) *TodoListRenderer {
    todoListR := new(TodoListRenderer)
    todoListR.todoListParentEl = htmlrender.GetFirstElementByClass(documentEl, todoListClassname)
    return todoListR
}

func (this *TodoListRenderer) clickOnTodoList(_this js.Value, args []js.Value) interface{} {
    target := args[0].Get("target")
    if htmlrender.ElementHasClass(target, "todo-delete") {
        todoIdStr := target.Get("dataset").Get("todoId").String()
        todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
        this.onDeleteCb(todoId)
    }
    return ""
}

func (this *TodoListRenderer) OnDelete(cb func(todoId int64)) {
    this.onDeleteCb = cb
}

func (this *TodoListRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: services.Classnames(
            "mb-5",
            todoListClassname,
        ),
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
   htmlrender.RenderElement(
       this.todoListParentEl,
       itemEl,
   )
}
