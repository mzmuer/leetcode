## 题目：[最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)

给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

**示例1:**
>输入:  
>[  
>  [1,3,1],  
>  [1,5,1],  
>  [4,2,1],  
>]  
>输出: 7  
>解释: 因为路径 1→3→1→1→1 的总和最小。  
     
## 思路
1. 因为只能向下或者向右运动，所以n,m的最小路径可以使用动态规划
2. 动态规划`dp(n,m) = min(dp(n-1,m),(n,m-1))`
3. 考虑情况: 
* `n== 0` 那么 `dp(0,m) = dp(0,m-1)+ grad[0][m]`
* `m==0` 那么 `dp(n,0) = dp(n-1,0)+ grad[n][0]`
* `dp(n,m) = min(dp(n-1,m),(n,m-1))+ grad[n][m]`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question64/answer_test.go)
```go
func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				f[0][0] = grid[0][0]
				continue
			} else if i == 0 {
				f[i][j] = f[i][j-1] + grid[i][j]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + grid[i][j]
			} else {
				f[i][j] = _min(f[i][j-1], f[i-1][j]) + grid[i][j]
			}
		}
	}

	return f[n-1][m-1]
}

func _min(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}
```