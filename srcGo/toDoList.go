package main

import (
    "./models"
)

// ToDoList is list of ToDo structs
type ToDoList []models.ToDoItem

var toDoList ToDoList

func model_addToDo(title string, done bool) models.ToDoItem {
    lastTodoId := 0
    if len(toDoList) > 0 {
        lastTodo := toDoList[len(toDoList) - 1]
        lastTodoId = lastTodo.ID
    }
    id := lastTodoId + 1
    toDoItem := models.ToDoItem{
        ID: id,
        Title: title,
        Done: done,
    }
    toDoList = append(toDoList, toDoItem)
    return toDoItem
}
