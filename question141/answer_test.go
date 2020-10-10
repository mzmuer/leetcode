// 141. 环形链表
// https://leetcode-cn.com/problems/linked-list-cycle/

package question141

import (
	"fmt"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	cycle1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: -4,
			},
		},
	}
	cycle1.Next.Next.Next = cycle1
	l1 := &ListNode{
		Val:  3,
		Next: cycle1,
	}

	fmt.Println(hasCycle(l1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next
	for fast != head {
		if fast == nil || fast.Next == nil {
			return false
		}

		head = head.Next
		fast = fast.Next.Next
	}

	return true
}
