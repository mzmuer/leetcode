// 563. 二叉树的坡度
// https://leetcode-cn.com/problems/binary-tree-tilt/

package question563

import (
	"testing"
)

func Test_findTilt(t *testing.T) {
	// t.Log(findTilt(2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	var ans int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		a := dfs(root.Left)
		b := dfs(root.Right)
		ans += abs(a - b)
		return a + b + root.Val
	}

	dfs(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
