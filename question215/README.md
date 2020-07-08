## 题目：[数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

在未排序的数组中找到第 `k` 个最大的元素。请注意，你需要找的是数组排序后的第 `k` 个最大的元素，而不是第 `k` 个不同的元素。

**示例1:**
>输入: [3,2,1,5,6,4] 和 k = 2  
>输出: 5  

**示例1:**
>输入: [3,2,3,1,2,4,5,5,6] 和 k = 4  
>输出: 4  

**说明:**  
你可以假设 `k` 总是有效的，且 `1 ≤ k ≤` 数组的长度。
     
## 思路
1. 使用堆排序的方法。
2. 首先构建一个大根堆。
3. 然后移除堆顶元素`k-1`次，此时的堆顶元素就是第`k`个最大的元素。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question215/answer_test.go)
```go
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
```