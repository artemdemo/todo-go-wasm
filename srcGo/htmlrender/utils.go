package htmlrender

import (
    "strings"
    "syscall/js"
)

func GetFirstElementByClass(baseEl js.Value, className string) js.Value {
    return baseEl.Call("getElementsByClassName", className).Index(0)
}

func GetElementById(baseEl js.Value, id string) js.Value {
    return baseEl.Call("getElementById", id)
}

func ElementHasClass(el js.Value, needle string) bool {
    haystack := el.Get("className").String()
    return strings.Contains(haystack, needle)
}
