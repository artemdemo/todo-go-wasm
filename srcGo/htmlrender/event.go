package htmlrender

import "syscall/js"

type Event struct {
    Ev js.Value
}

func (e *Event) GetTarget() interface{} {
    el := e.Ev.Get("target")
    return wrapEl(el)
}
