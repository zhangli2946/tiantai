package _7

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 1002003001, reverse(1003002001))
    assert.Equal(t, -1002003001, reverse(-1003002001))

    assert.Equal(t, 0, reverse(MAX))
}
