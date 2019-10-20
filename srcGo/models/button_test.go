package models

import (
    "testing"
)

func TestGetElementDef(t *testing.T) {
    emptyBtn := Button{
        className:  "",
        text:       "",
        bgColor:    "",
        size:       "",
        attributes: nil,
    }
    emptyBtn_def := emptyBtn.GetElementDef()

    if emptyBtn_def.Tag != "button" {
        t.Errorf("GetElementDef() failed, expected \"button\", got \"%v\"", emptyBtn_def.Tag)
    }
}
