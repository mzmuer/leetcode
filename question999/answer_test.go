// 999. 车的可用捕获量
// https://leetcode-cn.com/problems/available-captures-for-rook/

package question999

import (
	"testing"
)

func Test_numRookCaptures(t *testing.T) {
	a1 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}
	t.Log(numRookCaptures(a1) == 3)

	a2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', 'p', '.', 'p', '.'},
		{'.', 'p', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'B', '.', 'R', '.', 'B', 'p'},
		{'.', 'p', '.', '.', '.', 'B', '.', '.'},
		{'.', '.', 'p', '.', '.', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}
	t.Log(numRookCaptures(a2) == 1)
}

//board.length == board[i].length == 8
//board[i][j] 可以是 'R'，'.'，'B' 或 'p'
//只有一个格子上存在 board[i][j] == 'R'
func numRookCaptures(board [][]byte) int {
	var ans int
	dt := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if board[row][col] == 'R' {
				for _, d := range dt {
					ans += _impel(board, row, col, d)
				}
				break
			}
		}
	}

	return ans
}

func _impel(board [][]byte, x, y int, d [2]int) int {
	for {
		x += d[0]
		y += d[1]
		if x < 0 || x >= 8 || y < 0 || y >= 8 {
			break
		}

		if board[x][y] == 'B' {
			return 0
		} else if board[x][y] == 'p' {
			return 1
		}
	}

	return 0
}
