// 面试题 17.16. 按摩师
// https://leetcode-cn.com/problems/the-masseuse-lcci/

package question17_16

import (
	"math"
	"testing"
)

func Test_massage(t *testing.T) {
	a1 := []int{1, 2, 3, 1}
	t.Log(massage(a1) == 4)
}

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp0 := 0
	dp1 := nums[0]

	for i := 1; i < len(nums); i++ {
		tdp0 := _max(dp0, dp1) // 计算 dp[i][0]
		tdp1 := dp0 + nums[i]  // 计算 dp[i][1]

		dp0 = tdp0 // 用 dp[i][0] 更新 dp_0
		dp1 = tdp1 // 用 dp[i][1] 更新 dp_1
	}

	return _max(dp0, dp1)
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
