package htmlrender

import "syscall/js"

type DocumentEl struct {
    *GeneralParentEl
    El js.Value
}

func (docEl *DocumentEl) getDocumentEl() js.Value {
    return js.Global().Get("document")
}

func (docEl *DocumentEl) CreateElement(tagName string) DomEl {
    return DomEl{
        El: docEl.getDocumentEl().Call("createElement", tagName),
    }
}

func (docEl *DocumentEl) CreateTextNode(text string) DomEl {
    return DomEl{
        El: docEl.getDocumentEl().Call("createTextNode", text),
    }
}
