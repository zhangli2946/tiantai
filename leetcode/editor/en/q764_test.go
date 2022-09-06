package solve

import (
    "testing"
)

func TestLargestPlusSign(t *testing.T) {
    _ = orderOfLargestPlusSign
}

//leetcode submit region begin(Prohibit modification and deletion)
func orderOfLargestPlusSign(n int, mines [][]int) int {
    table := map[[2]int]struct{}{}
    for _, mine := range mines {
        table[[2]int{mine[0], mine[1]}] = struct{}{}
    }
    maxK := (n + 1) / 2
    for ; ; maxK -= 1 {
        if _, found := table[[2]int{maxK, maxK}]; found {
            continue
        }
    }
    
    return 0
}

//leetcode submit region end(Prohibit modification and deletion)
