package models

import (
    "fmt"
    "strconv"
    "syscall/js"

    "../htmlrender"
    "../services"
)

type TodoItem struct {
    id    int64
    title string
    done  bool
}

// `TodoItemJson` is struct that used to transform todoItem to json
type TodoItemJson struct {
    ID    int64  `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

// Transform func: TodoItem -> TodoItemJson
func NewTodoItemJson(todoItem TodoItem) TodoItemJson {
    todoItemJson := new(TodoItemJson)
    todoItemJson.ID = todoItem.id
    todoItemJson.Title = todoItem.title
    todoItemJson.Done = todoItem.done
    return *todoItemJson
}

const (
    todoItemClassname = "todo-item"
    todoItemDeleteClassname = "todo-delete"
    todoItemDoneClassname = "todo-done"
    dataTodoId = "data-todo-id"
)

func (todoItem *TodoItem) GetItemIdClassname() string {
    return fmt.Sprintf("%s-%d", todoItemClassname, todoItem.id)
}

func (todoItem *TodoItem) GetTodoItemDeleteClassname() string {
    return todoItemDeleteClassname
}

func (todoItem *TodoItem) GetTodoItemDoneClassname() string {
    return todoItemDoneClassname
}

func (todoItem * TodoItem) GetTodoIdFromEl(el js.Value) int64 {
    todoIdStr := el.Get("dataset").Get("todoId").String()
    todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
    return todoId
}

func (todoItem *TodoItem) SetDone(done bool) {
    todoItem.done = done
}

func (todoItem *TodoItem) GetDone() bool {
    return todoItem.done
}

func (todoItem *TodoItem) SetTitle(title string) {
    todoItem.title = title
}

func (todoItem *TodoItem) GetTitle() string {
    return todoItem.title
}

func (todoItem *TodoItem) GetElementDef() htmlrender.ElementDef {
    deleteBtn := Button{
        Text:      "Delete",
        BgColor:   "orange",
        Size:       ButtonSizes.XS,
        ClassName: services.Classnames(
            "mr-1",
            todoItemDeleteClassname,
        ),
        Attributes: []htmlrender.ElementAttr{
            {
                Name: dataTodoId,
                Content: strconv.FormatInt(todoItem.id, 10),
            },
        },
    }
    doneBtn := Button{
        Text:       "Done",
        BgColor:    "green",
        Size:       ButtonSizes.XS,
        ClassName:  todoItemDoneClassname,
        Attributes: []htmlrender.ElementAttr{
            {
                Name: dataTodoId,
                Content: strconv.FormatInt(todoItem.id, 10),
            },
        },
    }
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: services.Classnames(
            todoItemClassname,
            todoItem.GetItemIdClassname(),
            "p-2 border-b-2 border-gray-200 flex justify-between",
        ),
        Children: []htmlrender.ElementDef{
            {
                Tag: "span",
                InnerText: todoItem.title,
            },
            {
                Tag: "div",
                Children: []htmlrender.ElementDef{
                    deleteBtn.GetElementDef(),
                    doneBtn.GetElementDef(),
                },
            },
        },
    }
}
