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

func (formRenderer *FormRenderer) getFormParentEl(baseEl js.Value) js.Value {
    if formRenderer.formParentEl.Type() == js.TypeUndefined {
        formRenderer.formParentEl = htmlrender.GetFirstElementByClass(baseEl, formParentClassName)
    }
    return formRenderer.formParentEl
}

func (formRenderer *FormRenderer) getSubmitBtnEl(baseEl js.Value) js.Value {
    if formRenderer.submitBtnEl.Type() == js.TypeUndefined {
        formRenderer.submitBtnEl = htmlrender.GetElementById(baseEl, "submit-todo")
    }
    return formRenderer.submitBtnEl
}

func (formRenderer *FormRenderer) getTitleInputEl(baseEl js.Value) js.Value {
    if formRenderer.titleInputEl.Type() == js.TypeUndefined {
        formRenderer.titleInputEl = htmlrender.GetElementById(baseEl, "todo-title")
    }
    return formRenderer.titleInputEl
}

func (formRenderer *FormRenderer) ClearTitleInput(baseEl js.Value) {
    formRenderer.getTitleInputEl(baseEl).Set("value", "")
}

func (formRenderer *FormRenderer) GetTitle(baseEl js.Value) string {
    return formRenderer.getTitleInputEl(baseEl).Get("value").String()
}

func (formRenderer *FormRenderer) OnSubmitCb(baseEl js.Value,
                                             cb func(js.Value, []js.Value) interface{}) {
    submitBtnEl := formRenderer.getSubmitBtnEl(baseEl)
    submitBtnEl.Call("addEventListener", "click", js.FuncOf(cb))
}

func (formRenderer *FormRenderer) RenderForm(documentEl js.Value,
                                             form models.Form) {
    formParentEl := formRenderer.getFormParentEl(documentEl)
    htmlrender.RenderElement(
        formParentEl,
        htmlrender.CreateElement(
            documentEl,
            form.GetElementDef(),
        ),
    )
}

func (formRenderer FormRenderer) GetBaseElDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: formParentClassName,
    }
}
