// 912. 排序数组
// https://leetcode-cn.com/problems/sort-an-array/

package question912

import (
	"testing"
)

func Test_sortArray(t *testing.T) {
	a1 := []int{5, 2, 3, 1}
	t.Log(sortArray(a1))

	a2 := []int{5, 1, 1, 2, 0, 0}
	t.Log(sortArray(a2))
}

// 1 <= nums.length <= 50000
// -50000 <= nums[i] <= 50000
func sortArray(nums []int) []int {
	bucket := [100000]int{}
	for _, v := range nums {
		bucket[v+50000]++
	}

	var j int
	for i := 0; i < 100000; {
		if bucket[i] != 0 {
			nums[j] = i - 50000
			bucket[i]--
			j++
		}

		if bucket[i] == 0 {
			i++
		}
	}

	return nums
}

//func sortArray(nums []int) []int {
//	for gap := len(nums) / 2; gap > 0; gap /= 2 {
//		for i := 0; i < gap; i++ {
//			// 直接插入排序
//			for j := i + gap; j < len(nums); j += gap {
//				k := j
//
//				for k-gap >= 0 {
//					if nums[k] < nums[k-gap] {
//						nums[k-gap], nums[k] = nums[k], nums[k-gap]
//						k = k - gap
//					} else {
//						break
//					}
//				}
//
//			}
//		}
//	}
//
//	return nums
//}
