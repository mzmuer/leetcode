// 543. 二叉树的直径
// https://leetcode-cn.com/problems/diameter-of-binary-tree/

package question543

import (
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	x := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	t.Log(diameterOfBinaryTree(x))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	dp(root, &ans)
	return ans - 1
}

func dp(node *TreeNode, maxDepth *int) int {
	if node == nil {
		return 0
	}

	l := dp(node.Left, maxDepth)
	r := dp(node.Right, maxDepth)
	if *maxDepth < l+r+1 {
		*maxDepth = l + r + 1
	}

	if l > r {
		return l + 1
	}

	return r + 1
}
