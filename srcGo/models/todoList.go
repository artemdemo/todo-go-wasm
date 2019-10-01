package models

import (
    "encoding/json"
    "fmt"
    "syscall/js"

   "../htmlrender"
)

type ToDoList struct {
    Items []ToDoItem
}

func (todoList *ToDoList) AddTodoItem(toDoItem ToDoItem) {
    todoList.Items = append(todoList.Items, toDoItem)
}

func (todoList ToDoList) GetItemsJson() interface{} {
    result, err := json.Marshal(todoList.Items)

    if err != nil {
        fmt.Println(err)
        return ""
    }

    return js.ValueOf(string(result))
}

func (todoList ToDoList) GetElementDef() htmlrender.ElementDef {
    var todoListEls []htmlrender.ElementDef
    for i := 0; i < len(todoList.Items); i++ {
        toDoItem := ToDoItem{
            ID: todoList.Items[i].ID,
            Title: todoList.Items[i].Title,
            Done: todoList.Items[i].Done,
        }
        todoListEls = append(
            todoListEls,
            toDoItem.GetElementDef(),
        )
    }
    return htmlrender.ElementDef{
        Tag: "div",
        Children: todoListEls,
    }
}
