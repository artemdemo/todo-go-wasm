package ui

import (
    "fmt"
    "strconv"

    "../htmlrender"
)

type ToDoItem struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

func (toDoItem ToDoItem) Delete() {}

func (toDoItem ToDoItem) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: fmt.Sprintf(
            "todo-item todo-item-%d p-2 border-b-2 border-gray-200 flex justify-between",
            toDoItem.ID,
        ),
        Children: []htmlrender.ElementDef{
            {
                Tag: "span",
                InnerText: toDoItem.Title,
            },
            {
                Tag: "button",
                ClassName: "todo-delete bg-gray-500 hover:bg-gray-600 text-xs text-white py-1 px-2 rounded",
                InnerText: "Delete",
                Attributes: []htmlrender.ElementAttr{
                    {
                        Name: "data-todo-id",
                        Content: strconv.Itoa(toDoItem.ID),
                    },
                },
            },
        },
    }
}
