package solve

import (
    "testing"
    
    "github.com/go-playground/assert/v2"
    . "github.com/zhangli2946/tiantai/common"
)

func TestPseudoPalindromicPathsInABinaryTree(t *testing.T) {
    assert.Equal(t, 2, pseudoPalindromicPaths(&TreeNode{
        Val: 2,
        Left: &TreeNode{
            Val: 3,
            Left: &TreeNode{
                Val:   3,
                Left:  nil,
                Right: nil,
            },
            Right: &TreeNode{
                Val:   1,
                Left:  nil,
                Right: nil,
            },
        },
        Right: &TreeNode{
            Val:  1,
            Left: nil,
            Right: &TreeNode{
                Val:   1,
                Left:  nil,
                Right: nil,
            },
        },
    }))
    assert.Equal(t, 2, pseudoPalindromicPaths1(0, &TreeNode{
        Val: 2,
        Left: &TreeNode{
            Val: 3,
            Left: &TreeNode{
                Val:   3,
                Left:  nil,
                Right: nil,
            },
            Right: &TreeNode{
                Val:   1,
                Left:  nil,
                Right: nil,
            },
        },
        Right: &TreeNode{
            Val:  1,
            Left: nil,
            Right: &TreeNode{
                Val:   1,
                Left:  nil,
                Right: nil,
            },
        },
    }))
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
func pseudoPalindromicPaths(root *TreeNode) (ret int) {
    var res int
    nums := make([]bool, 9)
    helper(root, &nums, &res)
    return res
}
func helper(root *TreeNode, nums *[]bool, res *int) {
    if (*nums)[root.Val-1] {
        (*nums)[root.Val-1] = false
    } else {
        (*nums)[root.Val-1] = true
    }
    if root.Right == nil && root.Left == nil {
        k := 0
        for _, i := range *nums {
            if i {
                k++
            }
            if k > 1 {
                break
            }
        }
        if k <= 1 {
            *res++
        }
    }
    
    if root.Left != nil {
        next := make([]bool, 9)
        copy(next, *nums)
        helper(root.Left, &next, res)
    }
    if root.Right != nil {
        next := make([]bool, 9)
        copy(next, *nums)
        helper(root.Right, &next, res)
    }
}

//leetcode submit region end(Prohibit modification and deletion)

func pseudoPalindromicPaths2(root *TreeNode) (ret int) {
    var queue = make([]*TreeNode, 0)
    if root == nil {
        return 0
    }
    root.Val = 1 << root.Val
    queue = append(queue, root)
    for i := 0; i < len(queue); i += 1 {
        l, r := queue[i].Left, queue[i].Right
        if l == nil && r == nil {
            ret += isPalindromicPath(queue[i].Val)
        }
        if l != nil {
            l.Val = queue[i].Val ^ 1<<l.Val
            queue = append(queue, l)
        }
        if r != nil {
            r.Val = queue[i].Val ^ 1<<r.Val
            queue = append(queue, r)
        }
    }
    return
}

func isPalindromicPath(state int) int {
    count := 0
    for ; state > 0; state >>= 1 {
        count += state & 1
    }
    if count < 2 {
        return 1
    } else {
        return 0
    }
}

func pseudoPalindromicPaths1(state uint16, root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    if state ^= 1 << root.Val; root.Left == nil && root.Right == nil {
        count := uint16(0)
        for ; state > 0; state >>= 1 {
            count += state & 1
        }
        if count < 2 {
            return 1
        } else {
            return 0
        }
    }
    return pseudoPalindromicPaths1(state, root.Left) + pseudoPalindromicPaths1(state, root.Right)
}

/*
    stack := make([]*TreeNode, 0)
    var u16 uint16 = 0
    for curr := root; curr != nil || len(stack) > 0; {
        for curr != nil {
            u16 ^= 1 << curr.Val
            stack, curr = append(stack, curr), curr.Left
        }
        curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
        fmt.Fprint(os.Stderr, curr.Val)
        curr = curr.Right
    }
    return

   var records = [10]int{}
        for i := range stack {
            records[stack[i].Val] += 1
        }
        if single := records[0]&1 + records[9]&1 +
            records[1]&1 + records[8]&1 +
            records[2]&1 + records[7]&1 +
            records[3]&1 + records[6]&1 +
            records[4]&1 + records[5]&1; single < 2 {
            ret += 1
        }
*/
