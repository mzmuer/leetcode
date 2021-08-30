// 92. 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii/

package question92

import (
	"testing"
)

func printlnTree(t *testing.T, l *ListNode) {
	tmp := l
	for tmp != nil {
		t.Log(tmp.Val)
		tmp = tmp.Next
	}
}

func Test_reverseBetween(t *testing.T) {
	t1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}
	ansT1 := reverseBetween(t1, 2, 4)
	printlnTree(t, ansT1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

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
