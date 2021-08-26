package _38

import (
    "strconv"
    "strings"
)

func countAndSay(n int) string {
    var newState = strings.Builder{}
    newState.WriteRune('1')
    for i := 1; i < n; i += 1 {
        // reset state recorder
        state := newState.String()
        currChar, count := int32(state[0]), 0
        newState.Reset()
        for _, char := range state {
            if char == currChar {
                count += 1
                continue
            }
            newState.WriteString(strconv.Itoa(count))
            newState.WriteRune(currChar)
            currChar, count = char, 1
        }
        newState.WriteString(strconv.Itoa(count))
        newState.WriteRune(currChar)
    }
    return newState.String()
}
