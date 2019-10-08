package models

import (
    "../htmlrender"
)

type Form struct {}

func (form *Form) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: "mb-4",
        Children: []htmlrender.ElementDef{
            {
                Tag: "input",
                ID: "todo-title",
                ClassName: "bg-white focus:outline-none focus:shadow-outline border border-gray-300 rounded-lg py-2 px-4 block w-full appearance-none leading-normal mb-4",
                Attributes: []htmlrender.ElementAttr{
                    {
                        Name: "placeholder",
                        Content: "Title",
                    },
                },
            },
            {
                Tag: "button",
                ID: "submit-todo",
                ClassName: "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded",
                InnerText: "Add ToDo",
            },
        },
    }
}
