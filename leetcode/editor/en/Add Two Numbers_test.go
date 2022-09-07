package solve

import (
    "testing"
    
    . "github.com/zhangli2946/tiantai/common"
)

func TestAddTwoNumbers(t *testing.T) {
    _ = addTwoNumbers
    _ = addTwoNumbers(&ListNode{9, nil}, &ListNode{9, &ListNode{9, &ListNode{9, nil}}})
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var longer, shorter = l1, l2
    var carry = 0
    for longer != nil {
        longer.Val += carry
        if shorter != nil {
            longer.Val = longer.Val + shorter.Val
        }
        longer.Val, carry = longer.Val%10, longer.Val/10
        switch {
        case shorter == nil:
        case longer.Next == nil && shorter.Next != nil:
            longer.Next, shorter.Next = shorter.Next, longer.Next
        }
        if tmp := longer.Next; tmp == nil && carry > 0 {
            longer.Next, carry = &ListNode{carry, nil}, 0
        }
        longer = longer.Next
        if shorter != nil {
            shorter = shorter.Next
        }
    }
    return l1
}

//leetcode submit region end(Prohibit modification and deletion)
