package _11

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
    assert.Equal(t, 49, maxArea([]int{2, 8, 6, 2, 5, 4, 8, 3, 7}))
    assert.Equal(t, 49, maxArea([]int{5, 8, 6, 2, 5, 4, 8, 3, 7}))
    assert.Equal(t, 49, maxArea([]int{7, 8, 6, 2, 5, 4, 98, 99, 7}))
}
