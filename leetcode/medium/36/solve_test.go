package _36

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    "testing"
)

type Case struct {
    board [][]byte
    valid bool
}

func TestSolve(t *testing.T) {
    for i, e := range []Case{
        {[][]byte{
            {'2', '.', '.'},
            {'.', '.', '.'},
            {'.', '.', '2'}}, false},
    } {
        t.Run(fmt.Sprintf("case_%03d", i), func(t *testing.T) {
            assert.Equal(t, isValidSudoku(e.board), e.valid)
        })
    }
}
