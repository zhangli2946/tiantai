package solve

import (
    "testing"
    
    "github.com/stretchr/testify/assert"
)

func TestSatisfiabilityOfEqualityEquations(t *testing.T) {
    assert.Equal(t, false, equationsPossible([]string{"a==b", "a!=b"}))
    assert.Equal(t, true, equationsPossible([]string{"a==b", "a!=c"}))
    assert.Equal(t, false, equationsPossible([]string{"a==b", "b!=a"}))
    assert.Equal(t, false, equationsPossible([]string{"a==b", "e==c", "b==c", "a!=e"}))
    assert.Equal(t, false, equationsPossible([]string{"a!=e", "e==c", "b==c", "a==b"}))
    assert.Equal(t, false, equationsPossible([]string{"f==a", "a==b", "f!=e", "a==c", "b==e", "c==f"}))
}

//leetcode submit region begin(Prohibit modification and deletion)

func equationsPossible(equations []string) bool {
    const eq, ne = '=', '!'
    array := [26]uint32{}
    for i := range equations {
        a, flag, b := equations[i][0]-'a', equations[i][1], equations[i][3]-'a'
        if flag == eq {
            merged := array[a] | array[b] | 1<<b | 1<<a
            for j, idx := merged, 0; j > 0; j, idx = j>>1, idx+1 {
                if j&1 == 1 {
                    array[idx] |= merged
                }
            }
        }
    }
    for i := range equations {
        if a, flag, b := equations[i][0]-'a', equations[i][1], equations[i][3]-'a'; flag == ne {
            if a == b {
                return false
            }
            if array[a]&(1<<b) > 0 {
                return false
            }
            if array[b]&(1<<a) > 0 {
                return false
            }
        }
    }
    return true
}

//leetcode submit region end(Prohibit modification and deletion)
