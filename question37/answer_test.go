// 37. 解数独
// https://leetcode-cn.com/problems/sudoku-solver/

package question37

import (
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	matrix := [][]byte{
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
	solveSudoku(matrix)
	t.Log(matrix)
}

func solveSudoku(board [][]byte) {
	var (
		row, column [9][9]bool
		block       [3][3][9]bool
		spaces      [][2]int
	)

	for i, rows := range board {
		for j, v := range rows {
			if v != '.' {
				digit := v - '1'
				row[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			} else {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}

	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}

		i, j := spaces[pos][0], spaces[pos][1]
		for digit := 0; digit < 9; digit++ {
			if !row[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				row[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = byte(digit + '1')
				if dfs(pos + 1) {
					return true
				}
				row[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}

		return false
	}

	dfs(0)
}
