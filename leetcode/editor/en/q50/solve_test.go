package solve

import (
    "testing"
)

func TestPowxN(t *testing.T) { _ = myPow }

//leetcode submit region begin(Prohibit modification and deletion)

func myPow(x float64, n int) float64 {
    if flag := n < 0; flag {
        return 1 / myPowCached(x, -n)
    }
    return myPowCached(x, n)
}
func myPowCached(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    if n&1 == 0 {
        ans := myPowCached(x, n/2)
        return ans * ans
    }
    ans := myPowCached(x, (n-1)/2)
    return x * ans * ans
}

//leetcode submit region end(Prohibit modification and deletion)
