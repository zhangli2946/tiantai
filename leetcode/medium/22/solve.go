package _22

func generateParenthesis(n int) []string {
    var (
        ret = []string{}
    )
    dfs(n, n, "", &ret)
    return ret
}

func dfs(l, r int, s string, ret *[]string) {
    if l+r == 0 {
        *ret = append(*ret, s)
        return
    }
    if l > 0 {
        dfs(l-1, r, s+"(", ret)
    }
    if r > l {
        dfs(l, r-1, s+")", ret)
    }
}
