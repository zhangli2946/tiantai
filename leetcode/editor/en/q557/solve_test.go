package solve

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestReverseWordsInAStringIii(t *testing.T) {
    
    assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
    ss := []byte(s)
    for i, j := 0, 0; i <= len(ss); i += 1 {
        if i == len(ss) || ss[i] == ' ' {
            for ii, ij := i-1, j; ij < ii; ii, ij = ii-1, ij+1 {
                ss[ii], ss[ij] = ss[ij], ss[ii]
            }
            j = i + 1
        }
    }
    return string(ss)
}

//leetcode submit region end(Prohibit modification and deletion)
