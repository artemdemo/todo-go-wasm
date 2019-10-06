package services

import (
    "fmt"
    "strings"
)

func argProcessor(arg interface{}) string {
    if argMap, ok := arg.(map[string]bool); ok {
        var classList []string
        for className, classAllowed := range argMap {
            if classAllowed {
                classList = append(
                    classList,
                    strings.Trim(className, " "),
                )
            }
        }
        return strings.Join(classList, " ")
    }
    return fmt.Sprintf("%v", arg)
}

// "How to check variable type is map in Go language"
// https://stackoverflow.com/a/20759949
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
