## 题目：[长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

给定一个含有 **n** 个正整数的数组和一个正整数 **s** ，找出该数组中满足其和 **≥ s** 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

**示例1:**
>输入：s = 7, nums = [2,3,1,2,4,3]  
>输出：2  
>解释：子数组 [4,3] 是该条件下的长度最小的连续子数组。  
     
## 思路
1. 一般对于子数组问题都采用前缀和或者滑动窗口。
2. 这里采用滑动窗口，不断扩大窗口直到`sum`满足条件。记录下最小连续，然后减小窗口，重复判断`sum`是否满足条件，更新`ans`。直达达到结尾。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question209/answer_test.go)
```go
func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	var (
		ans              = math.MaxInt32
		left, right, sum int
	)

	for right < l {
		sum += nums[right]

		for sum >= s {
			if diff := right - left + 1; diff < ans {
				if ans = diff; ans == 1 {
					break
				}
			}

			sum -= nums[left]
			left++
		}

		right++
	}

	if math.MaxInt32 == ans {
		return 0
	}

	return ans
}
```