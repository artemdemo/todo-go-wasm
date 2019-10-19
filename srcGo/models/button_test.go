package models

import (
    _ "fmt"
    _ "syscall/js"
    "testing"

    _ "../htmlrender"
    _ "../services"
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

    if emptyBtn_def.Tag != "" {
        t.Errorf("GetElementDef() failed, expected an empty string, got \"%v\"", emptyBtn_def.Tag)
    }
}
