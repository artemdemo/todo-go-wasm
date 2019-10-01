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

func (todoList *ToDoList) AddTodoItem(title string, done bool) ToDoItem {
    lastTodoId := -1
    if len(todoList.Items) > 0 {
        lastTodo := todoList.Items[len(todoList.Items) - 1]
        lastTodoId = lastTodo.ID
    }
    todoItem := ToDoItem{
        ID:    lastTodoId + 1,
        Title: title,
        Done:  done,
    }
    todoList.Items = append(todoList.Items, todoItem)
    return todoItem
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
