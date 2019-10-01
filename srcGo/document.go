package main

import "syscall/js"

var _document js.Value

func getDocumentEl() js.Value {
    if _document.Type() == js.TypeUndefined {
        _document = js.Global().Get("document")
    }
    return _document
}
