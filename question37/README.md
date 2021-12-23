## 题目：[解数独](https://leetcode-cn.com/problems/sudoku-solver/)

编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需**遵循如下规则**：
1. 数字 `1-9` 在每一行只能出现一次。
2. 数字 `1-9` 在每一列只能出现一次。
3. 数字 `1-9` 在每一个以粗实线分隔的 `3x3` 宫内只能出现一次。

空白格用 `'.'` 表示。

![](http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

一个数独。

![](http://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png)

答案被标成红色。

**Note**:
* 给定的数独序列只包含数字 `1-9` 和字符 `'.'` 
* 你可以假设给定的数独只有唯一解。
* 给定数独永远是 `9x9` 形式的。
     
## 思路
1. 使用`row[9][9]`来保存行的状态，使用`column[9][9]`来保存列的状态，使用`block[3][3][9]`来保存每一个块的状态。
2. 首先遍历一遍`board`初始化状态，同时使用`spaces[][2]`记录需要填充的坐标。
3. 递归填充空闲的位置，如果能填入到该位置，就继续向下尝试。直到全部填充完成。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question37/answer_test.go)
```go
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
```