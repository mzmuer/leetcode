// 64. 最小路径和
// https://leetcode-cn.com/problems/minimum-path-sum/

package question64

import (
	"math"
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	a1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	t.Log(minPathSum(a1))

	a2 := [][]int{
		{0},
	}
	t.Log(minPathSum(a2))

	a3 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	t.Log(minPathSum(a3))
}

// n,m的最小路径 = min{(n-1,m),(n,m-1)}
// TODO: 应该可以用滚动数组来优化
// func minPathSum(grid [][]int) int {
// 	n, m := len(grid), len(grid[0])
// 	f := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		f[i] = make([]int, m)
// 	}
//
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if i == 0 && j == 0 {
// 				f[0][0] = grid[0][0]
// 				continue
// 			} else if i == 0 {
// 				f[i][j] = f[i][j-1] + grid[i][j]
// 			} else if j == 0 {
// 				f[i][j] = f[i-1][j] + grid[i][j]
// 			} else {
// 				f[i][j] = _min(f[i][j-1], f[i-1][j]) + grid[i][j]
// 			}
// 		}
// 	}
//
// 	return f[n-1][m-1]
// }

// 原地修改
func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = _min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}

	return grid[n-1][m-1]
}

func _min(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}
