package htmlrender

import (
    "syscall/js"
)

// ElementAttr is an DOM element attribute
type ElementAttr struct {
    Name    string
    Content string
}

// ElementDef is DOM element definition
type ElementDef struct {
    Tag        string
    ClassName  string
    ID         string
    // InnerText, has higher priority than Children
    // In case user passed both, then InnerText will be rendered and Children will be ignored
    InnerText  string
    // Attributes will override ClassName and ID
    // (in case user used theme here as well)
    Attributes []ElementAttr
    Children   []ElementDef
}

var _documentEL js.Value

func GetDocumentEl() js.Value {
    if _documentEL.Type() == js.TypeUndefined {
        _documentEL = js.Global().Get("document")
    }
    return _documentEL
}

// CreateElement is creating DOM element based on ElementDef
func CreateElement(elDef ElementDef) js.Value {
    var el js.Value
    var documentEl = GetDocumentEl()
    // If there is no Tag name, then it's text node
    // and text node can't have attributes or children
    if elDef.Tag != "" {
        el = documentEl.Call("createElement", elDef.Tag)
        if elDef.ClassName != "" {
            el.Call("setAttribute", "class", elDef.ClassName)
        }
        if elDef.ID != "" {
            el.Call("setAttribute", "id", elDef.ID)
        }
        if len(elDef.Attributes) > 0 {
            attributesAmount := len(elDef.Attributes)
            for i := 0; i < attributesAmount; i++ {
                attr := elDef.Attributes[i]
                el.Call("setAttribute", attr.Name, attr.Content)
            }
        }
        // If there are both Tag name and InnerText
        // then innerText will come before Children
        if elDef.InnerText != "" {
            el.Set("innerText", elDef.InnerText)
        } else if len(elDef.Children) > 0 {
            childrenAmount := len(elDef.Children)
            for i := 0; i < childrenAmount; i++ {
                childEl := CreateElement(elDef.Children[i])
                el.Call("appendChild", childEl)
            }
        }
    // If there is InnerText, without Tag,
    // then I'll render text node
    } else if elDef.InnerText != "" {
        el = documentEl.Call("createTextNode", elDef.InnerText)
    }
    return el
}

// RenderElement is rendering DOM element in provided `baseEl`
func RenderElement(baseEl js.Value, elDef ElementDef) {
    baseEl.Call("appendChild", CreateElement(elDef))
}

// ClearElementContent is clearing element content
func ClearElementContent(el js.Value) {
    el.Set("innerHtml", "")
}
