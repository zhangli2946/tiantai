package _18

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t,
       [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
       fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
    assert.Equal(t,
        [][]int{{0, 0, 0, 0}},
        fourSum([]int{0, 0, 0, 0}, 0))
}
