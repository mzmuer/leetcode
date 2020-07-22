// 154. 寻找旋转排序数组中的最小值 II
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

package question154

import (
	"testing"
)

func Test_findMin(t *testing.T) {
	a1 := []int{1, 3, 5}
	t.Log(findMin(a1))

	a2 := []int{2, 2, 2, 0, 1}
	t.Log(findMin(a2))
}

// 暴力法
// func findMin(nums []int) int {
// 	if len(nums) < 2 {
// 		return nums[0]
// 	}
//
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i-1] > nums[i] {
// 			return nums[i]
// 		}
// 	}
//
// 	return nums[0]
// }

// 二分法
func findMin(nums []int) int {
	var (
		low  = 0
		high = len(nums) - 1
	)

	for low < high {
		mid := (high + low) / 2
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}

	return nums[low]
}
