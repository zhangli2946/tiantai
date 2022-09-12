package solve

import (
    "testing"
    
    . "github.com/zhangli2946/tiantai/common"
)

func TestBinaryTreeInorderTraversal(t *testing.T) {
    _ = inorderTraversal(&TreeNode{
        Val: 1,
        Right: &TreeNode{
            Val:  2,
            Left: &TreeNode{Val: 3},
        },
    })
    _ = inorderTraversal(nil)
    _ = inorderTraversal(&TreeNode{1, nil, nil})
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var ret = make([]int, 0, 100)
    stack, top := make([]*TreeNode, 32, 32), -1
    for curr := root; curr != nil || top > -1; {
        for curr != nil {
            stack[top+1], top, curr = curr, top+1, curr.Left
        }
        curr, top = stack[top], top-1
        ret = append(ret, curr.Val)
        curr = curr.Right
    }
    
    return ret
}

//leetcode submit region end(Prohibit modification and deletion)
func inorderTraversal1(root *TreeNode) []int {
    if nil == root {
        return nil
    }
    ret := append([]int{}, inorderTraversal(root.Left)...)
    ret = append(ret, root.Val)
    return append(ret, inorderTraversal(root.Right)...)
}
