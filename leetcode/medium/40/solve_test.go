package _40

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "sort"
    "testing"
)

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    ret, path := [][]int{}, []int{}
    var backtrace func(_, _ int)
    backtrace = func(curr, idx int) {
        switch {
        case curr == target:
            copyPath := make([]int, len(path))
            copy(copyPath, path)
            ret = append(ret, copyPath)
            return
        }
        for i := idx; i < len(candidates) && curr+candidates[i] <= target; i += 1 {
            if i > idx && candidates[i] == candidates[i-1] {
                continue
            }
            path = append(path, candidates[i])
            backtrace(curr+candidates[i], i+1)
            path = path[:len(path)-1]
        }
    }
    backtrace(0, 0)
    return ret
}

type Kase struct {
    Result     [][]int
    Candidates []int
    Target     int
}

func TestSolve(t *testing.T) {
    for i, kase := range []Kase{
        {[][]int{
            {1, 1, 6},
            {1, 2, 5},
            {1, 7},
            {2, 6},
        }, []int{10, 1, 2, 7, 6, 1, 5}, 8},
        {[][]int{{1, 2, 2}, {5}}, []int{2, 5, 2, 1, 2}, 5},
    } {
        t.Run(fmt.Sprintf("case%05d", i), func(t *testing.T) {
            assert.Equal(t, kase.Result, combinationSum2(kase.Candidates, kase.Target))
        })
    }
}
