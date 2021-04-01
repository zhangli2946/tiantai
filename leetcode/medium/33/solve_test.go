package _33

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 4, search([]int{4,5,6,7,0,1,2}, 4))
    assert.Equal(t, 3, search([]int{4,5,6,7,0,1,2}, 7))
    assert.Equal(t, -1, search([]int{4,5,6,7,0,1,2}, 8))
}
