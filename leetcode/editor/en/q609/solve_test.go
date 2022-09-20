package q609

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
)

func TestFindDuplicateFileInSystem(t *testing.T) {
    assert.Equal(t, [][2]string{
        {"abcd", "root/a/1.txt"},
        {"efgh", "root/a/2.txt"},
    }, digest("root/a 1.txt(abcd) 2.txt(efgh)"))
    _ = findDuplicate
}

//leetcode submit region begin(Prohibit modification and deletion)
func findDuplicate(paths []string) [][]string {
    ret := map[string][]string{}
    
    for _, path := range paths {
        for _, pair := range digest(path) {
            if _, found := ret[pair[0]]; !found {
                ret[pair[0]] = []string{}
            }
            ret[pair[0]] = append(ret[pair[0]], pair[1])
        }
    }
    res := [][]string{}
    for k := range ret {
        if len(ret[k]) > 1 {
            res = append(res, ret[k])
        }
    }
    return res
}

func digest(path string) (ret [][2]string) {
    var dir string
    var s = [2]string{}
    for i, state, d := 0, 0, 0; i < len(path); i += 1 {
        if 0 == state && path[i] == ' ' {
            state, dir, d = 1, path[:i]+"/", i+1
        }
        if 3 == state && path[i] == ' ' {
            s = [2]string{s[1]}
            state, s[1], d = 1, dir+path[d:i], i+1
        }
        if 1 == state && path[i] == '(' {
            state, s[1], d = 2, dir+path[d:i], i+1
        }
        if 2 == state && path[i] == ')' {
            state, s[0], d = 3, path[d:i], 0
            ret = append(ret, s)
        }
    }
    return
}

//leetcode submit region end(Prohibit modification and deletion)
