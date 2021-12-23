## 题目：[合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/)

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

**示例1:**
>输入: 
>
> 	    Tree 1                     Tree 2                  
>           1                         2                          
>          / \                       / \                          
>         3   2                     1   3                        
>        /                           \   \                      
>       5                             4   7                  
> 输出: 
> 合并后的树:
>
> 	     3
> 	    / \
> 	   4   5
> 	  / \   \ 
> 	 5   4   7
>

**注意:** 合并必须从两个树的根节点开始。
     
## 思路
1. 深度优先，递归回溯。
    * 如果对应的节点都为空，合并后的节点为空
    * 如果对应的节点有一个为空，合并后的节点为不为空的节点
    * 如果两个节点都不为空，合并后的节点为两个节点的和，以及两个节点的左右子节点合并后的节点。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question617/answer_test.go)
```go
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}

	left := mergeTrees(t1.Left, t2.Left)
	right := mergeTrees(t1.Right, t2.Right)

	return &TreeNode{Val: t1.Val + t2.Val, Left: left, Right: right}
}
```