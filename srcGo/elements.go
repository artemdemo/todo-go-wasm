package main

import (
    "syscall/js"
)

var _documentEL js.Value

func getDocumentEl() js.Value {
    if _documentEL.Type() == js.TypeUndefined {
        _documentEL = js.Global().Get("document")
    }
    return _documentEL
}
