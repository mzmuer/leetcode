## 题目：[翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)

翻转一棵二叉树。

**示例:**
>输入: 
> 
>          4
>        /   \
>       2     7
>      / \   / \
>     1   3 6   9
>
>输出： 
> 
>          4  
>        /   \  
>       7     2  
>      / \   / \  
>     9   6 3   1  

     
## 思路
1. 使用递归回溯，直接交换左右节点就好。
2. 交换 `root` 的 `left` 和 `right` 之前，交换好 `root.left` 的 `left` 和 `right`, 以此来递归，直到没有叶节点开始回溯。就可以完成整个树的翻转。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question226/answer_test.go)
```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
```