package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../models"
    "../services"
)

type itemCb func(todoId int64)

type TodoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl js.Value
    onDeleteCb       itemCb
    onDoneCb         itemCb
    // "dummyTodoItem" will be used here to retrieve locator classnames
    dummyTodoItem    models.ToDoItem
}

const (
    todoListClassname = "todo-list"
)

func NewTodoListRender() *TodoListRenderer {
    todoListR := new(TodoListRenderer)
    todoListR.todoListParentEl = htmlrender.GetFirstElementByClass(
        htmlrender.GetDocumentEl(),
        todoListClassname,
    )
    todoListR.dummyTodoItem = models.ToDoItem{}
    return todoListR
}

func (this *TodoListRenderer) clickOnTodoList(_this js.Value, args []js.Value) interface{} {
    target := args[0].Get("target")
    todoDeleteClassname := this.dummyTodoItem.GetTodoItemDeleteClassname()
    todoItemDoneClassname := this.dummyTodoItem.GetTodoItemDoneClassname();
    if htmlrender.ElementHasClass(target, todoDeleteClassname) {
        todoId := this.dummyTodoItem.GetTodoIdFromEl(target)
        this.onDeleteCb(todoId)
    } else if htmlrender.ElementHasClass(target, todoItemDoneClassname) {
        todoId := this.dummyTodoItem.GetTodoIdFromEl(target)
        this.onDoneCb(todoId)
    }
    return ""
}

func (this *TodoListRenderer) getItemEl(baseEl js.Value, todoItem models.ToDoItem) js.Value {
    return htmlrender.GetFirstElementByClass(
        baseEl,
        todoItem.GetItemIdClassname(),
    )
}

func (this *TodoListRenderer) OnDelete(cb itemCb) {
    this.onDeleteCb = cb
}

func (this *TodoListRenderer) OnDone(cb itemCb) {
    this.onDoneCb = cb
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

func (this *TodoListRenderer) RenderTodoList(todoList models.ToDoList) {
    htmlrender.ClearElementContent(this.todoListParentEl)
    htmlrender.RenderElement(
        this.todoListParentEl,
        todoList.GetElementDef(),
    )
    this.todoListParentEl.Call("addEventListener", "click", js.FuncOf(this.clickOnTodoList))
}

// AppendTodoItem is adding item to the DOM.
// And setting link to the corresponded DOM element.
func (this *TodoListRenderer) AppendTodoItem(todoItem *models.ToDoItem) {
   htmlrender.RenderElement(
       this.todoListParentEl,
       todoItem.GetElementDef(),
   )
}
