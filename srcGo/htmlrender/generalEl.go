package htmlrender

import (
    "strings"
    "syscall/js"
)

type GeneralEl struct {
    el
}

func (genEl *GeneralEl) AddEventListener(evtType string, cb func(evt Event)) {
    eventCb := func(this js.Value, args []js.Value) interface{} {
        target := args[0].Get("target")
        cb(Event{target})
        return nil
    }
    genEl.El.Call("addEventListener", evtType, js.FuncOf(eventCb))
}

func (genEl *GeneralEl) IsDefined() bool {
    return genEl.El.Type() == js.TypeUndefined
}

func (genEl *GeneralEl) HasClass(className string) bool  {
    haystack := genEl.El.Get("className").String()
    return strings.Contains(haystack, className)
}
