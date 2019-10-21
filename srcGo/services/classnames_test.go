package services

import (
    "sort"
    "strings"
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
    // map[] is not persisting order of the keys.
    // Therefore I'll need to sort it in order to test be consistent in results
    mapOfArgs_list := strings.Split(mapOfArgs, " ")
    sort.Strings(mapOfArgs_list)

    mapOfArgs_expected := "3 first"
    if strings.Join(mapOfArgs_list, " ") != mapOfArgs_expected {
        t.Errorf("Classnames() failed, expected \"%v\", got \"%v\"", mapOfArgs_expected, mapOfArgs)
    }
}
