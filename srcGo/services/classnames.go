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

func Classnames(args ...interface{}) string {
    var classList []string
    for _, arg := range args {
        className := argProcessor(arg)
        if className != "" {
            classList = append(
                classList,
                className,
            )
        }
    }
    return strings.Join(classList, " ")
}
