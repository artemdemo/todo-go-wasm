package models

import (
    "../htmlrender"
    "../services"
)

type Form struct {}

const (
    addTodoButtonClassname = "submit-todo"
    todoTitleInputClassname = "todo-title"
)

func (form *Form) GetAddTodoButtonClassname() string {
    return addTodoButtonClassname
}

func (form *Form) GetTodoTitleInputClassname() string {
    return todoTitleInputClassname
}

func (form *Form) GetElementDef() htmlrender.ElementDef {
    addTodoBtn := Button{
       text:       "Add ToDo",
       bgColor:    "blue",
       className:  addTodoButtonClassname,
    }
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: "mb-4",
        Children: []htmlrender.ElementDef{
            {
                Tag: "input",
                ClassName: services.Classnames(
                    todoTitleInputClassname,
                    "bg-white focus:outline-none focus:shadow-outline border border-gray-300 rounded-lg py-2 px-4 block w-full appearance-none leading-normal mb-4",
                ),
                Attributes: []htmlrender.ElementAttr{
                    {
                        Name: "placeholder",
                        Content: "Title",
                    },
                },
            },
            addTodoBtn.GetElementDef(),
        },
    }
}
