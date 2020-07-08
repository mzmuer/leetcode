// 739. 每日温度
// https://leetcode-cn.com/problems/daily-temperatures/

package question739

import (
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	// a1 := []int{9, 10, 11, 12}
	// t.Log(dailyTemperatures(a1))

	a2 := []int{73, 74, 75, 71, 72, 69, 72, 76, 73}
	t.Log("--", dailyTemperatures(a2))
}

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := make([]int, 0)

	for i, v := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < v {
			topIndex := len(stack) - 1
			ans[stack[topIndex]] = i - stack[topIndex]
			stack = stack[:topIndex]
		}

		stack = append(stack, i)
	}

	return ans
}
