## 题目：[爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

假设你正在爬楼梯。需要 *n* 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 *n* 是一个正整数。

**示例1:**
>输入： 2  
>输出： 2  
>解释： 有两种方法可以爬到楼顶。  
>1. 1 阶 + 1 阶  
>2. 2 阶  

**示例2:**
>输入： 3  
>输出： 3  
>解释： 有三种方法可以爬到楼顶。  
>1. 1 阶 + 1 阶 + 1 阶  
>2. 1 阶 + 2 阶  
>3. 2 阶 + 1 阶  
     
## 思路
1. 动态规划方程`dp[n] = dp[n-1]+dp[n-2]`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question70/answer_test.go)
```go
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	dp := [2]int{0, 1}
	for i := 2; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}

	return dp[0] + dp[1]
}
```