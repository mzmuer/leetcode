## 题目：[二叉树的直径](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

**示例1:**
>给定二叉树 
>          1  
>         / \  
>        2   3  
>       / \     
>      4   5    

返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
     
## 思路
1. 一个最长的路径，可以看作是一个根节点的左子树最深的遍历+右子树最深的遍历。
2. 用dfs搜索左右子树，获取最大结果。

## 实现
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	dp(root, &ans)
	return ans - 1
}

func dp(node *TreeNode, maxDepth *int) int {
	if node == nil {
		return 0
	}

	l := dp(node.Left, maxDepth)
	r := dp(node.Right, maxDepth)
	if *maxDepth < l+r+1 {
		*maxDepth = l + r + 1
	}

	if l > r {
		return l + 1
	}

	return r + 1
}
```