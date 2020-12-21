## 题目：[使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)

数组的每个索引作为一个阶梯，第 `i`个阶梯对应着一个非负数的体力花费值 `cost[i]`(索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

**示例1:**
>输入: cost = [10, 15, 20]  
>输出: 15  
>解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。  

**示例1:**
>输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]  
>输出: 6  
>解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费  6。

注意：
1. `cost` 的长度将会在 `[2, 1000]`。
2. 每一个 `cost[i]` 将会是一个Integer类型，范围为 `[0, 999]`。

## 思路
1. `dp[i]` = 到达第 `i` 级楼下，所需要的*最少*体力花费
2. 状态转移方程为：  
    `dp[i] = min{dp[i-1], dp[i-2]}`
3. 初始状态为 `dp[0] = cost[0]`, `dp[1] = cost[1]`
4. 最后消耗的最少体力应该为： `min{dp[n-1], dp[n-2]}`
5. 由于 `dp[i]` 只和 `dp[i-1]` 和 `dp[i-2]` 相关，所以可以用滚动数组优化空间到**O(1)**
    
## [实现](https://github.com/mzmuer/leetcode/blob/master/question746/answer_test.go)
```go
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	var dp = [2]int{cost[0], cost[1]}
	for i := 2; i < n; i++ {
		dp[i%2] = _min(dp[0], dp[1]) + cost[i]
	}

	return _min(dp[0], dp[1])
}

func _min(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}

```