package solve

import (
    "testing"
)

func TestPalindromeNumber(t *testing.T) { _ = isPalindrome }

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    var factors = []int{}
    for ; x > 0; x = x / 10 {
        factors = append(factors, x%10)
    }
    for i, j := 0, len(factors)-1; i < j; i, j = i+1, j-1 {
        if factors[i] != factors[j] {
            return false
        }
    }
    return true
}

//leetcode submit region end(Prohibit modification and deletion)
