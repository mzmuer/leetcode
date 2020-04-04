// 289. 生命游戏
// https://leetcode-cn.com/problems/game-of-life/

package question289

import (
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	board1 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(board1)
	t.Log(board1)
}

func gameOfLife(board [][]int) {
	copyBoard := make([][]int, len(board))
	for i := range board {
		copyBoard[i] = make([]int, len(board[i]))
		copy(copyBoard[i], board[i])
	}
	dt := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for row, rows := range board {
		for col := range rows {
			var live int
			for _, d := range dt {
				r := row + d[0]
				c := col + d[1]
				if r < 0 || r >= len(copyBoard) || c < 0 || c >= len(copyBoard[0]) {
					continue
				}

				if copyBoard[r][c] == 1 {
					live++
				}
			}

			// 规则 1 或规则 3
			if (copyBoard[row][col] == 1) && (live < 2 || live > 3) {
				board[row][col] = 0
			}
			// 规则 4
			if copyBoard[row][col] == 0 && live == 3 {
				board[row][col] = 1
			}
		}
	}
}
