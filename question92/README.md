## 题目：[反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)

给你单链表的头指针 `head` 和两个整数 `left` 和 `right` ，其中 `left <= right` 。请你反转从位置 `left` 到位置 `right` 的链表节点，返回 **反转后的链表** 。

**示例1:**
>![](https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg)
>
>输入：head = [1,2,3,4,5], left = 2, right = 4  
 输出：[1,4,3,2,5]

**示例2:**
>输入：head = [5], left = 1, right = 1  
 输出：[5]

提示：
* 链表中节点数目为 `n`
* `1 <= n <= 500`
* `-500 <= Node.val <= 500`
* `1 <= left <= right <= n`

## 思路
1. 当遍历到需要反转的位置的时候，逐个开始反转当前的节点。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question92/answer_test.go)
```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyNode.Next
}
```