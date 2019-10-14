package renderers

import (
    "../htmlrender"
    "../services"
)

type LoggerRenderer struct {
    // "loggerParentEl" is parent element where form itself will be rendered
    loggerParentEl htmlrender.DomEl
}

const (
    appLoggerClassname = "app-logger"
)

func NewLoggerRenderer() *LoggerRenderer {
    loggerR := new(LoggerRenderer)
    el := (htmlrender.DocumentEl{}).GetFirstElementByClass(appLoggerClassname)
    loggerParentEl, ok := el.(htmlrender.DomEl)
    if ok {
        loggerR.loggerParentEl = loggerParentEl
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
