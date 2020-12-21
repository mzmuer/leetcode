## 题目：[买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

给定一个整数数组 `prices`，其中第 `i` 个元素代表了第 `i` 天的股票价格 ；非负整数 `fee` 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

**示例1:**
>输入: prices = [1, 3, 2, 8, 4, 9], fee = 2  
>输出: 8  
>解释: 能够达到的最大利润:    
>在此处买入 prices[0] = 1  
>在此处卖出 prices[3] = 8  
>在此处买入 prices[4] = 4  
>在此处卖出 prices[5] = 9  
>总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.  
    
注意：
* `0 < prices.length <= 50000`
* `0 < prices[i] < 50000`
* `0 <= fee < 50000`

## 思路一
贪心算法：
1. 假设买入价格为`buy = prices[i]+fee`
2. 如果遇到更低买入的价格(`prices[i]+fee`)那么就把 `buy` 替换为更低的买入价格
3. 如果遇到价格 `prices[i] > buy`, 就卖出 获得利润 `prices[i] - buy`；但是此时获得利润可能不是最高的利润。  
所以后面价格继续上涨，应该用更高的价格卖出。设置`buy = prices[i]`,后面价格上涨的利润为 `prices[i+n] - buy`。 重复2和3。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question714/answer_test.go)
```go
func maxProfit(prices []int, fee int) int {
	var (
		buy = prices[0] + fee
		ans = 0
	)

	for i := 1; i < len(prices); i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			ans += prices[i] - buy
			buy = prices[i]
		}
	}

	return ans
}
```

## 思路二
动态规划：
1. `dp[i][0]` = 第 `i` 天，不持有股票的最大收益。  `dp[i][1]` = 第 `i` 天持有股票的最大收益
2. 第 i 天，不持有股票状态转移方程：  
    2.1 前一天持有股票，卖出 `dp[i][0] = dp[i-1][1]+prices[i]`   
    2.2 前一天没有持有股票，不操作 `dp[i][0] = dp[i-1][0]`
3. 第 i 天，持有状态转移方程：  
    3.1 前一天持有股票， 不操作 `dp[i][1] = dp[i-1][1]`  
    3.2 前一天不持有股票，买入 `dp[i][1] = dp[i-1][0] - prices[i]`

4. 第 i 天，获得最大利润的状态转移方程：  
    不持有股票的最大利润：`dp[i][0] = max{dp[i-1][0], dp[i-1][1]+prices[i]-fee}`  
    持有股票的最大利润：`dp[i][1] = max{dp[i-1][1], dp[i-1][0]-prices[i]}`
    
## [实现](https://github.com/mzmuer/leetcode/blob/master/question714/answer_test.go)
```go
func maxProfit(prices []int, fee int) int {
	// n := len(prices)
	// var dp = make([][2]int, n)
	// dp[0][1] = -prices[0]
	//
	// for i := 1; i < n; i++ {
	// 	dp[i][0] = _max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
	// 	dp[i][1] = _max(dp[i-1][1], dp[i-1][0]-prices[i])
	// }
	//
	// return dp[n-1][0]

	// 剪枝
	n := len(prices)
	var dp = [2]int{0, -prices[0]}

	for i := 1; i < n; i++ {
		dp[0] = _max(dp[0], dp[1]+prices[i]-fee)
		dp[1] = _max(dp[1], dp[0]-prices[i])
	}

	return dp[0]
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
```