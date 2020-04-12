// 面试题 01.07. 旋转矩阵
// https://leetcode-cn.com/problems/rotate-matrix-lcci/

package question01_07

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix1)
	t.Log(matrix1)
}

func rotate(matrix [][]int) {
	copyMatrix := make([][]int, len(matrix))
	for i := range matrix {
		copyMatrix[i] = make([]int, len(matrix[i]))
		copy(copyMatrix[i], matrix[i])
	}

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = copyMatrix[i][j]
		}
	}
}
