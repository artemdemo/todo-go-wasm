package models

import (
    "fmt"
    "strconv"
    "syscall/js"

    "../htmlrender"
)

type ToDoItem struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
    // Each item will have link to corresponded DOM element
    // It's not efficient!
    // But I wanted to test how it will work and to play with pointers
    el    js.Value
}

// TODO this method should be in the itemRenderer
func (toDoItem ToDoItem) getItemEl(baseEl js.Value) js.Value {
    return htmlrender.GetFirstElementByClass(
        baseEl,
        fmt.Sprintf("todo-item-%d", toDoItem.ID),
    )
}

func (toDoItem ToDoItem) Delete(baseEl js.Value) {}

func (toDoItem *ToDoItem) SetEl(el js.Value) {
    toDoItem.el = el
}

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
