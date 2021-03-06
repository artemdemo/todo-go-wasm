package renderers

import (
    "fmt"

    "../htmlrender"
    "../services"
)

type LoggerRenderer struct {
    // "loggerParentEl" is parent element where form itself will be rendered
    loggerParentEl *htmlrender.DomEl
}

const (
    appLoggerClassname = "app-logger"
)

func NewLoggerRenderer() *LoggerRenderer {
    loggerR := new(LoggerRenderer)
    loggerParent := htmlrender.NewDocumentEl().GetFirstElementByClass(appLoggerClassname)
    if loggerParentEl, ok := loggerParent.(*htmlrender.DomEl); ok {
        loggerR.loggerParentEl = loggerParentEl
    } else {
        fmt.Printf("loggerParent is not of type *htmlrender.DomEl, got %T instead\n", loggerParent)
        panic("loggerParent is not of type *htmlrender.DomEl")
    }
    return loggerR
}

func (this *LoggerRenderer) AppendLogMsg(msg string) {
    this.loggerParentEl.AppendChild(
        htmlrender.ElementDef{
            Tag: "p",
            Children: []htmlrender.ElementDef{
                { InnerText: msg },
            },
        },
    )
}

func (this *LoggerRenderer) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag: "div",
        ClassName: services.Classnames(
            appLoggerClassname,
            "rounded bg-gray-100 p-3 text-gray-500",
        ),
    }
}
