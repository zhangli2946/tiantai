package _15

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, [][]int{}, threeSum([]int{}))
    assert.Equal(t, [][]int{}, threeSum([]int{0}))
    assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
    assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}, {0, 0, 0}}, threeSum([]int{-1, 0, 0, 0, 1, 2, -1, -4}))
}

func BenchmarkSolve(b *testing.B) {
    for i := 0; i < b.N; i++ {
        threeSum([]int{-1, 0, 0, 0, 1, 2, -1, -4})
    }
}