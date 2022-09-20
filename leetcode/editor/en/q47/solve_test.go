package q47

import (
    "sort"
    "testing"
)

func TestPermutationsIi(t *testing.T) {
    _ = permuteUnique
}

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }
    used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
    sort.Ints(nums)
    generatePermutation47(nums, 0, p, &res, &used)
    return res
}
func generatePermutation47(nums []int, index int, p []int, res *[][]int, used *[]bool) {
    if index == len(nums) {
        temp := make([]int, len(p))
        copy(temp, p)
        *res = append(*res, temp)
        return
    }
    for i := 0; i < len(nums); i++ {
        if !(*used)[i] {
            if i > 0 && nums[i-1] == nums[i] && !(*used)[i-1] {
                continue
            }
            (*used)[i] = true
            p = append(p, nums[i])
            generatePermutation47(nums, index+1, p, res, used)
            p = p[:len(p)-1]
            (*used)[i] = false
        }
    }
    return
}

//leetcode submit region end(Prohibit modification and deletion)
