## 题目：[最接近的三数之和](https://leetcode-cn.com/problems/3sum-closest/)

给定一个包括 n 个整数的数组 `nums` 和 一个目标值 `target`。找出 `nums` 中的三个整数，使得它们的和与 `target` 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

**示例1:**
>输入：nums = [-1,2,1,-4], target = 1  
>输出：2  
>解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。  

**提示：**
* `3 <= nums.length <= 10^3`
* `-10^3 <= nums[i] <= 10^3`
* `-10^4 <= target <= 10^4`
     
## 思路
1. 先对数组进行排序，然后使用双指针遍历。
2. 对于三元组，先固定第一个数，然后用双指针`left`和`right`,如果小于`target`那么`left`就往右边移动，否则`right`就往左边移动。直到找到最接近的值或者相等的值。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question16/answer_test.go)
```go
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
```