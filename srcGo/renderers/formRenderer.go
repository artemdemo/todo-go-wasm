package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../models"
)

type submitCb func(title string)

type FormRenderer struct {
    // "formParentEl" is parent element where form itself will be rendered
    formParentEl js.Value
    submitBtnEl  js.Value
    titleInputEl js.Value
    onSubmitCb   submitCb
    // "dummyForm" will be used here to retrieve locator classnames
    dummyForm    models.Form
}

const (
    formParentClassName = "form"
)

func NewFormRenderer() *FormRenderer {
    formR := new(FormRenderer)
    formR.formParentEl = htmlrender.GetFirstElementByClass(
        htmlrender.GetDocumentEl(),
        formParentClassName,
    )
    formR.dummyForm = models.Form{}
    return formR
}

func (this *FormRenderer) OnSubmit(cb submitCb) {
    this.onSubmitCb = cb
}

func (this *FormRenderer) clickOnSubmit(js.Value, []js.Value) interface{} {
    this.onSubmitCb(
        this.titleInputEl.Get("value").String(),
    )
    this.titleInputEl.Set("value", "")
    return ""
}

func (this *FormRenderer) RenderForm(form models.Form) {
    documentEl := htmlrender.GetDocumentEl()
    htmlrender.RenderElement(
        this.formParentEl,
        form.GetElementDef(),
    )
    this.submitBtnEl = htmlrender.GetFirstElementByClass(
        documentEl,
        this.dummyForm.GetAddTodoButtonClassname(),
    )
    this.titleInputEl = htmlrender.GetFirstElementByClass(
        documentEl,
        this.dummyForm.GetTodoTitleInputClassname(),
    )
    this.submitBtnEl.Call("addEventListener", "click", js.FuncOf(this.clickOnSubmit))
}

func (this *FormRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: formParentClassName,
    }
}
