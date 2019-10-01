package models

import (
   "../htmlrender"
)

type ToDoList struct {
    Items []ToDoItem
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
