package _1448

import (
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestSolve(t *testing.T) {
    root := &TreeNode{3, nil, nil}
    t.Run("only root", func(t *testing.T) {
        assert.Equal(t, goodNodes(root), 1)
    })
    t.Run("only root", func(t *testing.T) {
        root.Right = &TreeNode{4, nil, nil}
        assert.Equal(t, goodNodes(root), 2)
    })
    t.Run("only root", func(t *testing.T) {
        root.Right = &TreeNode{4, nil, nil}
        root.Left = &TreeNode{4, &TreeNode{5, nil, nil}, nil}
        assert.Equal(t, goodNodes(root), 4)
    })
}
