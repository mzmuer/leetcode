// 1047. 删除字符串中的所有相邻重复项
// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/

package question1047

import (
	"testing"
)

func Test_lastStoneWeight(t *testing.T) {
	t.Log(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	stack := make([]byte, 0)
	for i := range S {
		if len(stack) == 0 || stack[len(stack)-1] != S[i] {
			stack = append(stack, S[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
