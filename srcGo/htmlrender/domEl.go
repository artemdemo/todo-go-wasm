package htmlrender

type DomEl struct {
    // GeneralEl
    GeneralParentEl
}

func (domEl *DomEl) SetAttribute(attrName string, value string) {
    el := domEl.GetEl()
    el.Call("setAttribute", attrName, value)
}
