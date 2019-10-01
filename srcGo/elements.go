package main

import (
    "syscall/js"

    "./htmlrender"
)

var _documentEL js.Value
var _appLoggerEL js.Value

func getDocumentEl() js.Value {
    if _documentEL.Type() == js.TypeUndefined {
        _documentEL = js.Global().Get("document")
    }
    return _documentEL
}

func getAppLoggerEl() js.Value {
    if _appLoggerEL.Type() == js.TypeUndefined {
        documentEl := getDocumentEl()
        _appLoggerEL = htmlrender.GetFirstElementByClass(documentEl, "app-logger")
    }
    return _appLoggerEL
}
