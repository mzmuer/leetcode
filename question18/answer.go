// 18. 四数之和
// https://leetcode-cn.com/problems/4sum/

package question18

import "sort"

// 思路：套用三数之和，先确定两个数
func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		// 最小和都已经大于target
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// 最小和都已经大于target
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1          // 左边的数
			right := len(nums) - 1 // 右边的数

			for left < right {
				result := nums[i] + nums[j] + nums[left] + nums[right]
				if result == target { // 达到条件
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					// 可能还有其他组合
					left++
					right--
				} else if result < target { // 如果左边太小
					for left < right-1 && nums[left] == nums[left+1] {
						left++
					}
					left++
				} else { // 右边太小
					for left+1 < right && nums[right] == nums[right-1] {
						right--
					}
					right--
				}
			}
		}
	}

	return res
}
