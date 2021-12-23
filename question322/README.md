## 题目：[零钱兑换](https://leetcode-cn.com/problems/coin-change/)

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 `-1`。

**示例1:**
>输入: coins = [1, 2, 5], amount = 11  
>输出: 3  
>解释: 11 = 5 + 5 + 1

**示例2:**
>输入: coins = [2], amount = 3  
>输出: -1



## 思路
1. 想要硬币最少，应该尽可能多的使用最大的硬币。
2. 先选择尽可能多的最大的硬币，超出总额后再递归下一层选择较小的硬币。
3. 可以使用乘法`k := amount / coins[0]` 计算出最大能仍几个，再逐渐减小。
4. 因为最先找到的并不一定是最优解，所以还是要递归完所有情况。
5. 利用ans剪枝，可以快速去掉一部分不情况

## 实现
```go
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 排序
	sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j] })

	ans := math.MaxInt64
	dp(coins, amount, 0, &ans)
	if ans == math.MaxInt64 {
		return -1
	}

	return ans
}

func dp(coins []int, amount, count int, ans *int) {
	if amount == 0 {
		if *ans > count {
			*ans = count
		}
		return
	}

	if len(coins) == 0 {
		return
	}

	for k := amount / coins[0]; k >= 0 && k+count < *ans; k-- {
		dp(coins[1:], amount-k*coins[0], count+k, ans)
	}
}
```