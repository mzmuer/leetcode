// 122. 买卖股票的最佳时机 II
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

package question122

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	x := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(x)
	t.Log("ans", res, res == 7)

	x = []int{7, 6, 4, 3, 1}
	res = maxProfit(x)
	t.Log("ans", res, res == 0)
}

func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}

	return ans
}
