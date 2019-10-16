package renderers

import (
    "../htmlrender"
    "../models"
    "fmt"
)

type submitCb func(title string)

type FormRenderer struct {
    // "formParentEl" is parent element where form itself will be rendered
    formParentEl *htmlrender.DomEl
    submitBtnEl  *htmlrender.DomEl
    titleInputEl *htmlrender.InputEl
    onSubmitCb   submitCb
    // "dummyForm" will be used here to retrieve locator classnames
    dummyForm    models.Form
}

const (
    formParentClassName = "form"
)

func NewFormRenderer() *FormRenderer {
    formR := new(FormRenderer)
    formParent := htmlrender.NewDocumentEl().GetFirstElementByClass(formParentClassName)
    if formParentEl, ok := formParent.(*htmlrender.DomEl); ok {
        formR.formParentEl = formParentEl
    } else {
        fmt.Printf("formParent is not of type *htmlrender.DomEl, got %T instead\n", formParent)
        panic("formParent is not of type *htmlrender.DomEl")
    }
    formR.dummyForm = models.Form{}
    return formR
}

func (this *FormRenderer) OnSubmit(cb submitCb) {
    this.onSubmitCb = cb
}

func (this *FormRenderer) clickOnSubmit(evt *htmlrender.Event) {
    this.onSubmitCb(
        this.titleInputEl.GetValue(),
    )
    this.titleInputEl.SetValue("")
}

func (this *FormRenderer) RenderForm(form models.Form) {
    this.formParentEl.AppendChild(
        form.GetElementDef(),
    )

    submitBtn := htmlrender.NewDocumentEl().GetFirstElementByClass(
        this.dummyForm.GetAddTodoButtonClassname(),
    )
    if submitBtnEl, ok := submitBtn.(*htmlrender.DomEl); ok {
        this.submitBtnEl = submitBtnEl
        this.submitBtnEl.AddEventListener("click", this.clickOnSubmit)
    } else {
        fmt.Printf("submitBtn is not of type *htmlrender.DomEl, got %T instead\n", submitBtn)
        panic("submitBtn is not of type *htmlrender.DomEl")
    }

    titleInput := htmlrender.NewDocumentEl().GetFirstElementByClass(
        this.dummyForm.GetTodoTitleInputClassname(),
    )
    if titleInputEl, ok := titleInput.(*htmlrender.InputEl); ok {
        this.titleInputEl = titleInputEl
    } else {
        fmt.Printf("titleInput is not of type *htmlrender.InputEl, got %T instead\n", titleInput)
        panic("titleInput is not of type *htmlrender.InputEl")
    }
}

func (this *FormRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: formParentClassName,
    }
}
