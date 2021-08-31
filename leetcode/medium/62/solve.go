package _62

func uniquePathsWithObstacles(row, col int) int {

    mase := make([][]int, row)
    for i := range mase {
        mase[i] = make([]int, col)
    }
    // mark initial state visit start point only 1 way
    mase[0][0] = 1
    for i := range mase {
        for j := range mase[i] {
            if j != 0 {
                mase[i][j] += mase[i][j-1]
            }
            if i != 0 {
                mase[i][j] += mase[i-1][j]
            }
        }
    }
    // return the final state of how many ways
    return mase[row-1][col-1]
}
