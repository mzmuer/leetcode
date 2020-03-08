// 21. 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/

package question21

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：直接遍历合并
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		virtualNode = new(ListNode)
		curNode     = virtualNode
		cur1        = l1
		cur2        = l2
	)

	for cur1 != nil || cur2 != nil {
		if cur2 == nil || (cur1 != nil && cur1.Val <= cur2.Val) {
			curNode.Next = &ListNode{
				Val:  cur1.Val,
				Next: nil,
			}
			cur1 = cur1.Next
			curNode = curNode.Next
		} else if cur1 == nil || cur1.Val > cur2.Val {
			curNode.Next = &ListNode{
				Val:  cur2.Val,
				Next: nil,
			}
			cur2 = cur2.Next
			curNode = curNode.Next
		}
	}

	head := virtualNode.Next
	virtualNode.Next = nil
	return head
}
