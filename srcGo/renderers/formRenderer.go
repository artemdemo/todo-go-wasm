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
    // "dummyForm" will be used here to retrieve locator classnames
    dummyForm    models.Form
}

const (
    formParentClassName = "form"
)

func NewFormRenderer(documentEl js.Value) *FormRenderer {
    formR := new(FormRenderer)
    formR.formParentEl = htmlrender.GetFirstElementByClass(documentEl, formParentClassName)
    formR.dummyForm = models.Form{}
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
    this.submitBtnEl = htmlrender.GetFirstElementByClass(
        documentEl,
        this.dummyForm.GetAddTodoButtonClassname(),
    )
    this.titleInputEl = htmlrender.GetFirstElementByClass(
        documentEl,
        this.dummyForm.GetTodoTitleInputClassname(),
    )
}

func (this *FormRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: formParentClassName,
    }
}
