// 16. 最接近的三数之和
// https://leetcode-cn.com/problems/3sum-closest/

package question16

import (
	"math"
	"sort"
	"testing"
)

// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
func Test_threeSumClosest(t *testing.T) {
	a1 := []int{-100, -98, -2, -1}
	t.Log(threeSumClosest(a1, -101))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	minDiff := math.MaxInt32
	resSum := 0

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if (nums[i]+nums[i+1])-target > minDiff || minDiff == 0 { // 找到结果
			return resSum
		}

		left := i + 1          // 左边的数
		right := len(nums) - 1 // 右边的数

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := sum - target

			if diff == 0 {
				return sum
			} else if abs(diff) < minDiff {
				resSum = sum
				minDiff = abs(diff)
			}

			if diff < 0 {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else { // 右边太小
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}

	return resSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else if x == 0 {
		return 0
	}
	return x
}
