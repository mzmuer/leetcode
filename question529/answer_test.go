// 529. 扫雷游戏
// https://leetcode-cn.com/problems/minesweeper/

package question529

import (
	"bytes"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	b1 := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	t.Log(string(bytes.Join(updateBoard(b1, []int{3, 0}), []byte{'\n'})))

	b2 := [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}
	t.Log(string(bytes.Join(updateBoard(b2, []int{1, 2}), []byte{'\n'})))
}

func updateBoard(board [][]byte, click []int) [][]byte {
	// 如果点到雷 直接结束
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
	} else {
		bfs(board, click)
	}

	return board
}

func bfs(board [][]byte, click []int) {
	vis := make([][]bool, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, len(board[0]))
	}

	queue := [][]int{click}
	m, n := len(board), len(board[0])

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		// 获取8个相邻方向
		var mCount int
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				x := top[0] + i
				y := top[1] + j
				if x < 0 || x >= m || y < 0 || y >= n || (i == 0 && j == 0) {
					continue
				}

				if board[x][y] == 'M' {
					mCount++
				}
			}
		}

		if mCount > 0 {
			board[top[0]][top[1]] = byte('0' + mCount)
		} else {
			board[top[0]][top[1]] = 'B'
			// 如果当前位置是空的 那就继续获取周围的8个方向,获取8个相邻方向
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					x := top[0] + i
					y := top[1] + j
					if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'E' || (i == 0 && j == 0) || vis[x][y] {
						continue
					}

					queue = append(queue, []int{x, y})
					vis[x][y] = true
				}
			}
		}
	}
}
