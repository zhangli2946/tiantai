package _34

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    for i, e := range [][2][]int{
        {{}, {0, -1, -1}},
        {{1}, {0, -1, -1}},
        {{1}, {1, 0, 0}},
        {{2, 2}, {2, 0, 1}},
        {{1, 4}, {4, 1, 1}},
        {{1, 4}, {1, 0, 0}},
        {{1, 4}, {2, -1, -1}},
        {{1, 2, 2}, {2, 1, 2}},
        {{2, 2, 3}, {2, 0, 1}},
        {{1, 2, 2, 3}, {2, 1, 2}},
        {{1, 2, 4, 5}, {3, -1, -1}},
        {{2, 2, 3}, {3, 2, 2}},
    } {
        t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
            assert.Equal(t, searchRange(e[0], e[1][0]), e[1][1:])
        })
    }
}
