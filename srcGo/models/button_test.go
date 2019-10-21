package models

import (
    "testing"
)

func emptyButton(t *testing.T) {
    emptyBtn := Button{
        className:  "",
        text:       "",
        bgColor:    "",
        size:       "",
        attributes: nil,
    }
    emptyBtn_def := emptyBtn.GetElementDef()

    if emptyBtn_def.Tag != "button" {
        t.Fatalf("GetElementDef() [emptyButton() Tag] failed, expected \"button\", got \"%v\"", emptyBtn_def.Tag)
    }

    className_result := "py-2 px-4 text-white py-1 px-2 rounded"
    if emptyBtn_def.ClassName != className_result {
        t.Fatalf("GetElementDef() [emptyButton() ClassName] failed, expected \"%v\", got \"%v\"", className_result, emptyBtn_def.ClassName)
    }

    if emptyBtn_def.InnerText != "" {
        t.Fatalf("GetElementDef() [emptyButton() InnerText] failed, expected \"\", got \"%v\"", emptyBtn_def.InnerText)
    }

    if len(emptyBtn_def.Attributes) != 0 {
        t.Fatalf("GetElementDef() [emptyButton() len(Attributes)] failed, expected \"0\", got \"%v\"", (emptyBtn_def.Attributes))
    }
}

func className(t *testing.T) {
    btn := Button{
        className:  "some-class",
        text:       "",
        bgColor:    "",
        size:       "",
        attributes: nil,
    }
    btn_def := btn.GetElementDef()

    className_result := "some-class py-2 px-4 text-white py-1 px-2 rounded"
    if btn_def.ClassName != className_result {
        t.Fatalf("GetElementDef() [className() ClassName] failed, expected \"%v\", got \"%v\"", className_result, btn_def.ClassName)
    }
}

func text(t *testing.T) {
    btn := Button{
        className:  "",
        text:       "Some text",
        bgColor:    "",
        size:       "",
        attributes: nil,
    }
    btn_def := btn.GetElementDef()

    innerText_result := "Some text"
    if btn_def.InnerText != innerText_result {
        t.Fatalf("GetElementDef() [text() InnerText] failed, expected \"%v\", got \"%v\"", innerText_result, btn_def.InnerText)
    }
}

func bgColor(t *testing.T) {
    btn := Button{
        className:  "",
        text:       "",
        bgColor:    "blue",
        size:       "",
        attributes: nil,
    }
    btn_def := btn.GetElementDef()

    className_result := "bg-blue-500 hover:bg-blue-600 py-2 px-4 text-white py-1 px-2 rounded"
    if btn_def.ClassName != className_result {
        t.Fatalf("GetElementDef() [bgColor() ClassName] failed, expected \"%v\", got \"%v\"", className_result, btn_def.ClassName)
    }
}

func size(t *testing.T) {
    btn := Button{
        className:  "",
        text:       "",
        bgColor:    "",
        size:       ButtonSizes.XS,
        attributes: nil,
    }
    btn_def := btn.GetElementDef()

    className_result := "py-1 px-2 text-xs text-white py-1 px-2 rounded"
    if btn_def.ClassName != className_result {
        t.Fatalf("GetElementDef() [size() ClassName] failed, expected \"%v\", got \"%v\"", className_result, btn_def.ClassName)
    }
}

func Test_GetElementDef(t *testing.T) {
    emptyButton(t)
    className(t)
    text(t)
    bgColor(t)
    size(t)
}
