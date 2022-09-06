package solve

import (
    "testing"
    
    . "github.com/zhangli2946/tiantai/common"
)

func Test(t *testing.T) {
    _ = pruneTree
    _ = visitNR(nil)
    _ = visitR(nil)
    r := pruneTree(&TreeNode{
        Val: 1,
        Left: &TreeNode{
            0,
            &TreeNode{
                0,
                nil,
                nil,
            },
            &TreeNode{
                0,
                nil,
                nil,
            },
        },
        Right: &TreeNode{
            1,
            &TreeNode{
                0,
                nil,
                nil,
            },
            &TreeNode{
                1,
                nil,
                nil,
            },
        },
    })
    _ = r.Left
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
func pruneTree(n *TreeNode) *TreeNode {
    //stack := make([]*TreeNode, 8, 8)
    //si := 0
    if n == nil {
        return nil
    }
    n.Left, n.Right = pruneTree(n.Left), pruneTree(n.Right)
    if n.Left == nil && n.Right == nil && n.Val == 0 {
        return nil
    }
    return n
}

func visitNR(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    stack, top := make([]*TreeNode, 256, 256), 0
    var peek *TreeNode
    for stack[top] = root; top < 256; top += 1 {
        if peek = stack[top]; peek == nil {
            continue
        }
        if peek.Left != nil {
            stack[(top<<1)+1] = peek.Left
        }
        if peek.Right != nil {
            stack[(top+1)<<1] = peek.Right
        }
    }
    var pi int
    for top = 255; top > -1; top -= 1 {
        if peek = stack[top]; peek == nil {
            continue
        }
        if 0 == top&1 {
            pi = (top - 1) >> 1
        } else {
            pi = (top) >> 1
        }
        if peek.Left == nil && peek.Right == nil && peek.Val == 0 {
            if pi < 0 {
                return nil
            }
            if 0 == top&1 {
                stack[pi].Right = nil
            } else {
                stack[pi].Left = nil
            }
        }
    }
    return stack[0]
}

func visitR(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    n := TreeNode{
        Left:  visitR(root.Left),
        Right: visitR(root.Right),
        Val:   root.Val,
    }
    if n.Left == nil && n.Right == nil && root.Val == 0 {
        return nil
    }
    return &n
}

//leetcode submit region end(Prohibit modification and deletion)
