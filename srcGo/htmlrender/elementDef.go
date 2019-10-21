package htmlrender

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
