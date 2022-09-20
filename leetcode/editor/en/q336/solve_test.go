package q336

import (
    "testing"
)

func TestPalindromePairs(t *testing.T) {
    _ = palindromePairs([]string{"a", ""})
}

//leetcode submit region begin(Prohibit modification and deletion)

func palindromePairs(words []string) (ret [][]int) {
    reserved := make(map[string]int, len(words))
    for i := range words {
        reserved[reverse(words[i])] = i
    }
    if idx, found := reserved[""]; found {
        delete(reserved, "")
        for i, word := range words {
            if ii := i; ii != idx && isPalindrome(word) {
                ret = append(ret, []int{idx, ii}, []int{ii, idx})
            }
        }
    }
    
    for i, word := range words {
        for j, idx := 0, i; j < len(word); j += 1 {
            if jdx, found := reserved[word[:j]]; found && jdx != idx && isPalindrome(word[j:]) {
                ret = append(ret, []int{idx, jdx})
            }
            if jdx, found := reserved[word[j:]]; found && jdx != idx && isPalindrome(word[:j]) {
                ret = append(ret, []int{jdx, idx})
            }
        }
    }
    return ret
}

//    dict := map[string]int{}
//
//    for i, word := range words {
//        dict[reverse(word)] = i
//    }
//
//    result := map[[2]int]struct{}{}
//
//    for i, word := range words {
//        for j, max := 0, len(word); j <= max; j++ {
//            if k, ok := dict[word[:j]]; ok && i != k && isPalindrome(word[j:]) {
//                result[[2]int{i, k}] = struct{}{}
//            }
//            if k, ok := dict[word[max-j:]]; ok && i != k && isPalindrome(word[:max-j]) {
//                result[[2]int{k, i}] = struct{}{}
//            }
//        }
//    }
//
//    ans := make([][]int, 0, len(result))
//
//    for pair := range result {
//        now := pair
//        ans = append(ans, now[:])
//    }
//
//    return ans
//}

func reverse(s string) string {
    r := []byte(s)
    for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func isPalindrome(str string) bool {
    for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
        if str[i] != str[j] {
            return false
        }
    }
    return true
}

//leetcode submit region end(Prohibit modification and deletion)

func palindromePairs2(words []string) (ret [][]int) {
    reserved := make(map[string]int, len(words))
    for i := range words {
        reserved[reverse(words[i])] = i
    }
    if idx, found := reserved[""]; found {
        delete(reserved, "")
        for i, word := range words {
            if ii := i; ii != idx && isPalindrome(word) {
                ret = append(ret, []int{idx, ii}, []int{ii, idx})
            }
        }
    }
    
    for i, word := range words {
        for j, idx := 0, i; j < len(word); j += 1 {
            if jdx, found := reserved[word[:j]]; found && jdx != idx && isPalindrome(word[j:]) {
                ret = append(ret, []int{idx, jdx})
            }
            if jdx, found := reserved[word[j:]]; found && jdx != idx && isPalindrome(word[:j]) {
                ret = append(ret, []int{jdx, idx})
            }
        }
    }
    return ret
}
