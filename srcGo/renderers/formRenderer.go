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

func (this *FormRenderer) getFormParentEl(baseEl js.Value) js.Value {
    if this.formParentEl.Type() == js.TypeUndefined {
        this.formParentEl = htmlrender.GetFirstElementByClass(baseEl, formParentClassName)
    }
    return this.formParentEl
}

func (this *FormRenderer) getSubmitBtnEl(baseEl js.Value) js.Value {
    if this.submitBtnEl.Type() == js.TypeUndefined {
        this.submitBtnEl = htmlrender.GetElementById(baseEl, "submit-todo")
    }
    return this.submitBtnEl
}

func (this *FormRenderer) getTitleInputEl(baseEl js.Value) js.Value {
    if this.titleInputEl.Type() == js.TypeUndefined {
        this.titleInputEl = htmlrender.GetElementById(baseEl, "todo-title")
    }
    return this.titleInputEl
}

func (this *FormRenderer) ClearTitleInput(params ...js.Value) {
    if len(params) == 0 {
        this.titleInputEl.Set("value", "")
    } else {
        baseEl := params[0]
        this.getTitleInputEl(baseEl).Set("value", "")
    }
}

func (this *FormRenderer) GetTitle(baseEl js.Value) string {
    return this.getTitleInputEl(baseEl).Get("value").String()
}

func (this *FormRenderer) OnSubmitCb(baseEl js.Value,
                                     cb func(js.Value, []js.Value) interface{}) {
    submitBtnEl := this.getSubmitBtnEl(baseEl)
    submitBtnEl.Call("addEventListener", "click", js.FuncOf(cb))
}

func (this *FormRenderer) RenderForm(documentEl js.Value,
                                     form models.Form) {
    formParentEl := this.getFormParentEl(documentEl)
    htmlrender.RenderElement(
        formParentEl,
        htmlrender.CreateElement(
            documentEl,
            form.GetElementDef(),
        ),
    )
}

func (this FormRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: formParentClassName,
    }
}
