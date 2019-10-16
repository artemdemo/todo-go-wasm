package htmlrender

import (
    "strings"
    "syscall/js"
)

type GeneralParentEl struct {
    el
}

func wrapEl(el js.Value) interface{} {
    if el.Type() != js.TypeUndefined {
        tagName := strings.ToLower(el.Get("tagName").String())
        if tagName == "input" {
            return InputEl{El: el}
        }
        domEl := new(DomEl)
        domEl.El = el
        return domEl
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
    case *DomEl:
    case GeneralEl:
    case *GeneralEl:
        genParEl.El.Call("appendChild", childEl.El)
    case ElementDef:
        domEl := CreateElement(childEl)
        genParEl.El.Call("appendChild", domEl.El)
    default:
        panic("Unknown child type")
    }
}

////

func (genParEl *GeneralParentEl) AddEventListener(evtType string, cb func(evt Event)) {
    eventCb := func(this js.Value, args []js.Value) interface{} {
        target := args[0].Get("target")
        cb(Event{target})
        return nil
    }
    genParEl.El.Call("addEventListener", evtType, js.FuncOf(eventCb))
}

func (genParEl *GeneralParentEl) IsDefined() bool {
    return genParEl.El.Type() == js.TypeUndefined
}

func (genParEl *GeneralParentEl) HasClass(className string) bool  {
    haystack := genParEl.El.Get("className").String()
    return strings.Contains(haystack, className)
}
