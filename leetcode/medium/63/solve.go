package _63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    row := len(obstacleGrid)
    col := len(obstacleGrid[0])
    mase := make([][]int, row)
    for i := range mase {
        mase[i] = make([]int, col)
    }
    // mark initial state visit start point only 1 way
    mase[0][0] = 1
    for i := range obstacleGrid {
        for j := range obstacleGrid[i] {
            if 0 == obstacleGrid[i][j] {
                if j != 0 {
                    mase[i][j] += mase[i][j-1]
                }
                if i != 0 {
                    mase[i][j] += mase[i-1][j]
                }
            } else {
                mase[i][j] = 0
            }
        }
    }
    // return the final state of how many ways
    return mase[row-1][col-1]
}
