// 111. 二叉树的最小深度
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

package question111

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
	t.Log(minDepth(t1))

	t2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	t.Log(minDepth(t2))
}

// -----
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return _mix(minDepth(root.Left), minDepth(root.Right)) + 1
}

func _mix(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}
