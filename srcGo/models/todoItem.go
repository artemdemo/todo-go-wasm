package models

import (
    "fmt"
    "strconv"

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
    todoItemDeleteClassname = "todo-delete"
)

func (todoItem *ToDoItem) GetItemIdClassname() string {
    return fmt.Sprintf("%s-%d", todoItemClassname, todoItem.ID)
}

func (todoItem *ToDoItem) GetItemDeleteClassname() string {
    return todoItemDeleteClassname
}

func (todoItem *ToDoItem) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: services.Classnames(
            todoItemClassname,
            todoItem.GetItemIdClassname(),
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
                    todoItemDeleteClassname,
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
