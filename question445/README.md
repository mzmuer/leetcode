## 题目：[鸡蛋掉落](https://leetcode-cn.com/problems/super-egg-drop/)

给你两个 **非空** 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

**示例1:**
>输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)  
输出：7 -> 8 -> 0 -> 7

## 思路
1. 用栈存下数字
2. 逐个弹出相加，得到最好的结果，注意进位就好

## [实现](https://github.com/mzmuer/leetcode/blob/master/question445/answer_test.go)
```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var (
		ans    *ListNode
		offset int
	)

	for len(stack1) != 0 || len(stack2) != 0 {
		cur := &ListNode{Next: ans}
		ans = cur

		if len(stack1) == 0 {
			val := stack2[len(stack2)-1]
			cur.Val = (val + offset) % 10
			offset = (val + offset) / 10

			stack2 = stack2[:len(stack2)-1]
		} else if len(stack2) == 0 {
			val := stack1[len(stack1)-1]
			cur.Val = (val + offset) % 10
			offset = (val + offset) / 10

			stack1 = stack1[:len(stack1)-1]
		} else {
			val := stack1[len(stack1)-1] + stack2[len(stack2)-1]
			cur.Val = (val + offset) % 10
			offset = (val + offset) / 10

			stack1 = stack1[:len(stack1)-1]
			stack2 = stack2[:len(stack2)-1]
		}
	}

	if offset > 0 {
		cur := &ListNode{
			Val:  offset,
			Next: ans,
		}
		ans = cur
	}

	return ans
}
```