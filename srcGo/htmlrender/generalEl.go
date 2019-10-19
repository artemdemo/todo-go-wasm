package htmlrender

import (
    "strings"
    "syscall/js"
)

type generalEl struct {
    el js.Value
}

func (genEl *generalEl) GetEl() js.Value {
    return genEl.el
}

func (genEl *generalEl) AddEventListener(evtType string, cb func(evt *Event)) {
    eventCb := func(this js.Value, args []js.Value) interface{} {
        evt := new(Event)
        evt.ev = args[0]
        cb(evt)
        return nil
    }
    genEl.el.Call("addEventListener", evtType, js.FuncOf(eventCb))
}

func (genEl *generalEl) IsDefined() bool {
    return genEl.el.Type() == js.TypeUndefined
}

func (genEl *generalEl) HasClass(className string) bool  {
    haystack := genEl.el.Get("className").String()
    return strings.Contains(haystack, className)
}
