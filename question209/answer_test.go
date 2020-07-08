// 209. 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

package question209

import (
	"math"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	a1 := []int{2, 3, 1, 2, 4, 3}
	t.Log(minSubArrayLen(7, a1))
}

func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	var (
		ans              = math.MaxInt32
		left, right, sum int
	)

	for right < l {
		sum += nums[right]

		for sum >= s {
			if diff := right - left + 1; diff < ans {
				if ans = diff; ans == 1 {
					break
				}
			}

			sum -= nums[left]
			left++
		}

		right++
	}

	if math.MaxInt32 == ans {
		return 0
	}

	return ans
}
