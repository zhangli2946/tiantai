package _3

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 0, lengthOfLongestSubstring(""))
    assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
    assert.Equal(t, 4, lengthOfLongestSubstring("abcda"))
    assert.Equal(t, 6, lengthOfLongestSubstring("abcabcdef"))
    assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
