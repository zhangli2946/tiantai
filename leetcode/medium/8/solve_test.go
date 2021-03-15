package _8

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 1002002001, myAtoi("1002002001"))
    assert.Equal(t, -12, myAtoi("  -0012a42"))
    assert.Equal(t, -1002003001, myAtoi("-1002003001"))
    assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
    assert.Equal(t, -2147483648, myAtoi("-9223372036854775808"))
    assert.Equal(t, 2147483647, myAtoi("18446744073709551617"))
}
