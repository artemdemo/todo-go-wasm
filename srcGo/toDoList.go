package main

import (
    "./models"
)

var toDoList = models.ToDoList{
    Items: []models.ToDoItem{},
}

func addToDoItem(title string, done bool) models.ToDoItem {
    lastTodoId := 0
    if len(toDoList.Items) > 0 {
        lastTodo := toDoList.Items[len(toDoList.Items) - 1]
        lastTodoId = lastTodo.ID
    }
    id := lastTodoId + 1
    toDoItem := models.ToDoItem{
        ID: id,
        Title: title,
        Done: done,
    }
    toDoList.AddTodoItem(toDoItem)
    return toDoItem
}
