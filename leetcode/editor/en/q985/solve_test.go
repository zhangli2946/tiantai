package solve

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestSumOfEvenNumbersAfterQueries(t *testing.T) {
    assert.Equal(t, []int{8, 6, 2, 4}, sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{
        {1, 0}, {-3, 1}, {-4, 0}, {2, 3},
    }))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    ans := make([]int, len(queries)+1)
    for i := range nums {
        if v := nums[i]; (v >= 0 && v&1 == 0) || (v < 0 && (v*-1)&1 == 0) {
            ans[0] += v
        }
    }
    for i := range queries {
        val, idx := queries[i][0], queries[i][1]
        ans[i+1] = ans[i]
        if v := nums[idx]; (v >= 0 && v&1 == 0) || (v < 0 && (v*-1)&1 == 0) {
            ans[i+1] -= nums[idx]
        }
        nums[idx] += val
        if v := nums[idx]; (v >= 0 && v&1 == 0) || (v < 0 && (v*-1)&1 == 0) {
            ans[i+1] += v
        }
    }
    return ans[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
