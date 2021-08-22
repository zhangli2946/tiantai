package _37

func solveSudoku(board [][]byte) {

    boxUsed := [9][9]bool{{false}}
    rowUsed := [9][9]bool{{false}}
    colUsed := [9][9]bool{{false}}
    scan(board, &rowUsed, &colUsed, &boxUsed)

    solve(board, rowUsed, colUsed, boxUsed, 0, 0)

}

func solve(board [][]byte, rowUsed, colUsed, boxUsed [9][9]bool, row, col int) bool {
    if col > 8 {
        row, col = row+1, 0
    }
    if row > 8 {
        return true
    }
    char := board[row][col]
    if char != '.' {
        return solve(board, rowUsed, colUsed, boxUsed, row, col+1)
    }
    for ch := byte(0); ch < 9; ch += 1 {
        // check valid choice
        if rowUsed[row][ch] ||
            colUsed[col][ch] ||
            boxUsed[row/3*3+col/3][ch] {
            continue
        }
        // do choice
        board[row][col], rowUsed[row][ch], colUsed[col][ch], boxUsed[row/3*3+col/3][ch] = ch+'1', true, true, true
        if valid := solve(board, rowUsed, colUsed, boxUsed, row, col+1); valid {
            return true
        }
        // rollback choice
        board[row][col], rowUsed[row][ch], colUsed[col][ch], boxUsed[row/3*3+col/3][ch] = '.', false, false, false
    }
    return false
}

func scan(board [][]byte, rowUsed, colUsed, boxUsed *[9][9]bool) {
    for i, r := range board {
        for j, cell := range r {
            if cell != '.' {
                rowUsed[i][cell-'1'] = true
                colUsed[j][cell-'1'] = true
                boxUsed[i/3*3+j/3][cell-'1'] = true
            }
        }
    }
}
