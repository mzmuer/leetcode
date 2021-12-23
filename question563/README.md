## 题目：[二叉树的坡度](https://leetcode-cn.com/problems/binary-tree-tilt/)

给定一个二叉树，计算 *整个树* 的坡度 。

一个树的 *节点的坡度* 定义即为，该节点左子树的节点之和和右子树节点之和的 *差的绝对值* 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。

*整个树* 的坡度就是其所有节点的坡度之和。

**示例1:**
>![示例1](https://assets.leetcode.com/uploads/2020/10/20/tilt1.jpg)
>输入：root = [1,2,3]  
>输出：1  
>解释：  
>节点 2 的坡度：|0-0| = 0（没有子节点）  
>节点 3 的坡度：|0-0| = 0（没有子节点）  
>节点 1 的坡度：|2-3| = 1（左子树就是左子节点，所以和是 2 ；右子树就是右子节点，所以和是 3 ）  
>坡度总和：0 + 0 + 1 = 1  
>

提示：
* 树中节点数目的范围在 `[0, 104]` 内
* `-1000 <= Node.val <= 1000`

## 思路
深度优先，回溯过程中可以得到左右子树的值。计算并保存结果

## [实现](https://github.com/mzmuer/leetcode/blob/master/question563/answer_test.go)
```go
func findTilt(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	var ans int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		a := dfs(root.Left)
		b := dfs(root.Right)
		ans += abs(a - b)
		return a + b + root.Val
	}

	dfs(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
```