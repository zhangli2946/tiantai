package q119

import (
    "testing"
)

func TestPascalsTriangleIi(t *testing.T) {
    t.Log(getRow(0))
    t.Log(getRow(1))
    t.Log(getRow(2))
    t.Log(getRow(3))
    t.Log(getRow(4))
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
    ret := []int{1}
    for i := 1; i < rowIndex+1; i += 1 {
        ret = append(ret, 1)
        for j, k := 1, len(ret)-2; j <= k; j, k = j+1, k-1 {
            ret[j], ret[k] = ret[k]+ret[j], ret[k]+ret[j]
        }
    }
    return ret
}

//leetcode submit region end(Prohibit modification and deletion)
