package renderers

import (
    "../htmlrender"
    "../models"
    "../services"
    "fmt"
)

type itemFuncCb func(todoId int64)

type TodoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl *htmlrender.DomEl
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
    if todoListParentEl, ok := todoListParent.(*htmlrender.DomEl); ok {
        todoListR.todoListParentEl = todoListParentEl
    } else {
        fmt.Printf("todoListParent is not of type *htmlrender.DomEl, got %T instead\n", todoListParent)
        panic("todoListParent is not of type *htmlrender.DomEl")
    }
    todoListR.dummyTodoItem = models.ToDoItem{}
    return todoListR
}

func (this *TodoListRenderer) clickOnTodoList(evt *htmlrender.Event) {
    target := evt.GetTarget()
    todoDeleteClassname := this.dummyTodoItem.GetTodoItemDeleteClassname()
    todoItemDoneClassname := this.dummyTodoItem.GetTodoItemDoneClassname()
    if targetEl, ok := target.(*htmlrender.DomEl); ok {
        if targetEl.HasClass(todoDeleteClassname) {
            todoId := this.dummyTodoItem.GetTodoIdFromEl(targetEl.GetEl())
            this.onDeleteCb(todoId)
        } else if targetEl.HasClass(todoItemDoneClassname) {
            todoId := this.dummyTodoItem.GetTodoIdFromEl(targetEl.GetEl())
            this.onDoneCb(todoId)
        }
    } else {
        fmt.Printf("target is not of type *htmlrender.DomEl, got %T instead\n", target)
        panic("target is not of type *htmlrender.DomEl")
    }
}

func (this *TodoListRenderer) getItemEl(todoItem models.ToDoItem) *htmlrender.DomEl {
    el := this.todoListParentEl.GetFirstElementByClass(
        todoItem.GetItemIdClassname(),
    )
    if itemEl, ok := el.(*htmlrender.DomEl); ok {
        if itemEl.IsDefined() {
            return itemEl
        }
        fmt.Println("Can't locate DomEl, by given item:")
        fmt.Println(todoItem)
    } else {
        fmt.Printf("el is not of type *htmlrender.DomEl, got %T instead\n", el)
    }
    panic("Can't locate DomEl, by given item")
}

func (this *TodoListRenderer) OnDelete(cb itemFuncCb) {
    this.onDeleteCb = cb
}

func (this *TodoListRenderer) OnDone(cb itemFuncCb) {
    this.onDoneCb = cb
}

func (this *TodoListRenderer) DeleteTodoEl(todoItem models.ToDoItem) {
    el := this.getItemEl(todoItem)
    el.RemoveElFromDOM()
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

func (this *TodoListRenderer) AppendTodoEl(todoItem *models.ToDoItem) {
    this.todoListParentEl.AppendChild(
        todoItem.GetElementDef(),
    )
}
