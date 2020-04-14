// 445. 两数相加 II
// https://leetcode-cn.com/problems/add-two-numbers-ii/

package question445

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	res := addTwoNumbers(&l1, &l2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

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
