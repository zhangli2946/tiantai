package _1235

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}), 120)
    assert.Equal(t, jobScheduling([]int{6, 15, 7, 11, 1, 3, 16, 2}, []int{19, 18, 19, 16, 10, 8, 19, 8}, []int{2, 9, 1, 19, 5, 7, 3, 19}), 41)
    assert.Equal(t, jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}), 150)
    assert.Equal(t, jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}), 6)
}
