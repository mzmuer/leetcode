// 530. 二叉搜索树的最小绝对差
// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

package question530

import (
	"math"
	"testing"
)

func Test_getMinimumDifference(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}

	t.Log(getMinimumDifference(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var ans, pre = math.MaxInt32, -1

	var (
		cur  = root
		prev *TreeNode
	)

	// morris 中序遍历
	for cur != nil {
		if cur.Left == nil {
			if pre != -1 && cur.Val-pre < ans {
				ans = cur.Val - pre
				if ans == 0 {
					break
				}
			}
			pre = cur.Val
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
			if pre != -1 && cur.Val-pre < ans {
				ans = cur.Val - pre
				if ans == 0 {
					break
				}
			}
			pre = cur.Val
			cur = cur.Right
		}
	}

	return ans
}
