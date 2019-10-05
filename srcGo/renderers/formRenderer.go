package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../models"
)

type FormRenderer struct {
    // "formParentEl" is parent element where form itself will be rendered
    formParentEl js.Value
    submitBtnEl  js.Value
    titleInputEl js.Value
}

var formParentClassName = "form"

func NewFormRenderer(documentEl js.Value) *FormRenderer {
    formR := new(FormRenderer)
    formR.formParentEl = htmlrender.GetFirstElementByClass(documentEl, formParentClassName)
    return formR
}

func (this *FormRenderer) ClearTitleInput() {
    this.titleInputEl.Set("value", "")
}

func (this *FormRenderer) GetTitle() string {
    return this.titleInputEl.Get("value").String()
}

func (this *FormRenderer) OnSubmitCb(cb func(js.Value, []js.Value) interface{}) {
    this.submitBtnEl.Call("addEventListener", "click", js.FuncOf(cb))
}

func (this *FormRenderer) RenderForm(documentEl js.Value,
                                     form models.Form) {
    htmlrender.RenderElement(
        this.formParentEl,
        htmlrender.CreateElement(
            documentEl,
            form.GetElementDef(),
        ),
    )
    this.submitBtnEl = htmlrender.GetElementById(documentEl, "submit-todo")
    this.titleInputEl = htmlrender.GetElementById(documentEl, "todo-title")
}

func (this *FormRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: formParentClassName,
    }
}
