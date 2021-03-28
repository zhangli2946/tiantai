package _19

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var (
        curr = head
        ret  *ListNode
        loh  int
    )
    for ; curr != nil; curr, loh = curr.Next, loh+1 {
        if ret != nil {
            ret = ret.Next
        }
        if loh == n {
            ret = head
        }
    }
    switch {
    case loh == n:
        return head.Next
    case ret == nil:
        return head
    case ret.Next != nil:
        ret.Next = ret.Next.Next
    }
    return head
}
