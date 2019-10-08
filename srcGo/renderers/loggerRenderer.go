package renderers

import (
    "syscall/js"

    "../htmlrender"
    "../services"
)

type LoggerRenderer struct {
    // "loggerParentEl" is parent element where form itself will be rendered
    loggerParentEl js.Value
}

const (
    appLoggerClassname = "app-logger"
)

func NewLoggerRenderer(documentEl js.Value) *LoggerRenderer {
    loggerR := new(LoggerRenderer)
    loggerR.loggerParentEl = htmlrender.GetFirstElementByClass(documentEl, appLoggerClassname)
    return loggerR
}

func (this *LoggerRenderer) AppendLogMsg(documentEl js.Value, msg string) {
    htmlrender.RenderElement(
        this.loggerParentEl,
        htmlrender.CreateElement(
            documentEl,
            htmlrender.ElementDef{
                Tag: "p",
                Children: []htmlrender.ElementDef{
                    { InnerText: msg },
                },
            },
        ),
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
