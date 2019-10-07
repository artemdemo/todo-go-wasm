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
    // "dummyTodoItem" will be used here to retrieve locator classnames
    dummyTodoItem    models.ToDoItem
}

const (
    todoListClassname = "todo-list"
)

func NewTodoListRender(documentEl js.Value) *TodoListRenderer {
    todoListR := new(TodoListRenderer)
    todoListR.todoListParentEl = htmlrender.GetFirstElementByClass(documentEl, todoListClassname)
    todoListR.dummyTodoItem = models.ToDoItem{}
    return todoListR
}

func (this *TodoListRenderer) clickOnTodoList(_this js.Value, args []js.Value) interface{} {
    target := args[0].Get("target")
    todoDeleteClassname := this.dummyTodoItem.GetItemDeleteClassname()
    if htmlrender.ElementHasClass(target, todoDeleteClassname) {
        todoIdStr := target.Get("dataset").Get("todoId").String()
        todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
        this.onDeleteCb(todoId)
    }
    return ""
}

func (this *TodoListRenderer) getItemEl(baseEl js.Value, todoItem models.ToDoItem) js.Value {
    return htmlrender.GetFirstElementByClass(
        baseEl,
        todoItem.GetItemIdClassname(),
    )
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
