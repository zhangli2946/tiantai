package _5

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 0, len(longestPalindrome("")))
    assert.Equal(t, 1, len(longestPalindrome("a")))
    assert.Equal(t, 2, len(longestPalindrome("aa")))
    assert.Equal(t, 2, len(longestPalindrome("cbbd")))
    assert.Equal(t, 5, len(longestPalindrome("abcbac")))
    assert.Equal(t, 7, len(longestPalindrome("cabcbac")))
    assert.Equal(t, 10, len(longestPalindrome("abcbacccccccccc")))
    assert.Equal(t, 12, len(longestPalindrome("cabcbaabcbac")))
}
