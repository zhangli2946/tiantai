package solve

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestMultiplyStrings(t *testing.T) {
    assert.Equal(t, "121", multiply("11", "11"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
    if len(num2) == 0 || len(num1) == 0 {
        return "0"
    }
    if num2 == "0" || num1 == "0" {
        return "0"
    }
    return ""
}

//leetcode submit region end(Prohibit modification and deletion)
