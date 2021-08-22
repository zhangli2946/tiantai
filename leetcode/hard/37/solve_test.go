package _37

import (
    "fmt"
    "github.com/go-playground/assert/v2"
    _36 "github.com/zhangli2946/tiantai/leetcode/medium/36"
    "testing"
)

type Case struct {
    board [][]byte
}

func TestSolve(t *testing.T) {
    for i, e := range []Case{
        {[][]byte{
            {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
            {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
            {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
            {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
            {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
            {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
            {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
            {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
            {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
        }},
    } {
        t.Run(fmt.Sprintf("case_%03d", i), func(t *testing.T) {

            assert.Equal(t, true, _36.IsValidSudoku(e.board))
            show(e, '0')
            solveSudoku(e.board)
            show(e, '0')
            assert.Equal(t, true, _36.IsValidSudoku(e.board))
        })
    }
}

func show(e Case, flag byte) {
    fmt.Printf("---------------------------------------------\n")
    for _, row := range e.board {
        for _, cell := range row {
            if p := cell - flag; p < 10 {
                fmt.Printf("% 4d ", p)
                continue
            }
            fmt.Printf("%s ", "   .")
        }
        fmt.Printf("\n")
    }
}
