package _206

import "github.com/zhangli2946/tiantai/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *common.ListNode) *common.ListNode {
    var ret = &common.ListNode{
        -1, nil,
    }
    for ; head != nil; head = head.Next {
        ret.Next = &common.ListNode{Val: head.Val, Next: ret.Next}
    }
    return ret.Next
}
func reverseList2(head *common.ListNode) *common.ListNode {
    var pre, next *common.ListNode
    for ; head != nil; {
        next, head.Next = head.Next, pre
        pre, head = head, next
    }
    return pre
}
