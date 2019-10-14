package htmlrender

import "syscall/js"

type DomEl struct {
    *GeneralEl
    *GeneralParentEl
    El js.Value
}

func (domEl *DomEl) SetAttribute(attrName string, value string) {
    domEl.El.Call("setAttribute", attrName, value)
}
