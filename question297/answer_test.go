// 297. 二叉树的序列化与反序列化
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

package question297

import (
	"strconv"
	"strings"
	"testing"
)

func Test_Codec(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	obj := Constructor()
	data := obj.serialize(root)
	t.Log("serialize ----", data)

	ans := obj.deserialize(data)
	t.Log(ans)
}

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
