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
    Children []ElementDef
}

// CreateElement is creating DOM element based on ElementDef
func CreateElement(document js.Value, elDef ElementDef) js.Value {
    var el js.Value
    // If there is no Tag name, then it's text node
    // and text node can't have attributes or children
    if elDef.Tag != "" {
        el = document.Call("createElement", elDef.Tag)
        if elDef.ClassName != "" {
            el.Call("setAttribute", "class", elDef.ClassName)
        }
        if elDef.ID != "" {
            el.Call("setAttribute", "id", elDef.ID)
        }
        // If there are both Tag name and InnerText
        // then innerText will come before Children
        if elDef.InnerText != "" {
            el.Set("innerText", elDef.InnerText)
        } else if len(elDef.Children) > 0 {
            childrenAmount := len(elDef.Children)
            for i := 0; i < childrenAmount; i++ {
                childEl := CreateElement(document, elDef.Children[i])
                el.Call("appendChild", childEl)
            }
        }
    // If there is InnerText, without Tag,
    // then I'll render text node
    } else if elDef.InnerText != "" {
        el = document.Call("createTextNode", elDef.InnerText)
    }
    return el
}

// RenderElement is rendering DOM element in provided `baseEl`
func RenderElement(baseEl js.Value, el js.Value) {
    baseEl.Call("appendChild", el)
}
