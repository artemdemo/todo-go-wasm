package htmlrender

type DomEl struct {
    generalParentEl
}

func (domEl *DomEl) SetAttribute(attrName string, value string) {
    el := domEl.GetEl()
    el.Call("setAttribute", attrName, value)
}
