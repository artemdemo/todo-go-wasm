package models

import (
    "encoding/json"
    "fmt"
    "syscall/js"

   "../htmlrender"
)

type ToDoList struct {
    items []TodoItem
}

// AddTodoItem is adding TodoItem to the list of items
// It will return pointer to the item.
// This way user could add link to the DOM element later.
func (todoList *ToDoList) AddTodoItem(title string, done bool) *TodoItem {
    var lastTodoId int64
    if len(todoList.items) > 0 {
        lastTodo := todoList.items[len(todoList.items) - 1]
        lastTodoId = lastTodo.ID
    }
    todoItem := TodoItem{
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

func (todoList *ToDoList) GetTodoById(todoId int64) (*TodoItem, int, bool) {
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
        return &todoList.items[indexResult], indexResult, true
    }
    return &TodoItem{}, 0, false
}

// Remove `to do` from the list (by it's ID)
func (todoList *ToDoList) DeleteTodoById(todoId int64) (*TodoItem, bool) {
    if deletedTodo, indexResult, ok := todoList.GetTodoById(todoId); ok {
        // Removing item from slice, while keeping the order
        // @link https://stackoverflow.com/a/57213476
        result := make([]TodoItem, 0)
        result = append(result, todoList.items[:indexResult]...)
        todoList.items = append(result, todoList.items[indexResult + 1:]...)
        return deletedTodo, true
    }
    return &TodoItem{}, false
}
