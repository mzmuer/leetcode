// 108. 将有序数组转换为二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

package question108

import (
	"fmt"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	a1 := []int{-10, -3, 0, 5, 9}
	midEachTree(sortedArrayToBST(a1))

	// a2 := []int{3, 2, 1, 5, 6, 4}
	// t.Log(sortedArrayToBST(a2))
}

func midEachTree(root *TreeNode) {
	if root != nil {
		midEachTree(root.Left)
		fmt.Printf("%d ", root.Val)
		midEachTree(root.Right)
	}
}

// -----
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return dp(nums, 0, len(nums)-1)
}

func dp(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: arr[mid]}
	root.Left = dp(arr, left, mid-1)
	root.Right = dp(arr, mid+1, right)

	return root
}
