## 题目：[二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

给定一个二叉树，*原地*将它展开为一个单链表。

例如，给定二叉树
>         1  
>        / \  
>       2   5  
>      / \   \  
>     3   4   6

将其展开为：  
>     1  
>      \  
>       2  
>        \  
>         3  
>          \  
>           4  
>            \  
>             5  
>              \  
>               6  

## 思路
1. 因为最后形成的链表结构是前序遍历的结果。所以可以先前序遍历，但是不能在遍历过程中修改树。所以可以先用一个链表保存下遍历的节点顺序，然后组成结果。
2. 因为前序遍历的顺序为，根节点，左子树，右子树。所以  
    2.1 当左节点有右边子树的时候，那么左节点最右边的节点就是右节点的前置节点  
    2.2 当左节点没有右子树的时候，那么左节点自生就是右节点的前置指数

## [实现](https://github.com/mzmuer/leetcode/blob/master/question114/answer_test.go)
```go
func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
```