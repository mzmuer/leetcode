// 1162. 地图分析
// https://leetcode-cn.com/problems/as-far-from-land-as-possible/

package question1162

import (
	"testing"
)

func Test_maxDistance(t *testing.T) {
	a1 := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1}}
	t.Log(maxDistance(a1) == 2)

	a2 := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	t.Log(maxDistance(a2) == -1)
}

//1 <= grid.length == grid[0].length <= 100
//grid[i][j] 不是 0 就是 1
func maxDistance(grid [][]int) int {
	var ans = -1

	for x, rows := range grid {
		for y, v := range rows {
			if v != 0 {
				continue
			}

			if d := findNearestLand(x, y, grid); ans < d {
				ans = d
			}
		}
	}

	return ans
}

type coordinate struct {
	x, y, step int
}

func findNearestLand(x, y int, grid [][]int) int {
	var (
		vis   = [100][100]int{}
		dt    = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		queue = make([]coordinate, 0)
	)

	queue = append(queue, coordinate{x, y, 0})
	for len(queue) != 0 {
		f := queue[0]
		queue = queue[1:]
		for _, d := range dt {
			nx := f.x + d[0]
			ny := f.y + d[1]
			if !(nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0])) {
				continue
			}

			if vis[nx][ny] == 0 {
				queue = append(queue, coordinate{nx, ny, f.step + 1})
				vis[nx][ny] = 1
				if grid[nx][ny] == 1 { // 找到最近的陆地
					return f.step + 1
				}
			}
		} // for
	} // for

	return -1
}
