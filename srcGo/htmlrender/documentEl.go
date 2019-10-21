package htmlrender

import "syscall/js"

type documentEl struct {
    generalParentEl
}

func NewDocumentEl() *documentEl {
    docEl := new(documentEl)
    docEl.el = js.Global().Get("document")
    return docEl
}

func (docEl *documentEl) CreateElement(tagName string) *DomEl {
    domEl := new(DomEl)
    if !docEl.IsDefined() {
        panic("document element is not defined")
    }
    domEl.el = docEl.el.Call("createElement", tagName)
    return domEl
}

func (docEl *documentEl) CreateTextNode(text string) *DomEl {
    domEl := new(DomEl)
    domEl.el = docEl.el.Call("createTextNode", text)
    return domEl
}
