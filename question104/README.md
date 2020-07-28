## 题目：[二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

**说明**: 叶子节点是指没有子节点的节点。

**示例1:**
>给定二叉树 `[3,9,20,null,null,15,7]`，  
>  
>一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
>  
>        3
>       / \
>      9  20
>        /  \
>       15   7
>返回它的最大深度 3 。
     
## 思路
1. 深度优先递归，找到最大的值。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question104/answer_test.go)
```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return _max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
```