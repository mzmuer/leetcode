## 题目：[二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

**示例:**
>你可以将以下二叉树：  
>  
>    1  
>   / \  
>  2   3  
>     / \  
>    4   5  
>  
>序列化为 "[1,2,3,null,null,4,5]"  

**提示:** 你也可以采用其他序列划的格式或方法解决这个问题。

**说明:** 不要使用类的成员/全局/静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

## 思路
1. 用BFS遍历树(方法很多比如DFS)。
2. 遍历的时候记录下每个字段的值得到编码。
3. 解码的时候先获取root，用一个queue暂存一下当前层。每次获取两个值，填入父节点`(获取到的两个值，一定是当前节点的子节点)`。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question297/answer_test.go)
```go
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var (
		ans   []string
		queue = make([]*TreeNode, 0)
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		// pop
		item := queue[0]
		queue = queue[1:]

		if item != nil {
			ans = append(ans, strconv.Itoa(item.Val))
			queue = append(queue, item.Left)
			queue = append(queue, item.Right)
		} else {
			ans = append(ans, "null")
		}
	}

	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	arr := strings.Split(data, ",")
	val, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: val}

	var queue = make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 1; i < len(arr); i += 2 {
		// pop
		item := queue[0]
		queue = queue[1:]

		if arr[i] != "null" {
			leftVal, _ := strconv.Atoi(arr[i])
			item.Left = &TreeNode{Val: leftVal}
			queue = append(queue, item.Left)
		}

		if arr[i+1] != "null" {
			rightVal, _ := strconv.Atoi(arr[i+1])
			item.Right = &TreeNode{Val: rightVal}
			queue = append(queue, item.Right)
		}
	}

	return root
}
```