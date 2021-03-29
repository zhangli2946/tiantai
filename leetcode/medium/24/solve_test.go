package _24

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    assert.Equal(t, true, compare(build(0, 1), swapPairs(build(1, 0))))
    assert.Equal(t, true, compare(build(0, 1, 2), swapPairs(build(1, 0, 2))))
    assert.Equal(t, true, compare(build(0, 1, 2, 3), swapPairs(build(1, 0, 3, 2))))
}
func build(nums ...int) (ret *ListNode) {
    ret = &ListNode{}
    var (
        curr = ret
    )
    for i := range nums {
        curr.Next = &ListNode{Val: nums[i]}
        curr = curr.Next
    }
    return ret.Next
}
func compare(ret *ListNode, end *ListNode) bool {
    switch {
    case ret == nil && end == nil:
        return true
    case ret == nil && end != nil, ret != nil && end == nil:
        return false
    default:
        if ret.Val != end.Val {
            return false
        }
        return compare(ret.Next, end.Next)
    }
}
