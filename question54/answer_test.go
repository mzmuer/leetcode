// 54. 螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix/

package question54

import (
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	// matrix1 := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// t.Log(spiralOrder(matrix1))

	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	t.Log(spiralOrder(matrix2))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		ans = make([]int, 0)
		row = len(matrix)
		col = len(matrix[0])

		dt                = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		visited           = make([][]bool, row)
		total, turn, x, y = row * col, 0, 0, 0
	)

	// 初始化矩阵
	for i := range visited {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < total; i++ {
		visited[x][y] = true
		ans = append(ans, matrix[x][y])

		d := dt[turn]
		if x+d[0] >= row || x+d[0] < 0 || y+d[1] >= col || y+d[1] < 0 || visited[x+d[0]][y+d[1]] { // 转向
			turn = (turn + 1) % 4
		}

		d = dt[turn]
		x += d[0]
		y += d[1]
	}

	return ans
}
