package htmlrender

import "syscall/js"

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

var documentEL = NewDocumentEl()

// createElement is creating DOM element based on ElementDef
func createElement(elDef ElementDef) *DomEl {
    var el *DomEl
    // If there is no Tag name, then it's text node
    // and text node can't have attributes or children
    if elDef.Tag != "" {
        el = documentEL.CreateElement(elDef.Tag)
        if elDef.ClassName != "" {
            el.SetAttribute("class", elDef.ClassName)
        }
        if elDef.ID != "" {
            el.SetAttribute("id", elDef.ID)
        }
        if len(elDef.Attributes) > 0 {
            attributesAmount := len(elDef.Attributes)
            for i := 0; i < attributesAmount; i++ {
                attr := elDef.Attributes[i]
                el.SetAttribute(attr.Name, attr.Content)
            }
        }
        // If there are both Tag name and InnerText
        // then innerText will come before Children
        if elDef.InnerText != "" {
            el.SetInnerText(elDef.InnerText)
        } else if len(elDef.Children) > 0 {
            childrenAmount := len(elDef.Children)
            for i := 0; i < childrenAmount; i++ {
                childEl := createElement(elDef.Children[i])
                el.AppendChild(childEl)
            }
        }
        // If there is InnerText, without Tag,
        // then I'll render text node
    } else if elDef.InnerText != "" {
        el = documentEL.CreateTextNode(elDef.InnerText)
    }
    return el
}

func (elementDef ElementDef) GetEl() js.Value {
    el := createElement(elementDef)
    return el.GetEl()
}

// Get html representation of defined element
func (elementDef ElementDef) GetHTML() string {
    // I'm creating container, in order to capture html of the whole element
    elContainer := documentEL.CreateElement("div")
    elContainer.AppendChild(elementDef)
    return elContainer.GetEl().Get("innerHTML").String()
}
