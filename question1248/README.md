## 题目：[统计「优美子数组」](https://leetcode-cn.com/problems/count-number-of-nice-subarrays/)

给你一个整数数组 `nums` 和一个整数 `k`。

如果某个 连续 子数组中恰好有 `k` 个奇数数字，我们就认为这个子数组是「**优美子数组**」。

请返回这个数组中「优美子数组」的数目。

**示例1:**
>输入：nums = [1,1,2,1,1], k = 3  
输出：2  
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。

**示例2:**
>输入：nums = [2,4,6], k = 1  
输出：0  
解释：数列中不包含任何奇数，所以不存在优美子数组。

**示例3:**
>输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2  
输出：16

提示：
* `1 <= nums.length <= 50000`
* `1 <= nums[i] <= 10^5`
* `1 <= k <= nums.length`

## 思路
0. 滑动窗口，不断右移 `right` 指针来扩大滑动窗口，使其包含 `k` 个奇数
1. 统计第一个奇数左边的数字为`leftEvenCnt`，和最后一个奇数右边的偶数数字为`rightEvenCnt`。
2. 最终优美子数组的的组合数为`(leftEvenCnt + 1) * (rightEvenCnt + 1)`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1248/answer_test.go)
```go
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
```