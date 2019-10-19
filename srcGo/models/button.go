package models

import (
    "fmt"

    "../htmlrender"
    "../services"
)

type Button struct {
    className string
    text      string
    bgColor   string
    size      string
    attributes []htmlrender.ElementAttr
}

type buttonSizes struct {
    DEFAULT string
    XS      string
}

var ButtonSizes = &buttonSizes{
    DEFAULT: "",
    XS:      "xs",
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

func (button *Button) getPadding() string {
    switch button.size {
    case ButtonSizes.XS:
        return "py-1 px-2"
    default:
        return "py-2 px-4"
    }
}

func (button *Button) getTextSize() string {
    switch button.size {
    case ButtonSizes.XS:
        return "text-xs"
    default:
        return ""
    }
}

func (button *Button) GetElementDef() htmlrender.ElementDef {
    return htmlrender.ElementDef{
        Tag:        "button",
        ClassName:  services.Classnames(
            button.className,
            button.getBgColors(button.bgColor),
            button.getPadding(),
            button.getTextSize(),
            "text-white",
            buttonPadding,
            buttonRounded,
        ),
        InnerText:  button.text,
        Attributes: button.attributes,
    }
}
