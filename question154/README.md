## 题目：[寻找旋转排序数组中的最小值 II](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

注意数组中可能存在重复的元素。

**示例1:**
>输入: [1,3,5]  
>输出: 1  

**示例2:**
>输入: [2,2,2,0,1]  
输出: 0

## 思路一
1. 暴力解法。因为数组旋转为升序数组，所以直接从后往前遍历。找到第一个非递减的值，一定是最小值。

## 思路二
1. 二分查找。
2. 如果`nums[mid] < nums[high]`，那么mid右边的部分就没有意义，`high = mid`。
3. 如果`nums[mid] > nums[high]`，你那么mid左边的部分就没有意义,`low = mid+1`。
4. 如果`nums[mid] == nums[high]`，那么不知道最小值在左边还是右边，因为有相同值冗余，`hihg--`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question154/answer_test.go)
```go
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
```