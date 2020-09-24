// 501. 二叉搜索树中的众数
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

package question501

import (
	"testing"
)

func Test_findMode(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	// morrisTraversal(t1)
	// reverseMorrisTraversal(t1)
	t.Log(findMode(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {

	var (
		base, count, maxCount int
		answer                []int
	)
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var (
		cur  = root
		prev *TreeNode
	)

	// morris 中序遍历
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}

		prev = cur.Left
		for prev.Right != nil && prev.Right != cur {
			prev = prev.Right
		}

		if prev.Right == nil {
			prev.Right = cur
			cur = cur.Left
		} else {
			prev.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}

	return answer
}
