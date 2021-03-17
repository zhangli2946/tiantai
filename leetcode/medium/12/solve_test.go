package _12

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, "III", intToRoman3(3))
    assert.Equal(t, "IV", intToRoman3(4))
    assert.Equal(t, "LVIII", intToRoman3(58))
    assert.Equal(t, "MCMXCIV", intToRoman3(1994))
}

func TestLUT(t *testing.T) {
    for i := 3999; i > 0; i -= 1 {
        fmt.Printf("%d : \"%s\",\n", i, intToRoman(i))
    }
}

func BenchmarkSolve(b *testing.B) {
    for i := 0; i < b.N; i++ {
        intToRoman(1994)
    }
}
