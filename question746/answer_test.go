// 746. 使用最小花费爬楼梯
// https://leetcode-cn.com/problems/min-cost-climbing-stairs/

package question714

import (
	"math"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	t.Log(minCostClimbingStairs([]int{1, 3, 7, 5, 10, 3}))
	t.Log(minCostClimbingStairs([]int{1, 3, 2, 8, 4, 9}))
}

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
