// 42. 接雨水
// https://leetcode-cn.com/problems/trapping-rain-water/

package question42

import (
	"testing"
)

func Test_trap(t *testing.T) {
	a1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	t.Log(trap(a1))
}

func trap(height []int) int {
	var (
		left, leftMax, rightMax, ans int
		right                        = len(height) - 1
	)

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}

			right--
		}
	}
	return ans
}
