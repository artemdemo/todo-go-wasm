package models

import (
    "fmt"
    "strconv"
    "syscall/js"

    "../htmlrender"
    "../services"
)

type ToDoItem struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

const (
    todoItemClassname = "todo-item"
)

// TODO this method should be in the itemRenderer
func (todoItem ToDoItem) getItemEl(baseEl js.Value) js.Value {
    return htmlrender.GetFirstElementByClass(
        baseEl,
        fmt.Sprintf("%s-%d", todoItemClassname, todoItem.ID),
    )
}

func (todoItem ToDoItem) GetItemClassname(id interface{}) string {
    if idInt, ok := id.(int); ok {
        return fmt.Sprintf("%s-%d", todoItemClassname, idInt)
    }
    return todoItemClassname
}

func (todoItem ToDoItem) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: services.Classnames(
            todoItemClassname,
            todoItem.GetItemClassname(todoItem.ID),
            "p-2 border-b-2 border-gray-200 flex justify-between",
        ),
        Children: []htmlrender.ElementDef{
            {
                Tag: "span",
                InnerText: todoItem.Title,
            },
            {
                Tag: "button",
                ClassName: services.Classnames(
                    "todo-delete",
                    "bg-gray-500 hover:bg-gray-600 text-xs text-white py-1 px-2 rounded",
                ),
                InnerText: "Delete",
                Attributes: []htmlrender.ElementAttr{
                    {
                        Name: "data-todo-id",
                        Content: strconv.Itoa(todoItem.ID),
                    },
                },
            },
        },
    }
}
