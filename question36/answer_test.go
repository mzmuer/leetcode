// 36. 有效的数独
// https://leetcode-cn.com/problems/valid-sudoku/

package question36

import (
	"testing"
)

func Test_isValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	t.Log(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	cols := [9][9]int{}
	areas := [3][3][9]int{}

	for i, row := range board {
		for j, bv := range row {
			if bv == '.' {
				continue
			}

			v := bv - '1'
			if rows[i][v] == 0 && cols[j][v] == 0 && areas[i/3][j/3][v] == 0 {
				rows[i][v]++
				cols[j][v]++
				areas[i/3][j/3][v]++
			} else {
				return false
			}
		}
	}

	return true
}
