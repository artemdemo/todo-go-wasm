package htmlrender

import "syscall/js"

func GetFirstElementByClass(baseEl js.Value, className string) js.Value {
    return baseEl.Call("getElementsByClassName", className).Index(0)
}
