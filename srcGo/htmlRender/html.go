package htmlrender

import (
    "syscall/js"
)

// ElementDef is DOM element definition
type ElementDef struct {
    Tag       string
    ClassName string
}

// CreateElement is creating DOM element based on ElementDef
func CreateElement(document js.Value, elDef ElementDef) js.Value {
    el := document.Call("createElement", elDef.Tag)
    el.Call("setAttribute", "class", elDef.ClassName)
    return el
}

// RenderElement is rendering DOM element in provided `baseEl`
func RenderElement(baseEl js.Value, el js.Value) {
    baseEl.Call("appendChild", el)
}
