## 题目：[二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)

给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

**示例1:**
>输入: [1,2,3,null,5,null,4]  
>输出: [1, 3, 4]  
>解释:  
>  
>   1            <---  
> /   \  
>2     3         <---  
> \     \  
>  5     4       <---  

## 思路
1. 广度优先搜索,遍历每一层，记录最后一个节点就是最后的结果

## [实现](https://github.com/mzmuer/leetcode/blob/master/question199/answer_test.go)
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rightmostValueAtDepth := map[int]int{}
	maxDepth := -1

	nodeQueue := make([]*TreeNode, 0)
	depthQueue := make([]int, 0)
	nodeQueue = append(nodeQueue, root)
	depthQueue = append(depthQueue, 0)

	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		depth := depthQueue[0]
		depthQueue = depthQueue[1:]

		if node != nil {
			maxDepth = _max(maxDepth, depth)

			// 由于每一层最后一个访问到的节点才是我们要的答案，因此不断更新对应深度的信息即可
			rightmostValueAtDepth[depth] = node.Val

			nodeQueue = append(nodeQueue, node.Left)
			nodeQueue = append(nodeQueue, node.Right)
			depthQueue = append(depthQueue, depth+1)
			depthQueue = append(depthQueue, depth+1)
		}
	}

	rightView := make([]int, 0)
	for depth := 0; depth <= maxDepth; depth++ {
		rightView = append(rightView, rightmostValueAtDepth[depth])
	}

	return rightView
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
```