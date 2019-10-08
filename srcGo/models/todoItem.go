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
    deleteBtn := Button{
        Text:      "Delete",
        BgColor:   "orange",
        ClassName: "mr-1 text-xs",
        Attributes: []htmlrender.ElementAttr{
            {
                Name: dataTodoId,
                Content: strconv.Itoa(todoItem.ID),
            },
        },
    }
    doneBtn := Button{
        Text:       "Done",
        BgColor:    "green",
        ClassName:  "text-xs",
        Attributes: []htmlrender.ElementAttr{
            {
                Name: dataTodoId,
                Content: strconv.Itoa(todoItem.ID),
            },
        },
    }
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
                    deleteBtn.GetElementDef(),
                    doneBtn.GetElementDef(),
                },
            },
        },
    }
}
