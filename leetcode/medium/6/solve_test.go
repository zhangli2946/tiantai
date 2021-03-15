package _6

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, "", convert("", 1))
    assert.Equal(t, "ab", convert("ab", 1))
    assert.Equal(t, "aaaaaaabbbbbbb", convert("ababababababab", 2))
    assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
}
