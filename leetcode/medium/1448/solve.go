package _1448

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func goodNodes(root *TreeNode) int {
    return count(root, root.Val)
}

var count func(*TreeNode, int) int

func m1(root *TreeNode, max int) (cnt int) {
    if root == nil {
        return
    }
    if root.Val >= max {
        cnt += 1
        max = root.Val
    }
    cnt += m1(root.Right, max)
    cnt += m1(root.Left, max)
    return
}

func init() {
    count = m1
    count = m2
}

// 可选的优化
func m2(root *TreeNode, _ int) int {
    pop, stack, cnt := (*TreeNode)(nil), []*TreeNode{root}, 1
    for l := 1; l > 0; l = len(stack) {
        pop, stack = stack[l-1], stack[:l-1]
        if pop.Right != nil {
            if pop.Val > pop.Right.Val {
                pop.Right.Val = pop.Val
            } else {
                cnt += 1
            }
            stack = append(stack, pop.Right)
        }
        if pop.Left != nil {
            if pop.Val > pop.Left.Val {
                pop.Left.Val = pop.Val
            } else {
                cnt += 1
            }
            stack = append(stack, pop.Left)
        }
    }
    return cnt
}
