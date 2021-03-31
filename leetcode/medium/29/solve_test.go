package _29

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, 3, divide(10, 3))
    assert.Equal(t, -2, divide(7, -3))
    assert.Equal(t, -2, divide(-7, 3))
    assert.Equal(t, 1, divide(1, 1))
    assert.Equal(t, 2147483647, divide(-2147483648,-1))
    assert.Equal(t, -2147483648, divide(-2147483648,1))
}

func TestCommon(t *testing.T) {
    s := int32(-100)
    p := int32(10)
    s >>= 1 ;t.Log(s/p)
    s >>= 1 ;t.Log(s/p)
    s >>= 1 ;t.Log(s/p)
    s >>= 1 ;t.Log(s/p)

}
