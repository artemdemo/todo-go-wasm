package htmlrender

import (
    "strings"
    "syscall/js"
)

type GeneralEl struct {
    El js.Value
}

func (genEl *GeneralEl) AddEventListener(evtType string, cb js.Func) {
    genEl.El.Call("addEventListener", evtType, cb)
}

func (genEl *GeneralEl) IsDefined() bool {
    return genEl.El.Type() == js.TypeUndefined
}

func (genEl *GeneralEl) ElementHasClass(className string) bool  {
    haystack := genEl.El.Get("className").String()
    return strings.Contains(haystack, className)
}
