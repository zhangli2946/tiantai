package solve

import (
    "testing"
)

func TestValidParentheses(t *testing.T) {
    _ = isValid("[[]]")
    _ = isValid("[{}]")
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
    stack := []uint8{}
    for l, i := len(s), 0; i < l; i += 1 {
        switch {
        case '(' == s[i], '[' == s[i], '{' == s[i]:
            stack = append(stack, s[i])
        case len(stack) <= 0:
            return false
        case ')' == s[i] && '(' == stack[len(stack)-1],
            ']' == s[i] && '[' == stack[len(stack)-1],
            '}' == s[i] && '{' == stack[len(stack)-1]:
            stack = stack[:len(stack)-1]
        default:
            return false
        }
    }
    return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
