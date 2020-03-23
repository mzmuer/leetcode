// 876. 链表的中间结点
// https://leetcode-cn.com/problems/middle-of-the-linked-list/

package question876

import (
	"testing"
)

func Test_middleNode(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	res := middleNode(l1)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}

	t.Log("========================================")
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res = middleNode(l2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}

		p1 = p1.Next
	}

	return p1
}
