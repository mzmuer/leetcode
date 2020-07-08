// 215. 数组中的第K个最大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

package question215

import (
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	a1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	t.Log(findKthLargest(a1, 4) == 4)

	a2 := []int{3, 2, 1, 5, 6, 4}
	t.Log(findKthLargest(a2, 2) == 5)
}

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}

	l := len(nums)
	for i := 1; i < k; i++ {
		nums[0], nums[l-i] = nums[l-i], nums[0]
		adjustHeap(nums, 0, l-i-1)
	}

	return nums[0]
}

func adjustHeap(arr []int, i, l int) {
	tmp := arr[i]
	for k := i*2 + 1; k < l; k = k*2 + 1 {
		if k+1 < l && arr[k] < arr[k+1] {
			k++
		}

		if arr[k] > tmp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}

	arr[i] = tmp
}
