package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../models"
)

type TodoListRenderer struct {
    todoListEl js.Value
}

func (todoListRenderer TodoListRenderer) getTodoListEL(documentEl js.Value) js.Value {
    if todoListRenderer.todoListEl.Type() == js.TypeUndefined {
        todoListRenderer.todoListEl = htmlrender.GetFirstElementByClass(documentEl, "todo-list")
    }
    return todoListRenderer.todoListEl
}

func (todoListRenderer TodoListRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: "todo-list mb-5",
    }
}

func (todoListRenderer TodoListRenderer) RenderTodoList(documentEl js.Value,
                                                        todoList models.ToDoList) {
    todoListEl := todoListRenderer.getTodoListEL(documentEl)
    htmlrender.ClearElementContent(todoListEl)
    htmlrender.RenderElement(
        todoListEl,
        htmlrender.CreateElement(
            documentEl,
            todoList.GetElementDef(),
        ),
    )
}

func (todoListRenderer TodoListRenderer) AppendTodoItem(documentEl js.Value,
                                                        todoItem models.ToDoItem) {
    htmlrender.RenderElement(
        todoListRenderer.getTodoListEL(documentEl),
        htmlrender.CreateElement(
            documentEl,
            todoItem.GetElementDef(),
        ),
    )
}
