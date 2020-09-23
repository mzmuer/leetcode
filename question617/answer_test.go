// 617. 合并二叉树
// https://leetcode-cn.com/problems/merge-two-binary-trees/

package question695

import (
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}
	t2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}

	root := mergeTrees(t1, t2)

	// 打印
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			t.Log(root.Val)
			dfs(root.Left)
		}
	}
	dfs(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}

	left := mergeTrees(t1.Left, t2.Left)
	right := mergeTrees(t1.Right, t2.Right)

	return &TreeNode{Val: t1.Val + t2.Val, Left: left, Right: right}
}
