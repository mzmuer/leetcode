// 70. 爬楼梯
// https://leetcode-cn.com/problems/climbing-stairs/

package question70

import (
	"testing"
)

func Test_climbStairs(t *testing.T) {
	t.Log(climbStairs(2))  // 3
	t.Log(climbStairs(3))  // 3
	t.Log(climbStairs(44)) // 1134903170
}

// dp[n] = dp[n-1]+dp[n-2]
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	dp := [2]int{0, 1}
	for i := 2; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}

	return dp[0] + dp[1]
}
