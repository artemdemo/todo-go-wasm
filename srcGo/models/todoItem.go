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
    todoItemDoneClassname = "todo-done"
    dataTodoId = "data-todo-id"
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
                Tag: "div",
                Children: []htmlrender.ElementDef{
                    {
                        Tag: "button",
                        ClassName: services.Classnames(
                            todoItemDeleteClassname,
                            "bg-orange-500 hover:bg-orange-600 text-xs text-white",
                            "mr-1 py-1 px-2 rounded",
                        ),
                        InnerText: "Delete",
                        Attributes: []htmlrender.ElementAttr{
                            {
                                Name: dataTodoId,
                                Content: strconv.Itoa(todoItem.ID),
                            },
                        },
                    },
                    {
                        Tag: "button",
                        ClassName: services.Classnames(
                            todoItemDoneClassname,
                            "bg-green-500 hover:bg-green-600 text-xs text-white",
                            "py-1 px-2 rounded",
                        ),
                        InnerText: "Done",
                        Attributes: []htmlrender.ElementAttr{
                            {
                                Name: dataTodoId,
                                Content: strconv.Itoa(todoItem.ID),
                            },
                        },
                    },
                },
            },
        },
    }
}
