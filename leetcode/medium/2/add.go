package _2

type ListNode struct {
    Val  int
    Next *ListNode
}

// 8ms 5m
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var (
        carry = 0
        ret   = &ListNode{}
        s     = ret
    )
    for l1 != nil || l2 != nil || carry > 0 {
        s.Next = &ListNode{}
        s = s.Next
        if l1 != nil {
            s.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            s.Val += l2.Val
            l2 = l2.Next
        }
        s.Val += carry
        if ans := s.Val - 10; ans >= 0 {
            carry, s.Val = 1, ans
        } else {
            carry = 0
        }
    }
    return ret.Next
}

// 4ms 4.8M
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
    var (
        carry           = 0
        shorter, longer = l1, l2
        prev            *ListNode
    )
    for ; longer != nil; {
        if shorter != nil {
            longer.Val = longer.Val + carry + shorter.Val
        } else {
            longer.Val = longer.Val + carry
        }
        if ans := longer.Val - 10; ans >= 0 {
            carry, longer.Val = 1, ans
        } else {
            carry = 0
        }
        switch {
        case shorter == nil:
            longer, prev = longer.Next, longer
        case longer.Next == nil && shorter.Next != nil:
            longer.Next, shorter.Next = shorter.Next, longer.Next
            fallthrough
        default:
            longer, shorter, prev = longer.Next, shorter.Next, longer
        }
    }
    if carry == 1 {
        prev.Next = &ListNode{1, nil}
    }
    return l2
}
