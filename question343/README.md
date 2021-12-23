## 题目：[整数拆分](https://leetcode-cn.com/problems/integer-break/)

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:  
>输入: 2  
>输出: 1  
>解释: 2 = 1 + 1, 1 × 1 = 1。  

示例 2:  
>输入: 10  
>输出: 36  
>解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。  

**说明**: 你可以假设 n 不小于 2 且不大于 58。
     
## 思路
1. 对于一个整数n可以拆分为`j`和`n-j`，那么`j`可能的范围为`[1,n-1]`
2. 可以得出状态转移方程 `dp[n] = max(j*(n-j), j*dp[n-j])`
3. 要得到最大的`dp[n]`，需要对`j`可能范围内的每一个值计算，得到最大的值。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question343/answer_test.go)
```go
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = _max(curMax, _max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}

	return dp[n]
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
```