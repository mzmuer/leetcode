// 122. 买卖股票的最佳时机 III
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

package question123

import (
	"sort"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	//a1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	//res := maxProfit(a1)
	//t.Log("ans", res, res == 6)
	//
	//a2 := []int{1, 2, 3, 4, 5}
	//res = maxProfit(a2)
	//t.Log("ans", res, res == 4)

	a3 := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	res := maxProfit(a3)
	t.Log("ans", res, res == 13)
}

func maxProfit(prices []int) int {
	arr := make([]int, 0)

	for i := 0; i < len(prices); i++ {
		tmp := 0
		for ; i+1 < len(prices) && prices[i] < prices[i+1]; i++ {
			tmp += prices[i+1] - prices[i]
		}
		arr = append(arr, tmp)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	var ans int
	for i := 0; i < len(arr) && i < 2; i++ {
		ans += arr[i]
	}

	return ans
}
