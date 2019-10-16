package renderers

import (
    "../htmlrender"
    "../models"
    "../services"
)

type itemFuncCb func(todoId int64)

type TodoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl htmlrender.DomEl
    onDeleteCb       itemFuncCb
    onDoneCb         itemFuncCb
    // "dummyTodoItem" will be used here to retrieve locator classnames
    dummyTodoItem    models.ToDoItem
}

const (
    todoListClassname = "todo-list"
)

func NewTodoListRender() *TodoListRenderer {
    todoListR := new(TodoListRenderer)
    todoListParent := htmlrender.NewDocumentEl().GetFirstElementByClass(todoListClassname)
    if todoListParentEl, ok := todoListParent.(htmlrender.DomEl); ok {
        todoListR.todoListParentEl = todoListParentEl
    }
    todoListR.dummyTodoItem = models.ToDoItem{}
    return todoListR
}

func (this *TodoListRenderer) clickOnTodoList(evt htmlrender.Event) {
    target := evt.GetTarget()
    todoDeleteClassname := this.dummyTodoItem.GetTodoItemDeleteClassname()
    todoItemDoneClassname := this.dummyTodoItem.GetTodoItemDoneClassname()
    if targetEl, ok := target.(htmlrender.DomEl); ok {
        if targetEl.HasClass(todoDeleteClassname) {
            todoId := this.dummyTodoItem.GetTodoIdFromEl(targetEl.El)
            this.onDeleteCb(todoId)
        } else if targetEl.HasClass(todoItemDoneClassname) {
            todoId := this.dummyTodoItem.GetTodoIdFromEl(targetEl.El)
            this.onDoneCb(todoId)
        }
    }
}

func (this *TodoListRenderer) getItemEl(baseEl htmlrender.DomEl, todoItem models.ToDoItem) interface{} {
    el := baseEl.GetFirstElementByClass(
        todoItem.GetItemIdClassname(),
    )
    if itemEl, ok := el.(htmlrender.DomEl); ok {
        return itemEl
    }
    return nil
}

func (this *TodoListRenderer) OnDelete(cb itemFuncCb) {
    this.onDeleteCb = cb
}

func (this *TodoListRenderer) OnDone(cb itemFuncCb) {
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
    this.todoListParentEl.SetInnerHtml("")
    this.todoListParentEl.AppendChild(
        todoList.GetElementDef(),
    )
    this.todoListParentEl.AddEventListener(
        "click",
        this.clickOnTodoList,
    )
}

func (this *TodoListRenderer) AppendTodoItem(todoItem *models.ToDoItem) {
    this.todoListParentEl.AppendChild(
        todoItem.GetElementDef(),
    )
}
