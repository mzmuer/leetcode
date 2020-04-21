// 1248. 统计「优美子数组」
// https://leetcode-cn.com/problems/count-number-of-nice-subarrays/

package question1248

import "testing"

func Test_numberOfSubarrays(t *testing.T) {
	a1 := []int{1, 1, 2, 1, 1}
	t.Log(numberOfSubarrays(a1, 3) == 2)
}

func numberOfSubarrays(nums []int, k int) int {
	var left, right, oddCnt, ans int

	for right < len(nums) {
		if nums[right]%2 == 1 {
			oddCnt++
		}

		right++
		if oddCnt == k {
			tmp := right
			for right < len(nums) && nums[right]%2 == 0 {
				right++
			}

			rightEvenCnt := right - tmp
			leftEvenCnt := 0

			for nums[left]%2 == 0 {
				leftEvenCnt++
				left++
			}

			ans += (leftEvenCnt + 1) * (rightEvenCnt + 1)
			left++
			oddCnt--
		}
	}

	return ans
}
