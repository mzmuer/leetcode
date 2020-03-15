// 300. 最长上升子序列
// https://leetcode-cn.com/problems/longest-increasing-subsequence/

package question300

import (
	"math"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	a1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	t.Log(lengthOfLIS(a1) == 4)

	a2 := []int{2, 2}
	t.Log(lengthOfLIS(a2) == 1)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res = 1
	for i := 0; i < len(nums)-1; i++ {
		var upup, up, tmp = math.MaxInt64, nums[i], 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > up {
				tmp++
				upup = up
				up = nums[j]
				if tmp > res {

					res = tmp
				}
			} else if nums[j] < up && nums[j] > upup {
				up = nums[j]
			}
		}
	}

	return res
}
