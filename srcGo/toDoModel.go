package main

import (
    "./ui"
)

// ToDoList is list of ToDo structs
type ToDoList []ui.ToDoItem

var toDoList ToDoList

func model_addToDo(title string, done bool) ui.ToDoItem {
    lastTodoId := 0
    if len(toDoList) > 0 {
        lastTodo := toDoList[len(toDoList) - 1]
        lastTodoId = lastTodo.ID
    }
    id := lastTodoId + 1
    toDoItem := ui.ToDoItem{
        ID: id,
        Title: title,
        Done: done,
    }
    toDoList = append(toDoList, toDoItem)
    return toDoItem
}
