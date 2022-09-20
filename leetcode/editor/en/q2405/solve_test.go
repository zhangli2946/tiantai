package q2405

import (
    "testing"
)

func TestOptimalPartitionOfString(t *testing.T) {
    _ = partitionString
}

//leetcode submit region begin(Prohibit modification and deletion)
func partitionString(s string) (ret int) {
    //ret = 0
    for i, max, feq := 0, len(s), [26]bool{}; i < max; i += 1 {
        if feq[s[i]-'a'] {
            ret, feq = ret+1, [26]bool{}
        }
        feq[s[i]-'a'] = true
    }
    ret += 1
    return
}

//leetcode submit region end(Prohibit modification and deletion)
