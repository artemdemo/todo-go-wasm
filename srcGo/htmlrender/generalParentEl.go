package htmlrender

import (
    "strings"
    "syscall/js"
)

type generalParentEl struct {
    el js.Value
}

func (genParEl *generalParentEl) GetEl() js.Value {
    return genParEl.el
}

func (genParEl *generalParentEl) SetInnerText(text string) {
    genParEl.el.Set("innerText", text)
}

func (genParEl *generalParentEl) SetInnerHtml(html string) {
    genParEl.el.Set("innerHTML", html)
}

func (genParEl *generalParentEl) GetElementById(id string) interface{} {
    el := genParEl.el.Call("getElementById", id)
    return wrapEl(el)
}

func (genParEl *generalParentEl) GetFirstElementByClass(className string) interface{} {
    el := genParEl.el.Call("getElementsByClassName", className).Index(0)
    return wrapEl(el)
}

func (genParEl *generalParentEl) AppendChild(childEl generalElI) {
    genParEl.el.Call("appendChild", childEl.GetEl())
}

func (genParEl *generalParentEl) ReplaceInDOM(newEl generalElI) {
    genParEl.el.Get("parentNode").Call("replaceChild", newEl.GetEl(), genParEl.el)
}

////

func (genParEl *generalParentEl) RemoveElFromDOM() {
    genParEl.el.Call("remove")
}

func (genParEl *generalParentEl) AddEventListener(evtType string, cb func(evt *Event)) {
    eventCb := func(this js.Value, args []js.Value) interface{} {
        evt := new(Event)
        evt.ev = args[0]
        cb(evt)
        return nil
    }
    genParEl.el.Call("addEventListener", evtType, js.FuncOf(eventCb))
}

func (genParEl *generalParentEl) IsDefined() bool {
    return genParEl.el.Type() != js.TypeUndefined
}

func (genParEl *generalParentEl) HasClass(className string) bool  {
    haystack := genParEl.el.Get("className").String()
    return strings.Contains(haystack, className)
}
