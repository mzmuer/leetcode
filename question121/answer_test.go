// 121. 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

package question121

import (
	"math"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	x := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(x)
	t.Log("res", res, res == 5)

	x = []int{7, 6, 4, 3, 1}
	res = maxProfit(x)
	t.Log("res", res, res == 0)
}

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
