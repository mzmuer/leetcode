## 题目：[岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**示例1:**
>输入:  
11110  
11010  
11000  
00000  
输出: 1

**示例2:**
>输入:  
11000  
11000  
00100  
00011  
输出: 3  
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。  

## 思路
1. 广度优先搜索
2. 如果位置是1，就加入队列搜索，搜索过的部分就标记为0，完成一次搜索岛屿数量+1

## [实现](https://github.com/mzmuer/leetcode/blob/master/question200/answer_test.go)
```go
type pair struct {
	r, c int
}

func numIslands(grid [][]byte) int {
	var (
		dt  = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		ans int
	)

	for row, rows := range grid {
		for col, v := range rows {
			if v == '1' {
				queue := []pair{{row, col}}
				for len(queue) > 0 {
					f := queue[0]
					queue = queue[1:]

					for _, d := range dt {
						r := f.r + d[0]
						c := f.c + d[1]

						if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
							continue
						}

						if grid[r][c] == '1' {
							grid[r][c] = '0'
							queue = append(queue, pair{r, c})
						}
					}
				}

				// ---
				ans++
			}
		}
	}

	return ans
}
```