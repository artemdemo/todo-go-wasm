package htmlrender

import (
    "strings"
    "syscall/js"
)

type GeneralParentEl struct {
    El js.Value
}

func wrapEl(el js.Value) interface{} {
    if el.Type() != js.TypeUndefined {
        tagName := strings.ToLower(el.Get("tagName").String())
        if tagName == "input" {
            return InputEl{El: el}
        }
        return DomEl{El: el}
    }
    return nil
}

func (genParEl *GeneralParentEl) SetInnerText(text string) {
    genParEl.El.Set("innerText", text)
}

func (genParEl *GeneralParentEl) SetInnerHtml(html string) {
    genParEl.El.Set("innerHtml", html)
}

func (genParEl *GeneralParentEl) GetElementById(id string) interface{} {
    el := genParEl.El.Call("getElementById", id)
    return wrapEl(el)
}

func (genParEl *GeneralParentEl) GetFirstElementByClass(className string) interface{} {
    el := genParEl.El.Call("getElementsByClassName", className).Index(0)
    return wrapEl(el)
}

func (genParEl *GeneralParentEl) AppendChild(child interface{}) {
    switch childEl := child.(type) {
    case DomEl:
    case GeneralEl:
        genParEl.El.Call("appendChild", childEl.El)
    case ElementDef:
        domEl := CreateElement(childEl)
        genParEl.El.Call("appendChild", domEl.El)
    default:
        panic("Unknown child type")
    }
}
