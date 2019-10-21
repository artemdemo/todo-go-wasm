package htmlrender

var documentEL = NewDocumentEl()

// CreateElement is creating DOM element based on ElementDef
func CreateElement(elDef ElementDef) *DomEl {
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
                childEl := CreateElement(elDef.Children[i])
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
