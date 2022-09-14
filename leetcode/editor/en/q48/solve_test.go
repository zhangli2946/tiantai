package solve

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestRotateImage(t *testing.T) {
    array := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    rotate(array)
    assert.Equal(t, [][]int{
        {7, 4, 1},
        {8, 5, 2},
        {9, 6, 3},
    }, array)
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
    l := len(matrix)
    for i := 0; i < l/2; i += 1 {
        for j := i; j < l-1-i; j += 1 {
            matrix[j][l-1-i], matrix[l-1-i][l-1-j], matrix[l-1-j][i], matrix[i][j] =
                matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j], matrix[l-1-j][i]
        }
    }
}

//leetcode submit region end(Prohibit modification and deletion)
