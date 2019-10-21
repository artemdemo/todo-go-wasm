package models

import (
    "testing"
)

func Test_GetElementDef(t *testing.T) {
    emptyBtn := Button{
        className:  "",
        text:       "",
        bgColor:    "",
        size:       "",
        attributes: nil,
    }
    emptyBtn_def := emptyBtn.GetElementDef()

    if emptyBtn_def.Tag != "button" {
        t.Fatalf("GetElementDef() failed, expected \"button\", got \"%v\"", emptyBtn_def.Tag)
    }
}
