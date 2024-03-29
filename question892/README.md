## 题目：[三维形体的表面积](https://leetcode-cn.com/problems/surface-area-of-3d-shapes/)

在 `N * N` 的网格上，我们放置一些 `1 * 1 * 1`  的立方体。

每个值 `v = grid[i][j]` 表示 `v` 个正方体叠放在对应单元格 `(i, j)` 上。

请你返回最终形体的表面积。

**示例1:**
>输入：[[2]]  
>输出：10

**示例2:**
>输入：[[1,2],[3,4]]  
输出：34

**示例3:**
>输入：[[1,0],[0,2]]  
输出：16

**示例4:**
>输入：[[1,1,1],[1,0,1],[1,1,1]]
输出：32

**示例5:**
>输入：[[2,2,2],[2,1,2],[2,2,2]]
输出：46

提示：
* `1 <= N <= 50`
* `0 <= grid[i][j] <= 50`

## 思路
1. `v > 0` 那么地面和顶部可以各贡献`1`的面积
2. 其余四面的面积，取决于网格中四个方向的高度，也就是`v`的大小

## [实现](https://github.com/mzmuer/leetcode/blob/master/question892/answer_test.go)
```go
func surfaceArea(grid [][]int) int {
	var ans int
	dt := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for row, rows := range grid {
		for col, v := range rows {
			if v == 0 {
				continue
			}

			// 上下两个面
			ans += 2

			// 其余四个面
			for _, d := range dt {
				r := row + d[0]
				c := col + d[1]
				if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
					ans += v
				} else if grid[r][c] < v {
					ans += v - grid[r][c]
				}
			}
		}
	}

	return ans
}
```