## 题目：[排序数组](https://leetcode-cn.com/problems/sort-an-array/)

给你一个整数数组 `nums`，请你将该数组升序排列。

**示例1:**
>输入：nums = [5,2,3,1]
>输出：[1,2,3,5]

**示例2:**
>输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]

提示：
* `1 <= nums.length <= 50000`
* `-50000 <= nums[i] <= 50000`

## 思路
1. 排序的方法很多，这里提供几种

## [实现](https://github.com/mzmuer/leetcode/blob/master/question912/answer_test.go)

####1. 希尔排序
```go
func sortArray(nums []int) []int {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := 0; i < gap; i++ {
			// 直接插入排序
			for j := i + gap; j < len(nums); j += gap {
				k := j
				
				for k-gap >= 0 {
					if nums[k] < nums[k-gap] {
						nums[k-gap], nums[k] = nums[k], nums[k-gap]
						k = k - gap
					} else {
						break
					}
				}

			}
		}
	}

	return nums
}
```

####2. 基数排序
```go
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
```