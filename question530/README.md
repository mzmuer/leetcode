## 题目：[二叉搜索树的最小绝对差](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)

给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

**示例1:**
>输入：
>
>     1
>      \
>       3
>      /
>     2
> 
> 输出：  
> 1
>
>解释：  
>最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。

## 思路
1. 对于二叉搜索树，使用中序遍历可以输出升序序列。
2. 对于升序结果，最小的绝对值差，一定是两个相邻的数字之间。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question530/answer_test.go)
```go
func getMinimumDifference(root *TreeNode) int {
	var ans, pre = math.MaxInt32, -1

	var (
		cur  = root
		prev *TreeNode
	)

	// morris 中序遍历
	for cur != nil {
		if cur.Left == nil {
			if pre != -1 && cur.Val-pre < ans {
				ans = cur.Val - pre
				if ans == 0 {
					break
				}
			}
			pre = cur.Val
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
			if pre != -1 && cur.Val-pre < ans {
				ans = cur.Val - pre
				if ans == 0 {
					break
				}
			}
			pre = cur.Val
			cur = cur.Right
		}
	}

	return ans
}
```