package main

// ToDo is basic model of todo list
type ToDo struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

// ToDoList is list of ToDo structs
type ToDoList []ToDo

var toDoList ToDoList

func model_addToDo(title string, done bool) int {
    id := len(toDoList) + 1;
    toDoList = append(toDoList, ToDo{
        ID:    id,
        Title: title,
        Done:  done,
    })
    return id
}
