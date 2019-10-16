package htmlrender

import "syscall/js"

type documentEl struct {
    GeneralParentEl
}

func NewDocumentEl() *documentEl {
    docEl := new(documentEl)
    docEl.El = js.Global().Get("document")
    return docEl
}

func (docEl *documentEl) CreateElement(tagName string) *DomEl {
    domEl := new(DomEl)
    domEl.El = docEl.El.Call("createElement", tagName)
    return domEl
}

func (docEl *documentEl) CreateTextNode(text string) *DomEl {
    domEl := new(DomEl)
    domEl.El = docEl.El.Call("createTextNode", text)
    return domEl
}
