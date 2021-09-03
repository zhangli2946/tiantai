package _565

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "testing"
)

type Kase struct {
    Array  []int
    Result int
}

func TestSolve(t *testing.T) {
    for i, kase := range []Kase{
        {[]int{5, 4, 0, 3, 1, 6, 2}, 4},
        {[]int{0,1,2}, 1},
    } {
        t.Run(fmt.Sprintf("case%05d", i), func(t *testing.T) {
            assert.Equal(t, arrayNesting(kase.Array), kase.Result)
        })
    }
}
