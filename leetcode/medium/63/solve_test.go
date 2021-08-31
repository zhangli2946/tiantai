package _63

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "testing"
)

type Kase struct {
    Result       int
    ObstacleGrid [][]int
}

func TestSolve(t *testing.T) {
    for i, kase := range []Kase{
        {1, [][]int{{0, 1}, {0, 0}}},
        {2, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
    } {
        t.Run(fmt.Sprintf("%05d", i), func(t *testing.T) {
            assert.Equal(t, kase.Result, uniquePathsWithObstacles(kase.ObstacleGrid))
        })
    }
}
