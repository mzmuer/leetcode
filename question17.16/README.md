## 题目：[按摩师](https://leetcode-cn.com/problems/the-masseuse-lcci/)

一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

**示例1:**
>输入： [1,2,3,1]  
>输出： 4  
>解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。

**示例2:**
>输入： [2,7,9,3,1]  
>输出： 12  
>解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。

**示例3:**
>输入： [2,1,4,5,3,1,1,3]  
>输出： 12  
>解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。

## 思路
1. 动态规划,TODO:

## 实现
```go
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp0 := 0
	dp1 := nums[0]

	for i := 1; i < len(nums); i++ {
		tdp0 := _max(dp0, dp1) // 计算 dp[i][0]
		tdp1 := dp0 + nums[i]  // 计算 dp[i][1]

		dp0 = tdp0 // 用 dp[i][0] 更新 dp_0
		dp1 = tdp1 // 用 dp[i][1] 更新 dp_1
	}

	return _max(dp0, dp1)
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}
```