package services

import (
    "testing"
)

func Test_Classnames(t *testing.T) {
    // 1 test
    emptyResult := Classnames()

    if emptyResult != "" {
        t.Errorf("Classnames() failed, expected an empty string, got \"%v\"", emptyResult)
    }

    // 2 test
    listOfArgs := Classnames(
        "first",
        "second",
        "3",
    )

    listOfArgs_expected := "first second 3"
    if listOfArgs != listOfArgs_expected {
        t.Errorf("Classnames() failed, expected \"%v\", got \"%v\"", listOfArgs_expected, listOfArgs)
    }

    // 3 test
    mapOfArgs := Classnames(map[string]bool{
        "first": true,
        "second": false,
        "3": true,
    })

    mapOfArgs_expected := "3 first"
    if mapOfArgs != mapOfArgs_expected {
        t.Errorf("Classnames() failed, expected \"%v\", got \"%v\"", mapOfArgs_expected, mapOfArgs)
    }
}
