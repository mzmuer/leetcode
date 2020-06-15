## 题目：[每日温度](https://leetcode-cn.com/problems/daily-temperatures/)

请根据每日 `气温` 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 `0` 来代替。

例如，给定一个列表 `temperatures = [73, 74, 75, 71, 69, 72, 76, 73]`，你的输出应该是 `[1, 1, 4, 2, 1, 1, 0, 0]`。

提示：气温 列表长度的范围是 `[1, 30000]`。每个气温的值的均为华氏度，都是在 `[30, 100]` 范围内的整数。
     
## 思路
1. 单调栈解法。栈中存储元素的**下标**
2. 遍历温度数组，如果栈为空，或者元素不小于栈顶元素，就入栈。
3. 如果小于栈顶元素，就弹出栈顶元素，`i-index` 得到 `ans[index]` 的结果。直到栈为空，或者元素不小于栈顶元素。将`i`入栈

## [实现](https://github.com/mzmuer/leetcode/blob/master/question739/answer_test.go)
```go
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
```