package models

import (
    "encoding/json"
    "fmt"
    "syscall/js"

   "../htmlrender"
)

type ToDoList struct {
    items []ToDoItem
}

func (todoList *ToDoList) AddTodoItem(title string, done bool) ToDoItem {
    lastTodoId := -1
    if len(todoList.items) > 0 {
        lastTodo := todoList.items[len(todoList.items) - 1]
        lastTodoId = lastTodo.ID
    }
    todoItem := ToDoItem{
        ID:    lastTodoId + 1,
        Title: title,
        Done:  done,
    }
    todoList.items = append(todoList.items, todoItem)
    return todoItem
}

func (todoList ToDoList) GetItemsJson() interface{} {
    result, err := json.Marshal(todoList.items)

    if err != nil {
        fmt.Println(err)
        return ""
    }

    return js.ValueOf(string(result))
}

func (todoList ToDoList) GetElementDef() htmlrender.ElementDef {
    var todoListEls []htmlrender.ElementDef
    for i := 0; i < len(todoList.items); i++ {
        todoListEls = append(
            todoListEls,
            todoList.items[i].GetElementDef(),
        )
    }
    return htmlrender.ElementDef{
        Tag: "div",
        Children: todoListEls,
    }
}
