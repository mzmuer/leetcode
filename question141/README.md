## 题目：[二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 `pos` 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 `pos` 是 `-1`，则在该链表中没有环。**注意：`pos` 不作为参数进行传递**，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 `true` 。 否则，返回 `false` 。

**进阶：**  

你能用 *O(1)*（即，常量）内存解决此问题吗？

**示例1:**
>![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png)
>
> **输入：** head = [3,2,0,-4], pos = 1  
> **输出：** true  
> **解释：** 链表中有一个环，其尾部连接到第二个节点。

**示例2:**
>![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)
>
> **输入：** head = [1,2], pos = 0  
> **输出：** true  
> **解释：** 链表中有一个环，其尾部连接到第一个节点。  

**示例3:**
>![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)
>
> **输入：** head = [1], pos = -1  
> **输出：** false  
> **解释：** 链表中没有环。  
  
## 思路
1. 使用快慢指针，当慢指针和快指针相等时，链表中一定有环
2. 当快指针为空时，链表中一定没有环

## [实现](https://github.com/mzmuer/leetcode/blob/master/question114/answer_test.go)
```go
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
```