package q188

import (
    "math"
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestBestTimeToBuyAndSellStockIv(t *testing.T) {
    assert.Equal(t, 2, maxProfit(2, []int{2, 4, 1}))
    assert.Equal(t, 2, maxProfit(2, []int{1, 0, 2}))
    assert.Equal(t, 8, maxProfit(2, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
    assert.Equal(t, 2, maxProfit(2, []int{1, 1, 3}))
    assert.Equal(t, 1, maxProfit(1, []int{1, 2}))
    assert.Equal(t, 3, maxProfit(2, []int{1, 2, 4}))
    assert.Equal(t, 7, maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
    _ = maxProfit
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) (profit int) {
    lop := len(prices)
    if lop == 0 {
        return 0
    }
    if k > lop/2 {
        for i, p := range prices[1:] {
            profit += max(0, p-prices[i])
        }
        return
    }
    dp := make([][]int, k+1, k+1)
    for i := range dp {
        dp[i] = make([]int, lop, lop)
    }
    for t := 1; t < k+1; t += 1 {
        profit = math.MinInt32
        for d := 1; d < lop; d += 1 {
            profit = max(profit, dp[t-1][d-1]-prices[d-1])
            dp[t][d] = max(dp[t][d-1], profit+prices[d])
        }
    }
    return dp[k][lop-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//leetcode submit region end(Prohibit modification and deletion)
