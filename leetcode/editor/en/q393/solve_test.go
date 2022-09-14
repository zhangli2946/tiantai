package q393

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestUtf8Validation(t *testing.T) {
    assert.Equal(t, true, validUtf8([]int{197, 130, 1}))
    assert.Equal(t, false, validUtf8([]int{235, 140, 1}))
    assert.Equal(t, false, validUtf8([]int{237}))
}

//leetcode submit region begin(Prohibit modification and deletion)
const (
    State0 = 0
    State1 = 1
    State2 = 2
    State3 = 3
)

func validUtf8(s []int) bool {
    for i, max := 0, len(s); i < max; {
        switch {
        case 0b0000_0000 == s[i]&0b1000_0000 && max-i >= 1:
            i += 1
        case 0b1111_0000 == s[i]&0b1111_1000 &&
            max-i >= 4 && 0b1000_0000 == s[i+1]&0b1100_0000 &&
            0b1000_0000 == s[i+2]&0b1100_0000 && 0b1000_0000 == s[i+3]&0b1100_0000:
            i += 4
        case 0b1110_0000 == s[i]&0b1111_0000 &&
            max-i >= 3 && 0b1000_0000 == s[i+1]&0b1100_0000 &&
            0b1000_0000 == s[i+2]&0b1100_0000:
            i += 3
        case 0b1100_0000 == s[i]&0b1110_0000 &&
            max-i >= 2 && 0b1000_0000 == s[i+1]&0b1100_0000:
            i += 2
        default:
            return false
        }
    }
    return true
}

//leetcode submit region end(Prohibit modification and deletion)
