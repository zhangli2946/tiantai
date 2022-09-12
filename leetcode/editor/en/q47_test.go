package solve

import (
    "testing"
)

func TestPermutationsIi(t *testing.T) {
    _ = permuteUnique
}

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
    if len(nums) == 1 || len(nums) == 2 && nums[0] == nums[1] {
        return [][]int{nums}
    }
    if len(nums) == 2 {
        return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
    }
    //sort.sor(nums)
    factor, ret := nums[0], permute(nums[1:])
    lor := len(ret)
    for i := 0; i < lor; i += 1 {
        for start, loc := 0, len(ret[i]); start < loc+1; start += 1 {
            row := make([]int, loc+1, loc+1)
            row[start] = factor
            for j := start; j-start < loc; j += 1 {
                row[(j+1)%(loc+1)] = ret[i][j-start]
            }
            if row[start] == row[(start+1)%(loc+1)] {
                start += 1
            }
            ret = append(ret, row)
        }
    }
    return ret[lor:]
}

//leetcode submit region end(Prohibit modification and deletion)
