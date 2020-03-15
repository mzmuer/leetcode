// 695. 岛屿的最大面积
// https://leetcode-cn.com/problems/max-area-of-island/

package question695

import (
	"testing"
)

func Test_maxAreaOfIsland(t *testing.T) {
	a1 := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	t.Log(maxAreaOfIsland(a1) == 6)
}

func maxAreaOfIsland(grid [][]int) int {
	var res int

	for row, rows := range grid {
		for col, v := range rows {
			if v != 1 {
				continue
			}

			tmp := dp(row, col, grid)
			if res < tmp {
				res = tmp
			}
		}
	}

	return res
}

func dp(row, col int, grid [][]int) int {
	grid[row][col] = 0
	var res = 1

	// 左
	if col > 0 && grid[row][col-1] == 1 {
		res += dp(row, col-1, grid)
	}

	// 上
	if row > 0 && grid[row-1][col] == 1 {
		res += dp(row-1, col, grid)
	}

	// 右
	if col < len(grid[0])-1 && grid[row][col+1] == 1 {
		res += dp(row, col+1, grid)
	}

	// 下
	if row < len(grid)-1 && grid[row+1][col] == 1 {
		res += dp(row+1, col, grid)
	}

	return res
}
