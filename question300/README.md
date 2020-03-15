## 题目：[最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

给定一个无序的整数数组，找到其中最长上升子序列的长度。

**示例1:**
>输入: [10,9,2,5,3,7,101,18]  
>输出: 4   
>解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:
* 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
* 你算法的时间复杂度应该为 O(n2) 。
     
## 思路
1. 循环以每一个元素为起点遍历
2. 遍历到比上一个个元素大的值就增加子序列的长度
3. 为了使上升序列最长，新增的序列值应该尽可能接近上一个值

## 实现
```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res = 1
	for i := 0; i < len(nums)-1; i++ {
		var upup, up, tmp = math.MaxInt64, nums[i], 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > up {
				tmp++
				upup = up
				up = nums[j]
				if tmp > res {

					res = tmp
				}
			} else if nums[j] < up && nums[j] > upup {
				up = nums[j]
			}
		}
	}

	return res
}

```