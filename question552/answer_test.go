// 552. 学生出勤记录 II
// https://leetcode-cn.com/problems/student-attendance-record-ii/

package question552

import (
	"testing"
)

func Test_checkRecord(t *testing.T) {
	// t.Log(checkRecord(2))
	// t.Log(checkRecord(10101))

	t.Log(checkRecordTrim(2))
	t.Log(checkRecordTrim(10101))
}

// 动态规划
// 第i天的出勤情况 d[i][j][k]  第i天 有j次缺勤 有连续k次迟到的 几种情况的和(i<n, j<2, k<3)
// 根据限制条件 第i天出勤情况也就是：
// P: d[i][j][0] = d[i-1][j][k] {j<=1 k<=2}
// L: d[i][j][k+1] = d[i-1][j][k] {j<=1 k<2}
// A: d[i][1][0] = d[i-1][0][k] {k<=2}
func checkRecord(n int) int {
	dp := make([][2][3]int, n+1)
	dp[0][0][0] = 1
	var mod int = 1e9 + 7

	for i := 1; i <= n; i++ {
		// P
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
			}
		}
		// L
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 1; k++ {
				dp[i][j][k+1] = (dp[i][j][k+1] + dp[i-1][j][k]) % mod
			}
		}
		// A
		for k := 0; k <= 2; k++ {
			dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
		}
	}

	var ans int
	for j := 0; j <= 1; j++ {
		for k := 0; k <= 2; k++ {
			ans = (ans + dp[n][j][k]) % mod
		}
	}

	return ans
}

// 剪枝优化
// 第i天只和第i-1天有关，可以优化减少一维数组
func checkRecordTrim(n int) int {
	dp := [2][3]int{}
	dp[0][0] = 1
	var mod int = 1e9 + 7

	for i := 1; i <= n; i++ {
		newDp := [2][3]int{}
		// P
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				newDp[j][0] = (newDp[j][0] + dp[j][k]) % mod
			}
		}
		// L
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 1; k++ {
				newDp[j][k+1] = (newDp[j][k+1] + dp[j][k]) % mod
			}
		}
		// A
		for k := 0; k <= 2; k++ {
			newDp[1][0] = (newDp[1][0] + dp[0][k]) % mod
		}
		dp = newDp
	}

	var ans int
	for j := 0; j <= 1; j++ {
		for k := 0; k <= 2; k++ {
			ans = (ans + dp[j][k]) % mod
		}
	}

	return ans
}
