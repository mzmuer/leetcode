// 19. 删除链表的倒数第N个节点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

package question19

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路一: 遍历得到list长度，再次遍历删除
// 思路二：双指针 ptr和 ptr-n。当ptr移动到终点时 删除ptr-n
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	virtualNode := &ListNode{Next: head}
	ptr1 := head
	ptr2 := virtualNode

	// 移动达到 ptr1和ptr2的差值为n
	for i := 0; i < n; i++ {
		ptr1 = ptr1.Next
	}

	// 同时移动
	for ptr1 != nil {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	ptr2.Next = ptr2.Next.Next
	head = virtualNode.Next
	virtualNode.Next = nil

	return head
}
