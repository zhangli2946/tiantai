package solve

import (
    "testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
    _ = removeDuplicates([]int{
        0, 0,
    })
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
    if l := len(nums); l <= 1 {
        return l
    }
    i := 1
    for _, x := range nums[1:] {
        if x != nums[i-1] {
            nums[i] = x
            i++
        }
    }
    return i
}

//leetcode submit region end(Prohibit modification and deletion)
func removeDuplicates2(nums []int) int {
    var dupIdx, currIdx, end = 0, 1, len(nums)
    for ; currIdx < end; currIdx += 1 {
        for ; currIdx < end && nums[dupIdx] == nums[currIdx]; currIdx += 1 { // move on same
        }
        if currIdx == end {
            break
        }
        dupIdx, nums[dupIdx+1] = dupIdx+1, nums[currIdx]
    }
    return dupIdx + 1
}
