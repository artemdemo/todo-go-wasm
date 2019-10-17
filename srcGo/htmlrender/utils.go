package htmlrender

import (
    "strings"
    "syscall/js"
)

func wrapEl(el js.Value) interface{} {
    if el.Type() != js.TypeUndefined {
        tagName := strings.ToLower(el.Get("tagName").String())
        if tagName == "input" {
            inputEl := new(InputEl)
            inputEl.el = el
            return inputEl
        }
        domEl := new(DomEl)
        domEl.el = el
        return domEl
    }
    return nil
}
