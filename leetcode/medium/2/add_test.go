package _2

import "testing"

func TestAddTwoNumbers(t *testing.T) {
    longer, shorter, _ := Build(999), Build(99), Build(807)
    result2 := addTwoNumbers(longer, shorter)
    show(t, result2)
}

func show(t *testing.T, r *ListNode) {
    for ; r != nil; r = r.Next {
        t.Logf("%d", r.Val)
    }
}

func TestBuild(t *testing.T) {
    show(t, Build(123456))
}

func Build(i int) (ret *ListNode) {
    ret = &ListNode{}
    h := ret
    for ; i >= 10; i /= 10 {
        h.Val = i % 10
        h.Next = &ListNode{}
        h = h.Next
    }
    h.Val = i
    return
}
