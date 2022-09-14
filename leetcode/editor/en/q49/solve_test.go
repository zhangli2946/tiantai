package solve

import (
    "testing"
)

func TestGroupAnagrams(t *testing.T) { _ = groupAnagrams }

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
    bag := map[[26]uint8][]string{}
    for i := 0; i < len(strs); i += 1 {
        array := [26]uint8{}
        for _, c := range strs[i] {
            array[c-'a'] += 1
        }
        bag[array] = append(bag[array], strs[i])
    }
    ret := make([][]string, 0)
    for k := range bag {
        ret = append(ret, bag[k])
    }
    return ret
}

//leetcode submit region end(Prohibit modification and deletion)
