package q70

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestClimbingStairs(t *testing.T) {
    assert.Equal(t, 3, climbStairs(3))
    assert.Equal(t, 5, climbStairs(4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    d2 := 1
    for i, d0, d1 := 0, 1, 1; i < n-1; i += 1 {
        d2 = d0 + d1
        d0, d1 = d1, d2
    }
    return d2
}

//leetcode submit region end(Prohibit modification and deletion)
