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
    todoItemDeleteClassname = "todo-delete"
    todoItemDoneClassname = "todo-done"
    dataTodoId = "data-todo-id"
)

func (todoItem *ToDoItem) GetItemIdClassname() string {
    return fmt.Sprintf("%s-%d", todoItemClassname, todoItem.ID)
}

func (todoItem *ToDoItem) GetTodoItemDeleteClassname() string {
    return todoItemDeleteClassname
}

func (todoItem *ToDoItem) GetTodoItemDoneClassname() string {
    return todoItemDoneClassname
}

func (todoItem * ToDoItem) GetTodoIdFromEl(el js.Value) int64 {
    todoIdStr := el.Get("dataset").Get("todoId").String()
    todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
    return todoId
}

func (todoItem *ToDoItem) GetElementDef() htmlrender.ElementDef {
    deleteBtn := Button{
        Text:      "Delete",
        BgColor:   "orange",
        Size:       ButtonSizes.XS,
        ClassName: services.Classnames(
            "mr-1",
            todoItemDeleteClassname,
        ),
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
        Size:       ButtonSizes.XS,
        ClassName:  todoItemDoneClassname,
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
