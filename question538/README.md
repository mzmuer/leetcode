## 题目：[把二叉搜索树转换为累加树](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/)

给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。

**示例1:**
>输入: 原始二叉搜索树:  
>
>              5
>            /   \
>           2     13
>  
>输出: 转换为累加树:  
>  
>             18
>            /   \
>          20     13
     
## 思路
1. 对于二叉搜索树，反向中序遍历的结果就是降序，累加就可以得到最后的结果。
2. 使用 Morris Traversal算法来中序遍历，可以降低空间复杂度到O(1)

## 实现
```go
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	
	dfs(root)

	return root
}
```