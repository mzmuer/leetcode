// 322. 零钱兑换
// https://leetcode-cn.com/problems/coin-change/

package question322

import (
	"math"
	"sort"
	"testing"
)

func Test_coinChange(t *testing.T) {
	x := []int{346, 29, 395, 188, 155, 109}
	res := coinChange(x, 9401)
	t.Log("res", res, res == 26)
}

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
