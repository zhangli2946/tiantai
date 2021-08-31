package _62

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "testing"
)

type Kase struct {
    Result int
    r, c   int
}

func TestSolve(t *testing.T) {
    for i, kase := range []Kase{
        {28, 3, 7},
        {6, 3, 3},
    } {
        t.Run(fmt.Sprintf("%05d", i), func(t *testing.T) {
            assert.Equal(t, kase.Result, uniquePathsWithObstacles(kase.r, kase.c))
        })
    }
}
