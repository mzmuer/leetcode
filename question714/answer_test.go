// 714. 买卖股票的最佳时机含手续费
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

package question714

import (
	"math"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	x := []int{1, 3, 7, 5, 10, 3}
	res := maxProfit(x, 3)
	t.Log("ans", res, res == 6)

	// x = []int{7, 6, 4, 3, 1}
	// res = maxProfit(x, 2)
	// t.Log("ans", res, res == 0)
	//
	x = []int{1, 3, 2, 8, 4, 9}
	res = maxProfit(x, 2)
	t.Log("ans", res, res == 8)
}

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

// func maxProfit(prices []int, fee int) int {
// 	var (
// 		buy = prices[0] + fee
// 		ans = 0
// 	)
//
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i]+fee < buy {
// 			buy = prices[i] + fee
// 		} else if prices[i] > buy {
// 			ans += prices[i] - buy
// 			buy = prices[i]
// 		}
// 	}
//
// 	return ans
// }
