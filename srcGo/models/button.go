package models

import (
    "fmt"

    "../htmlrender"
    "../services"
)

type Button struct {
    ClassName string
    Text      string
    BgColor   string
    Attributes []htmlrender.ElementAttr
}

const (
    buttonPadding = "py-1 px-2"
    buttonRounded = "rounded"
)

func (button *Button) getBgColors(color string) string {
    return fmt.Sprintf(
        "bg-%s-500 hover:bg-%s-600",
        color,
        color,
    )
}

func (button *Button) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag:        "button",
        ClassName:  services.Classnames(
            button.ClassName,
            button.getBgColors(button.BgColor),
            "text-white",
            buttonPadding,
            buttonRounded,
        ),
        InnerText:  button.Text,
        Attributes: button.Attributes,
    }
}
