// 226. 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/

package question226

import (
	"testing"
)

func Test_invertTree(t *testing.T) {
	t1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	var printTree func(root *TreeNode)
	printTree = func(root *TreeNode) {
		if root == nil {
			return
		}

		t.Log(root.Val)
		printTree(root.Left)
		printTree(root.Right)
	}

	printTree(invertTree(t1))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
