package _36

var IsValidSudoku = isValidSudoku

func isValidSudoku(board [][]byte) bool {
    for i, row := range board {
        for j, val := range row {
            switch {
            case val == '.':
                continue
            case !validSingleRow(i, board):
                return false
            case !validSingleCol(j, board):
                return false
            case !validSquire(i, j, board):
                return false
            }
        }
    }
    return true
}

func validSquire(i int, j int, board [][]byte) bool {
    startCol, startRow := (j/3)*3, (i/3)*3
    table := [10]int{0}
    for i := startRow; i < startRow+3; i += 1 {
        for j := startCol; j < startCol+3; j += 1 {
            val := board[i][j]
            switch {
            case val == '.':
                continue
            default:
                if table[val-'0'] == 0 {
                    table[val-'0'] += 1
                } else {
                    return false
                }
            }
        }
    }
    return true
}

func validSingleCol(col int, board [][]byte) bool {
    table := [10]int{0}
    for i := len(board) - 1; i >= 0; i -= 1 {
        val := board[i][col]
        switch {
        case val == '.':
            continue
        default:
            if table[val-'0'] == 0 {
                table[val-'0'] += 1
            } else {
                return false
            }
        }
    }
    return true
}
func validSingleRow(rowIdx int, board [][]byte) bool {
    table := [10]int{0}
    for _, val := range board[rowIdx] {
        switch {
        case val == '.':
            continue
        default:
            if table[val-'0'] == 0 {
                table[val-'0'] += 1
            } else {
                return false
            }
        }
    }
    return true
}
