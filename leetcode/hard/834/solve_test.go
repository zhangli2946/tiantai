package _834

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, sumOfDistancesInTree(3, [][]int{{0, 1}, {0, 2}}), []int{2, 3, 3})
    assert.Equal(t, sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {3, 2}, {4, 2}, {5, 2}}), []int{8,12,6,10,10,10})
}


