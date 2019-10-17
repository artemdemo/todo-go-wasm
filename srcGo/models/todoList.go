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

// AddTodoItem is adding ToDoItem to the list of items
// It will return pointer to the item.
// This way user could add link to the DOM element later.
func (todoList *ToDoList) AddTodoItem(title string, done bool) *ToDoItem {
    var lastTodoId int64
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
    return &todoItem
}

func (todoList *ToDoList) GetListJson() interface{} {
    result, err := json.Marshal(todoList.items)

    if err != nil {
        fmt.Println(err)
        return ""
    }

    return js.ValueOf(string(result))
}

func (todoList *ToDoList) GetElementDef() htmlrender.ElementDef {
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

// Remove `to do` from the list (by it's ID)
// @link https://stackoverflow.com/a/55381756
func (todoList *ToDoList) DeleteTodoById(todoId int64) (ToDoItem, bool) {
    var indexResult int
    indexFound := false
    for index, item := range todoList.items {
        if item.ID == todoId {
            indexResult = index
            indexFound = true
            break
        }
    }
    if indexFound {
        deletedTodo := todoList.items[indexResult]
        copy(todoList.items[:indexResult], todoList.items[indexResult + 1:])
        todoList.items = todoList.items[:len(todoList.items) - 1]
        return deletedTodo, true
    }
    return ToDoItem{}, false
}
