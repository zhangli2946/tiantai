package q834

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestSumOfDistancesInTree(t *testing.T) {
    assert.Equal(t, []int{1, 1}, sumOfDistancesInTree(2, [][]int{
        {0, 1},
    }))
    assert.Equal(t, []int{2, 3, 3}, sumOfDistancesInTree(3, [][]int{
        {0, 1}, {0, 2},
    }))
    assert.Equal(t, []int{8, 12, 6, 10, 10, 10}, sumOfDistancesInTree(6, [][]int{
        {0, 1}, {0, 2},
        {2, 3}, {2, 4}, {2, 5},
    }))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumOfDistancesInTree(n int, edges [][]int) (ans []int) {
    relation := make([][]int, n, n)
    for i := 0; i < n; i += 1 {
        relation[i] = make([]int, n, n)
    }
    for i, edge := range edges {
        relation[i][i], relation[edge[0]][edge[1]], relation[edge[1]][edge[0]] = 0, 1, 1
    }
    
    for from := n - 1; from > -1; from -= 1 {
        for to := n - 1; to > -1; to -= 1 {
            if from == to  || relation[from][to] != 0 { continue }
            // deep first
            stack := make([]int, n)
            var i, v  = 0
            for stack = append(stack, to); len(stack) > 0; {
                pop := stack[len(stack)-1]
                for i, v = range relation[pop] {
                    if i  == to { continue }
                    if v != 0 {  break }
                }
                relation[][] = v +1
            }
        }
    }
    ans = make([]int, n, n)
    for i := range ans {
        for j := 0; j < n; j += 1 {
            if i == j {
                continue
            }
            ans[i] += relation[i][j]
        }
    }
    return
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

//leetcode submit region end(Prohibit modification and deletion)
