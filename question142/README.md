## 题目：[二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 `null`。

为了表示给定链表中的环，我们使用整数 `pos` 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 `pos` 是 `-1`，则在该链表中没有环。

说明：不允许修改给定的链表。

**进阶：**  
你是否可以不用额外空间解决此题？

**示例1:**
>![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png)
>
> **输入：** head = [1,2], pos = 0  
> **输出：** tail connects to node index 0  
> **解释：** 链表中有一个环，其尾部连接到第一个节点。

**示例2:**
>![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)
>
> **输入：** head = [1,2], pos = 0  
> **输出：** tail connects to node index 0  
> **解释：** 链表中有一个环，其尾部连接到第一个节点。

>**示例3:**
>![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)
>
> **输入：** head = [1], pos = -1  
> **输出：** no cycle  
> **解释：** 链表中没有环。 

  
## 思路一
1. 使用map来存在访问过的节点，当再次访问到的时候就是入环的第一个节点。

## 思路二
1. 使用快慢指针。![](https://assets.leetcode-cn.com/solution-static/142/142_fig1.png)
2. 当slow和fast相遇时  
`a+n(b+c)+b=a+(n+1)b+nc`
3. 任意时刻，*fast* 指针走过的距离都为 *slow* 指针的 **2** 倍  
`a+(n+1)b+nc=2(a+b)⟹a=c+(n−1)(b+c)`
4. 所以从相遇点到入环点的距离加上 n-1 圈的环长，恰好等于从链表头部到入环点的距离。
5. 因此，当发现 *slow* 与 *fast* 相遇时，我们再额外使用一个指针 *ptr*。起始，它指向链表头部；随后，它和 *slow* 每次向后移动一个位置。最终，它们会在入环点相遇。


## [实现](https://github.com/mzmuer/leetcode/blob/master/question142/answer_test.go)
```go
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
```