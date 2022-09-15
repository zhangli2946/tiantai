package solve

import (
    "testing"
)

func TestDivideIntervalsIntoMinimumNumberOfGroups(t *testing.T) {
    _ = minGroups
}

//leetcode submit region begin(Prohibit modification and deletion)

func minGroups(intervals [][]int) (ret int) {
    m := make([]int, 1e6+10)
    for i := range intervals {
        m[intervals[i][0]], m[intervals[i][1]+1] = m[intervals[i][0]]+1, m[intervals[i][1]+1]-1
    }
    var ans = 0
    for i := range m {
        if ans += m[i]; ret < ans {
            ret = ans
        }
    }
    return ret
}

//leetcode submit region end(Prohibit modification and deletion)
