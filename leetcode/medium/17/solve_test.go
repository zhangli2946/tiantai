package _17

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 0, len(letterCombinations("")))
    assert.Equal(t, 3, len(letterCombinations("2")))
    assert.Equal(t, 9, len(letterCombinations("23")))
    assert.Equal(t, 27, len(letterCombinations("234")))
    assert.Equal(t, 81, len(letterCombinations("2345")))
}
