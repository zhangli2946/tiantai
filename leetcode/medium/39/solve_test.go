package _39

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "sort"
    "testing"
)

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    ret := [][]int{}
    for i := len(candidates) - 1; i >= 0; i -= 1 {
        solve(candidates, target-candidates[i], i, []int{candidates[i]}, &ret)
    }
    return ret
}

func solve(candidates []int, target int, upper int, acc []int, ret *[][]int) {
    // fast path
    switch {
    case target < 0:
        return
    case 0 == target:
        *ret = append(*ret, append([]int{}, acc...))
        return
    }
    for i := upper; i >= 0; i -= 1 {
        solve(candidates, target-candidates[i], i, append(acc, candidates[i]), ret)
    }
}

type Kase struct {
    Result     [][]int
    Candidates []int
    Target     int
}

func TestSolve(t *testing.T) {
    for i, kase := range []Kase{
        //{[][]int{{1, 1}}, []int{1}, 2},
        //{[][]int{}, []int{2}, 1},
        //{[][]int{{5, 3}, {3, 3, 2}, {2, 2, 2, 2}}, []int{2, 3, 5}, 8},
        {[][]int{{7, 2},
            {7, 1, 1},
            {6, 3},
            {6, 2, 1},
            {6, 1, 1, 1},
            {5, 3, 1},
            {5, 2, 2},
            {5, 2, 1, 1},
            {5, 1, 1, 1, 1},
            {3, 3, 3},
            {3, 3, 2, 1},
            {3, 3, 1, 1, 1},
            {3, 2, 2, 2},
            {3, 2, 2, 1, 1},
            {3, 2, 1, 1, 1, 1},
            {3, 1, 1, 1, 1, 1, 1},
            {2, 2, 2, 2, 1},
            {2, 2, 2, 1, 1, 1},
            {2, 2, 1, 1, 1, 1, 1},
            {2, 1, 1, 1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1, 1, 1, 1, 1}}, []int{2, 7, 6, 3, 5, 1}, 9},
    } {
        t.Run(fmt.Sprintf("case%05d", i), func(t *testing.T) {
            assert.Equal(t, kase.Result, combinationSum(kase.Candidates, kase.Target))
        })
    }
}
