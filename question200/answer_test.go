// 200. 岛屿数量
// https://leetcode-cn.com/problems/number-of-islands/

package question200

import "testing"

func Test_numIslands(t *testing.T) {
	g := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	t.Log(numIslands(g))
}

type pair struct {
	r, c int
}

func numIslands(grid [][]byte) int {
	var (
		dt  = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		ans int
	)

	for row, rows := range grid {
		for col, v := range rows {
			if v == '1' {
				queue := []pair{{row, col}}
				for len(queue) > 0 {
					f := queue[0]
					queue = queue[1:]

					for _, d := range dt {
						r := f.r + d[0]
						c := f.c + d[1]

						if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
							continue
						}

						if grid[r][c] == '1' {
							grid[r][c] = '0'
							queue = append(queue, pair{r, c})
						}
					}
				}

				// ---
				ans++
			}
		}
	}

	return ans
}
