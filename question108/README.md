## 题目：[将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

**示例1:**
>给定有序数组: [-10,-3,0,5,9],  
>  
>一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
>  
>          0  
>         / \
>       -3   9  
>       /   /  
>     -10  5  
     
## 思路
1. BST的中序遍历是升序的,所以以升序数组中的任意一个元素作为根节点，左边的树为左子树，右边的树为右子树。那么就是一个BST。
2. 为了满足高度平衡，我们选择中间节点为根节点。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question108/answer_test.go)
```go
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
```