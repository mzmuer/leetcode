// 26. 删除排序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

package question26

// 思路 遍历数组，不重复长度++， 替换不重复的值到前面
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l++
		}
	}

	return l
}
