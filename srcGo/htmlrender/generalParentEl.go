package htmlrender

import (
    "strings"
    "syscall/js"
)

type GeneralParentEl struct {
    el js.Value
}

func (genParEl *GeneralParentEl) GetEl() js.Value {
    return genParEl.el
}

func (genParEl *GeneralParentEl) SetInnerText(text string) {
    genParEl.el.Set("innerText", text)
}

func (genParEl *GeneralParentEl) SetInnerHtml(html string) {
    genParEl.el.Set("innerHtml", html)
}

func (genParEl *GeneralParentEl) GetElementById(id string) interface{} {
    el := genParEl.el.Call("getElementById", id)
    return wrapEl(el)
}

func (genParEl *GeneralParentEl) GetFirstElementByClass(className string) interface{} {
    el := genParEl.el.Call("getElementsByClassName", className).Index(0)
    return wrapEl(el)
}

func (genParEl *GeneralParentEl) AppendChild(child interface{}) {
    switch childEl := child.(type) {
    case DomEl:
    case *DomEl:
        genParEl.el.Call("appendChild", childEl.GetEl())
    case ElementDef:
        domEl := CreateElement(childEl)
        genParEl.el.Call("appendChild", domEl.GetEl())
    default:
        panic("Unknown child type")
    }
}

func (genParEl *GeneralParentEl) ReplaceInDOM(newEl interface{}) {
    switch _el := newEl.(type) {
    case DomEl:
    case *DomEl:
        genParEl.el.Get("parentNode").Call("replaceChild", _el.el, genParEl.el)
    case ElementDef:
        domEl := CreateElement(_el)
        genParEl.el.Get("parentNode").Call("replaceChild", domEl.el, genParEl.el)
    default:
        panic("Unknown child type")
    }
}

////

func (genParEl *GeneralParentEl) RemoveElFromDOM() {
    genParEl.el.Call("remove")
}

func (genParEl *GeneralParentEl) AddEventListener(evtType string, cb func(evt *Event)) {
    eventCb := func(this js.Value, args []js.Value) interface{} {
        evt := new(Event)
        evt.ev = args[0]
        cb(evt)
        return nil
    }
    genParEl.el.Call("addEventListener", evtType, js.FuncOf(eventCb))
}

func (genParEl *GeneralParentEl) IsDefined() bool {
    return genParEl.el.Type() != js.TypeUndefined
}

func (genParEl *GeneralParentEl) HasClass(className string) bool  {
    haystack := genParEl.el.Get("className").String()
    return strings.Contains(haystack, className)
}
