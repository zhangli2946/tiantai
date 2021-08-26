package _38

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, countAndSay(2), "11")
    assert.Equal(t, countAndSay(3), "21")
    assert.Equal(t, countAndSay(4), "1211")
    assert.Equal(t, countAndSay(5), "111221")
    assert.Equal(t, countAndSay(6), "312211")
}
