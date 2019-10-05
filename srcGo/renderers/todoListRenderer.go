package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../models"
)

type TodoListRenderer struct {
    // "todoListParentEl" is parent element where to-do list itself will be rendered
    todoListParentEl js.Value
}

func (todoListRenderer *TodoListRenderer) getTodoListParentEL(baseEl js.Value) js.Value {
    if todoListRenderer.todoListParentEl.Type() == js.TypeUndefined {
        todoListRenderer.todoListParentEl = htmlrender.GetFirstElementByClass(baseEl, "todo-list")
    }
    return todoListRenderer.todoListParentEl
}

func (todoListRenderer TodoListRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: "todo-list mb-5",
    }
}

func (todoListRenderer *TodoListRenderer) RenderTodoList(documentEl js.Value,
                                                         todoList models.ToDoList) {
    todoListParentEl := todoListRenderer.getTodoListParentEL(documentEl)
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
func (todoListRenderer *TodoListRenderer) AppendTodoItem(documentEl js.Value,
                                                         todoItem *models.ToDoItem) {
   itemEl := htmlrender.CreateElement(
       documentEl,
       todoItem.GetElementDef(),
   )
   todoItem.SetEl(itemEl)
   htmlrender.RenderElement(
       todoListRenderer.getTodoListParentEL(documentEl),
       itemEl,
   )
}
