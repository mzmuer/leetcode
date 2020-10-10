// 142. 环形链表 II
// https://leetcode-cn.com/problems/linked-list-cycle-ii/

package question142

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
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

	t.Log(detectCycle(l1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}
//
// 	nodeMap := make(map[*ListNode]struct{})
// 	node := head
//
// 	for node != nil {
// 		if _, ok := nodeMap[node]; ok {
// 			return node
// 		} else {
// 			nodeMap[node] = struct{}{}
// 		}
// 		node = node.Next
// 	}
//
// 	return nil
// }
