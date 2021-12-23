## 题目：[扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)

让我们一起来玩扫雷游戏！

给定一个代表游戏板的二维字符矩阵。 **'M'** 代表一个**未挖出的**地雷，**'E'** 代表一个**未挖出的**空方块，**'B'** 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的**已挖出的**空白方块，**数字**（'1' 到 '8'）表示有多少地雷与这块**已挖出的**方块相邻，**'X'** 则表示一个**已挖出的**地雷。

现在给出在所有**未挖出的**方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 **'X'**。
如果一个**没有相邻地雷**的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的**未挖出**方块都应该被递归地揭露。
如果一个**至少与一个地雷相邻**的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
如果在此次点击中，若无更多方块可被揭露，则返回面板。

示例 1：
>**输入:**   
>  
>[['E', 'E', 'E', 'E', 'E'],  
> ['E', 'E', 'M', 'E', 'E'],  
> ['E', 'E', 'E', 'E', 'E'],  
> ['E', 'E', 'E', 'E', 'E']]  
>  
>Click : [3,0]  
>  
>**输出:**   
>  
>[['B', '1', 'E', '1', 'B'],  
> ['B', '1', 'M', '1', 'B'],  
> ['B', '1', '1', '1', 'B'],  
> ['B', 'B', 'B', 'B', 'B']]
>  
> **解释：**  
![示例1](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/minesweeper_example_1.png)

示例 2：
>**输入:**   
>  
>[['B', '1', 'E', '1', 'B'],  
> ['B', '1', 'M', '1', 'B'],  
> ['B', '1', '1', '1', 'B'],  
> ['B', 'B', 'B', 'B', 'B']]  
>  
>Click : [1,2]  
>  
>**输出:**   
>  
>[['B', '1', 'E', '1', 'B'],  
> ['B', '1', 'X', '1', 'B'],  
> ['B', '1', '1', '1', 'B'],  
> ['B', 'B', 'B', 'B', 'B']]
>  
> **解释：**  
![示例2](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/minesweeper_example_2.png)

## 思路
1. 使用广度优先搜索。
2. 先将点击位置加入队列，判断该位置的相邻方向。记录下相邻地雷的数量。
3. 如果周围没有地雷，就将周围**未挖开的**且**未加入队列的**的位置加入队列。直到队列为空

## [实现](https://github.com/mzmuer/leetcode/blob/master/question529/answer_test.go)
```go
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
```