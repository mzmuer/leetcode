// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/

package question2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路: 同步遍历l1和l2,新建一个链表存储新的结果
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		newL    *ListNode
		carried int
		curNode = newL
	)

	for l1 != nil || l2 != nil {
		total := 0
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}

		total += carried
		if total >= 10 {
			carried = 1
			total -= 10
		} else {
			carried = 0
		}

		if curNode == nil {
			curNode = &ListNode{
				Val:  total,
				Next: nil,
			}
			newL = curNode
		} else {
			curNode.Next = &ListNode{
				Val:  total,
				Next: nil,
			}

			curNode = curNode.Next
		}
	}

	if carried != 0 && curNode != nil {
		curNode.Next = &ListNode{
			Val:  carried,
			Next: nil,
		}
	}

	return newL
}
