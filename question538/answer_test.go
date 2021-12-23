// 538. 把二叉搜索树转换为累加树
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

package question538

import (
	"fmt"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	t1 := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 4},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val:   9,
				Right: &TreeNode{Val: 10},
			},
			Right: &TreeNode{Val: 15},
		},
	}

	// morrisTraversal(t1)
	// reverseMorrisTraversal(t1)
	t.Log(convertBST(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	// sum := 0
	// var dfs func(root *TreeNode)
	// dfs = func(root *TreeNode) {
	// 	if root != nil {
	// 		dfs(root.Right)
	// 		sum += root.Val
	// 		root.Val = sum
	// 		dfs(root.Left)
	// 	}
	// }
	//
	// dfs(root)

	var (
		cur  = root
		prev *TreeNode
		sum  int
	)

	for cur != nil {
		if cur.Right == nil {
			// fmt.Println(cur.Val)
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
			continue
		}

		prev = cur.Right
		for prev.Left != nil && prev.Left != cur {
			prev = prev.Left
		}

		if prev.Left == nil {
			prev.Left = cur
			cur = cur.Right
		} else {
			prev.Left = nil
			// fmt.Println(cur.Val)
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
		}
	}

	return root
}

// 中序遍历
func inorderMorrisTraversal(root *TreeNode) {
	var (
		cur  = root
		prev *TreeNode
	)

	for cur != nil {
		// 如果当前节点的左孩子为空，
		// 则输出当前节点并将其右孩子作为当前节点
		if cur.Left == nil {
			fmt.Println(cur.Val)
			cur = cur.Right
			continue
		}

		// 找到 当前节点(cur) 的前驱节点
		// 且前驱节点的右孩子不为 当前节点(cur)
		prev = cur.Left
		for prev.Right != nil && prev.Right != cur {
			prev = prev.Right
		}

		// 如果前驱节点的右孩子为空，
		// 将它的右孩子设置为当前节点
		// 当前节点(cur) 更新为当前节点的左孩子。
		if prev.Right == nil {
			prev.Right = cur
			cur = cur.Left
		} else {
			// 如果前驱节点的右孩子为当前节点
			// 将它的右孩子重新设为空（恢复树的形状）
			// 输出当前节点
			// 当前节点更新为当前节点的右孩子

			prev.Right = nil
			fmt.Println(cur.Val)
			cur = cur.Right
		}
	}
}

// 反向中序遍历
func reverseInorderMorrisTraversal(root *TreeNode) {
	var (
		cur  = root
		prev *TreeNode
	)

	for cur != nil {
		if cur.Right == nil {
			fmt.Println(cur.Val)
			cur = cur.Left
			continue
		}

		prev = cur.Right
		for prev.Left != nil && prev.Left != cur {
			prev = prev.Left
		}

		if prev.Left == nil {
			prev.Left = cur
			cur = cur.Right
		} else {
			prev.Left = nil
			fmt.Println(cur.Val)
			cur = cur.Left
		}
	}
}
