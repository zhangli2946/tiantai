package q948

import (
    "sort"
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestBagOfTokens(t *testing.T) {
    assert.Equal(t, 1, bagOfTokensScore([]int{200, 100}, 150))
    assert.Equal(t, 2, bagOfTokensScore([]int{200, 100,300,400}, 200))
}

//leetcode submit region begin(Prohibit modification and deletion)
func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    res := 0
    for start, end, score := 0, len(tokens)-1, 0; start <= end; {
        if power >= tokens[start] {
            power, score = power-tokens[start], score+1
            res, start = max(res, score), start+1
        } else if score > 0 {
            power, score = power+tokens[end], score-1
            end -= 1
        } else {
            break
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//leetcode submit region end(Prohibit modification and deletion)
