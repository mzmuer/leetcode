## 题目：[两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

**你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。 

**示例1:**
>给定 1->2->3->4, 你应该返回 2->1->4->3.

## 思路
1. 迭代链表，找出要交换的两个节点 `node1` 和 `node2`
2. 然后交换 `node1` 和 `node2`，再找出下一组要交换的节点，直到链表末尾

## [实现](https://github.com/mzmuer/leetcode/blob/master/question24/answer_test.go)
```go
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	return dummyHead.Next
}
```