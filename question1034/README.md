## 题目：[边界着色](https://leetcode-cn.com/problems/coloring-a-border/)

给你一个大小为 `m x n` 的整数矩阵 `grid` ，表示一个网格。另给你三个整数 `row`、`col` 和 `color` 。网格中的每个值表示该位置处的网格块的颜色。

当两个网格块的颜色相同，而且在四个方向中任意一个方向上相邻时，它们属于同一 *连通分量* 。

*连通分量的边界* 是指连通分量中的所有与不在分量中的网格块相邻（四个方向上）的所有网格块，或者在网格的边界上（第一行/列或最后一行/列）的所有网格块。

请你使用指定颜色 `color` 为所有包含网格块 `grid[row][col]` 的 *连通分量的边界* 进行着色，并返回最终的网格 `grid` 。

**示例1:**
>输入：grid = [[1,1],[1,2]], row = 0, col = 0, color = 3  
>输出：[[3,3],[3,2]]

**示例2:**
>输入：grid = [[1,2,2],[2,3,2]], row = 0, col = 1, color = 3  
>输出：[[1,3,3],[2,3,3]]

**示例3:**
>输入：grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2  
输出：[[2,2,2],[2,1,2],[2,2,2]]
 
提示:
* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 50`
* `1 <= grid[i][j], color <= 1000`
* `0 <= row < m`
* `0 <= col < n`
 
## 思路
1. 这个题首先要理解这个边界，相同颜色所连起来的区域
2. 所以我们从给定的中心点出发，往四个方向遍历（dfs和bfs都可以），找到遍历，并放入数组中，最后遍历边界数组，修改颜色

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1034/answer_test.go)
```go
var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	borders := make([]point, 0)
	originalColor := grid[row][col]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int)
	dfs = func(x, y int) {
		vis[x][y] = true
		isBorder := false
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isBorder {
			borders = append(borders, point{x, y})
		}
	}
	dfs(row, col)

	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}
```