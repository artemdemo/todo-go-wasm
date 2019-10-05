package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../models"
)

type todoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl js.Value
}

func NewTodoListRender() *todoListRenderer {
    todoListR := new(todoListRenderer)
    return todoListR
}

func (this *todoListRenderer) getTodoListParentEL(baseEl js.Value) js.Value {
    if this.todoListParentEl.Type() == js.TypeUndefined {
        this.todoListParentEl = htmlrender.GetFirstElementByClass(baseEl, "todo-list")
    }
    return this.todoListParentEl
}

func (this *todoListRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: "todo-list mb-5",
    }
}

func (this *todoListRenderer) RenderTodoList(documentEl js.Value,
                                             todoList models.ToDoList) {
    todoListParentEl := this.getTodoListParentEL(documentEl)
    htmlrender.ClearElementContent(todoListParentEl)
    htmlrender.RenderElement(
        todoListParentEl,
        htmlrender.CreateElement(
            documentEl,
            todoList.GetElementDef(),
        ),
    )
}

// AppendTodoItem is adding item to the DOM.
// And setting link to the corresponded DOM element.
func (this *todoListRenderer) AppendTodoItem(documentEl js.Value,
                                             todoItem *models.ToDoItem) {
   itemEl := htmlrender.CreateElement(
       documentEl,
       todoItem.GetElementDef(),
   )
   todoItem.SetEl(itemEl)
   htmlrender.RenderElement(
       this.getTodoListParentEL(documentEl),
       itemEl,
   )
}
