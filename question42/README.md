## 题目：[接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)


给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

**示例1:**
>输入: [0,1,0,2,1,0,1,3,2,1,2,1]  
>输出: 6


## 思路
1. 当`left[i]<right[j]`（元素0-6）时,积水高度由`leftMax`决定; 当`left[i]>right[j]`时,积水的高度由于`rightMax`决定。
2. 所以我们使用双指针来遍历列表。根据`left[i]`和`right[j]`的关系，交替遍历列表。同时记录下`leftMax`和`rightMax`。累加就可以得到最后的结果。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question42/answer_test.go)
```go
func trap(height []int) int {
	var (
		left, leftMax, rightMax, ans int
		right                        = len(height) - 1
	)

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}
```