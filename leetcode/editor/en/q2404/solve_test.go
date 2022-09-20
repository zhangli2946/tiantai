package q2404

import (
    "testing"
)

func TestMostFrequentEvenElement(t *testing.T) {
    _ = mostFrequentEven
}

//leetcode submit region begin(Prohibit modification and deletion)
func mostFrequentEven(nums []int) int {
    feq := [50000]int{}
    for i := range nums {
        if nums[i]&1 == 0 {
            feq[nums[i]/2] += 1
        }
    }
    factor, max := -1, 0
    for i := 0; i < 50000; i += 1 {
        if max < feq[i] {
            factor, max = i, feq[i]
        }
    }
    if factor != -1 {
        return factor * 2
    }
    return factor
}

//leetcode submit region end(Prohibit modification and deletion)
