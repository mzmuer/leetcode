// 343. 整数拆分
// https://leetcode-cn.com/problems/integer-break/

package question343

import (
	"math"
	"testing"
)

func Test_integerBreak(t *testing.T) {
	t.Log(integerBreak(10))

	t.Log(integerBreak(2))
}

// 将n拆分成 j 和 n-j
// 那么 j 和 n-j的最大乘积取决于 k*(n-j) 或 k*dp[n-j]的最大乘积。
// dp[n-j]也就等于 n-j继续拆分的乘积最大值。直到 dp[1]=0
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = _max(curMax, _max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}

	return dp[n]
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
