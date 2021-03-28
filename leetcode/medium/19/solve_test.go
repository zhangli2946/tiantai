package _19

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestGetPrev(t *testing.T) {
    assert.Equal(t, true,
        compareNode(removeNthFromEnd(build(0, 1, 2, 3, 4, 5, 6, 7), 1), build(6)),
    )
    assert.Equal(t, true,
        compareNode(removeNthFromEnd(build(0, 1, 2, 3, 4, 5, 6, 7), 7), build(0)),
    )
}

func TestSolve(t *testing.T) {
    assert.Equal(
        t, true, compare(build(0, 1, 2, 3, 4, 5, 6),
            removeNthFromEnd(build(0, 1, 2, 3, 4, 5, 6, 7), 1)))
    assert.Equal(
        t, true, compare(build(0, 1, 2, 3, 4, 5, 7),
            removeNthFromEnd(build(0, 1, 2, 3, 4, 5, 6, 7), 2)))
    assert.Equal(
        t, true, compare(build(0, 2, 3, 4, 5, 6, 7),
            removeNthFromEnd(build(0, 1, 2, 3, 4, 5, 6, 7), 7)))
    assert.Equal(
        t, true, compare(build(1, 2, 3, 4, 5, 6, 7),
            removeNthFromEnd(build(0, 1, 2, 3, 4, 5, 6, 7), 8)))
}

func compareNode(o1, o2 *ListNode) bool {
    switch {
    case o1 == nil && o2 != nil, o2 == nil && o1 != nil, o1 == o2 && o1 == nil:
        return false
    default:
        return o1.Val == o2.Val
    }
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

func TestBuild(t *testing.T) {
    s := build(1, 2, 3, 4, 5, 6, 7)
    t.Log(s)
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
