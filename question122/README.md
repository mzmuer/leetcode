## 题目：[买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

**示例1:**
>输入: [7,1,5,3,6,4]  
>输出: 5  
>解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
     
**示例2:**
>输入: [7,6,4,3,1]  
>输出: 0  
>解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

## 思路
1. 最低点，最高点卖出才能实现利润最大化
2. 遍历一遍数组，想象成日期推进，不断以尽量最低的点买进，更高的点卖出。  
    2.1 如果有更低的价格，就记录下来。  
    2.2 如果有更高的利润就记录下利润。

## 实现
```go
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var (
		maxProfit int
		minPrice  = math.MaxInt32
	)

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}

	return maxProfit
}

```