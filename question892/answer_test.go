// 892. 三维形体的表面积
// https://leetcode-cn.com/problems/surface-area-of-3d-shapes/

package question892

import (
	"testing"
)

func Test_surfaceArea(t *testing.T) {
	a1 := [][]int{{2}}
	t.Log(surfaceArea(a1) == 10)

	a2 := [][]int{{1, 2}, {3, 4}}
	t.Log(surfaceArea(a2) == 34)

	a3 := [][]int{{1, 0}, {0, 2}}
	t.Log(surfaceArea(a3) == 16)

	a4 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	t.Log(surfaceArea(a4) == 32)

	a5 := [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}
	t.Log(surfaceArea(a5) == 46)
}

func surfaceArea(grid [][]int) int {
	var ans int

	for row, rows := range grid {
		for col, v := range rows {
			if v == 0 {
				continue
			}

			// 上下两个面
			ans += 2

			// 上
			if row == 0 {
				ans += v
			} else if grid[row-1][col] < v {
				ans += v - grid[row-1][col]
			}

			// 左
			if col == len(rows)-1 {
				ans += v
			} else if grid[row][col+1] < v {
				ans += v - grid[row][col+1]
			}

			// 下
			if row == len(grid)-1 {
				ans += v
			} else if grid[row+1][col] < v {
				ans += v - grid[row+1][col]
			}

			// 右
			if col == 0 {
				ans += v
			} else if grid[row][col-1] < v {
				ans += v - grid[row][col-1]
			}
		}
	}

	return ans
}
