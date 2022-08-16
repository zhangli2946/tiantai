package _834

func sumOfDistancesInTree(n int, edges [][]int) []int {
    return n1(n, edges)
}

func n1(n int, edges [][]int) []int {
    var (
        children = make([][]int, n)

    )
    for i := range edges {
        children[edges[i][0]] = append(children[edges[i][0]], edges[i][1])
        children[edges[i][1]] = append(children[edges[i][1]], edges[i][0])
    }

}

var (
    _ = n2
)

func n2(n int, edges [][]int) []int {
    var (
        children = make([][]int, n)
        ans      = make([]int, n)
    )
    for i := range edges {
        children[edges[i][0]] = append(children[edges[i][0]], edges[i][1])
        children[edges[i][1]] = append(children[edges[i][1]], edges[i][0])
    }
    for i := range ans {
        ans[i] = solve(children, i, -1, 0)
    }
    return ans
}

func solve(children [][]int, curr int, prev int, level int) (ret int) {
    for i := range children[curr] {
        if children[curr][i] == prev {
            continue
        }
        ret += solve(children, children[curr][i], curr, level+1)
    }
    if ret == 0 {
        return level
    }
    return ret + level
}
