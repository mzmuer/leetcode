## 题目：[把二叉搜索树转换为累加树](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/)

给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：  
* 结点左子树中所含结点的值小于等于当前结点的值
* 结点右子树中所含结点的值大于等于当前结点的值
* 左子树和右子树都是二叉搜索树



**示例1:**
>给定 BST `[1,null,2,2]`   
>
>     1
>      \
>       2
>      /
>     2  
> `返回[2]`.
     
## 思路
1. 遍历树寻找众数的。

## 实现
```go
func findMode(root *TreeNode) []int {

	var (
		base, count, maxCount int
		answer                []int
	)
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var (
		cur  = root
		prev *TreeNode
	)

	// morris 中序遍历
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
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
			update(cur.Val)
			cur = cur.Right
		}
	}

	return answer
}
```