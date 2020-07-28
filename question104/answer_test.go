// 104. 二叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

package question104

import (
	"math"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	t1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	t.Log(maxDepth(t1))
}

// -----
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return _max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
