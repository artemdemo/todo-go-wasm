package htmlrender

import (
    "syscall/js"
)

// ElementDef is DOM element definition
type ElementDef struct {
    Tag       string
    ClassName string
    ID string
    // It's not a good approach, to have `InnerText` defined on `ElementDef`
    // It makes it unclear how to use `children: []` property in the future
    // The better solution will be to defined text element in `children: []`
    // But I'll do it later
    InnerText string
}

// CreateElement is creating DOM element based on ElementDef
func CreateElement(document js.Value, elDef ElementDef) js.Value {
    el := document.Call("createElement", elDef.Tag)
    if elDef.ClassName != "" {
        el.Call("setAttribute", "class", elDef.ClassName)
    }
    if elDef.ID != "" {
        el.Call("setAttribute", "id", elDef.ID)
    }
    el.Set("innerText", elDef.InnerText)
    return el
}

// RenderElement is rendering DOM element in provided `baseEl`
func RenderElement(baseEl js.Value, el js.Value) {
    baseEl.Call("appendChild", el)
}
