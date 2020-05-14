## 题目：[只出现一次的数字](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

**示例1:**
>输入: [2,2,1]
输出: 1

**示例2:**
>输入: [4,1,2,1,2]
输出: 4
     
## 思路
1. 位操作，使用异或(`XRO`)
2. 遍历数组，XRO操作每个数字，相同的元素XRO之后结果为0，最后剩下的值就是单一的元素。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question136/answer_test.go)
```go
func singleNumber(nums []int) int {
	var ret int
	for i := range nums {
		ret ^= nums[i]
	}

	return ret
}
```