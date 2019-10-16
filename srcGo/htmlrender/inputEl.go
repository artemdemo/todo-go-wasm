package htmlrender

import "syscall/js"

type InputEl struct {
    GeneralEl
}

func (inputEl *InputEl) GetEl() js.Value {
    return inputEl.el
}

func (inputEl *InputEl) GetValue() string {
    return inputEl.el.Get("value").String()
}

func (inputEl *InputEl) SetValue(value string) {
    inputEl.el.Set("value", value)
}
