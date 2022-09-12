package solve

import (
    "testing"
)

func TestJumpGameIi(t *testing.T) {
    t.Log(jump([]int{2, 3, 0, 1, 4}))
    t.Log(jump([]int{2, 3, 1, 1, 4}))
    t.Log(jump([]int{2, 0, 2, 0, 1}))
    t.Log(jump([]int{0}))
    t.Log(jump([]int{1}))
    t.Log(jump([]int{1, 1, 1, 2, 1}))
    t.Log(jump([]int{1, 2, 0, 1}))
    
}

//leetcode submit region begin(Prohibit modification and deletion)

func jump(nums []int) int {
    return solve(0, nums)
}

func solve(cnt int, nums []int) int {
    
    if len(nums) < 2 {
        return cnt
    }
    for try, end, last, curr := 0, len(nums)-1, 0, 0; try < end; try += 1 {
        curr = max(curr, try+nums[try])
        if try == last {
            last = curr
            cnt += 1
            if curr > end {
                break
            }
        }
    }
    return cnt
}
func max2(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//leetcode submit region end(Prohibit modification and deletion)
