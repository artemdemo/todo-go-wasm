package htmlrender

import "syscall/js"

type Event struct {
    ev js.Value
}

func (e *Event) GetTarget() interface{} {
    el := e.ev.Get("target")
    return wrapEl(el)
}
