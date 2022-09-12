package q118

import (
    "testing"
)

func TestPascalsTriangle(t *testing.T) {
    t.Log(generate(3))
    t.Log(generate(5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
    ret := make([][]int, 0, numRows)
    ret = append(ret, []int{1})
    for i := 1; i < numRows; i += 1 {
        row := make([]int, i+1, i+1)
        row[0], row[i] = 1, 1
        for d := 1; 2*d < (i + 1); d += 1 {
            row[d], row[i-d] = ret[i-1][d-1]+ret[i-1][d], ret[i-1][i-d]+ret[i-1][i-d-1]
        }
        ret = append(ret, row)
    }
    return ret
}

//leetcode submit region end(Prohibit modification and deletion)
