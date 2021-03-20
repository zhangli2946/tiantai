package _16

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
    assert.Equal(t, 0, threeSumClosest([]int{0, 2, 1, -3}, 1))
}
