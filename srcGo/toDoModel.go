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

func model_addToDo(title string, done bool) {
	toDoList = append(toDoList, ToDo{
		ID:    len(toDoList) + 1,
		Title: title,
		Done:  done,
	})
}
