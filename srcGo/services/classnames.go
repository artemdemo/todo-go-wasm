package services

import (
    "fmt"
    "sort"
    "strings"
)

func argProcessor(arg interface{}) string {
    if argMap, ok := arg.(map[string]bool); ok {
        // Order of the keys is not constant in map.
        // It will break tests, therefore I'm sorting it.
        keys := make([]string, 0, len(argMap))
        for itemKey := range argMap {
            keys = append(keys, itemKey)
        }
        sort.Strings(keys)
        var classList []string
        for _, keyItem := range keys {
            classAllowed := argMap[keyItem]
            if classAllowed {
                classList = append(
                    classList,
                    strings.Trim(keyItem, " "),
                )
            }
        }
        return strings.Join(classList, " ")
    }
    return fmt.Sprintf("%v", arg)
}

func Classnames(args ...interface{}) string {
    var classList []string
    for _, arg := range args {
        classList = append(
            classList,
            argProcessor(arg),
        )
    }
    return strings.Join(classList, " ")
}
