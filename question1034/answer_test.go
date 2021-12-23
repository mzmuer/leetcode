// 1034. 边界着色
// https://leetcode-cn.com/problems/coloring-a-border/

package question1034

import (
	"testing"
)

func Test_colorBorder(t *testing.T) {
	// // 输入：grid = [[1,1],[1,2]], row = 0, col = 0, color = 3
	// // 输出：[[3,3],[3,2]]
	t.Log(colorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3))

	// 输入：grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2
	// 输出：[[2,2,2],[2,1,2],[2,2,2]]
	t.Log(colorBorder([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2))
}

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
