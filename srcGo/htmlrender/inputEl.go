package htmlrender

import "syscall/js"

type InputEl struct {
    *GeneralEl
    El js.Value
}

func (inputEl *InputEl) GetValue() string {
    return inputEl.El.Get("value").String()
}

func (inputEl *InputEl) SetValue(value string) {
    inputEl.El.Set("value", value)
}
