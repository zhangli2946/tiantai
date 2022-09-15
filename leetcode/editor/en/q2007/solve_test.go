package solve

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestFindOriginalArrayFromDoubledArray(t *testing.T) {
    
    assert.Equal(t, []int{0, 0}, findOriginalArray([]int{0, 0, 0, 0}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findOriginalArray(changed []int) []int {
    feq := [1e5 + 1]int{}
    for i := range changed {
        feq[changed[i]] += 1
    }
    if feq[0]&1 == 1 {
        return []int{}
    }
    feq[0] = feq[0] / 2
    for i := range feq {
        if i == 0 || feq[i] == 0 {
            continue
        }
        i2 := 2 * i
        if i2 > 1e5 {
            return []int{}
        }
        if feq[i] > feq[i2] {
            return []int{}
        }
        feq[i2] = feq[i2] - feq[i]
    }
    ret := []int{}
    for i := range feq {
        for j := i; feq[i] > 0; feq[i] -= 1 {
            ret = append(ret, j)
        }
    }
    return ret
}

//leetcode submit region end(Prohibit modification and deletion)
