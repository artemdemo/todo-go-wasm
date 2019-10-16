package htmlrender

type DomEl struct {
    // GeneralEl
    GeneralParentEl
}

func (domEl *DomEl) SetAttribute(attrName string, value string) {
    domEl.El.Call("setAttribute", attrName, value)
}
