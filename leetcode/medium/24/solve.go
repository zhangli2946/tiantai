package _24

type ListNode struct {
    Val  int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    var (
        curr = head
    )
    for ; curr != nil; curr = curr.Next {
        if p := curr.Next; p != nil {
            curr.Val, p.Val = p.Val, curr.Val
            curr = curr.Next
        }
    }
    return head
}
