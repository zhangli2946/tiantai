package solve

import (
    "testing"
    
    "github.com/stretchr/testify/assert"
)

func TestFindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
    _ = strStr
    assert.Equal(t, 0, strStr("abc", "a"))
    assert.Equal(t, 1, strStr("abc", "bc"))
    assert.Equal(t, -1, strStr("cabc", "cbc"))
    assert.Equal(t, 1, strStr("ocabc", "cab"))
    assert.Equal(t, 0, strStr("aaa", "aaa"))
    assert.Equal(t, 1, strStr("caaa", "aaa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
    var hi, ni, hm, nm = 0, 0, len(haystack), len(needle)
    var offset = 0
    for ; hi+offset < hm && ni < nm; {
        if haystack[hi+offset] != needle[ni] {
            hi, ni, offset = hi+1, 0, 0
            continue
        }
        offset, ni = offset+1, ni+1
    }
    if offset != nm {
        return -1
    }
    return hi
}

//leetcode submit region end(Prohibit modification and deletion)
