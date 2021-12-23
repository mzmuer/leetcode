## 题目：[岛屿的最大面积](https://leetcode-cn.com/problems/max-area-of-island/)

给定一个包含了一些 0 和 1的非空二维数组 `grid` , 一个 岛屿 是由四个方向 (水平或垂直) 的 `1` (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

**示例1:**
>[[0,0,1,0,0,0,0,1,0,0,0,0,0],  
 [0,0,0,0,0,0,0,1,1,1,0,0,0],  
 [0,1,1,0,1,0,0,0,0,0,0,0,0],  
 [0,1,0,0,1,1,0,0,1,0,1,0,0],  
 [0,1,0,0,1,1,0,0,1,1,1,0,0],  
 [0,0,0,0,0,0,0,0,0,0,1,0,0],  
 [0,0,0,0,0,0,0,1,1,1,0,0,0],  
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]  

对于上面这个给定矩阵应返回 `6`。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

**示例2:**
>[[0,0,0,0,0,0,0,0]]

对于上面这个给定的矩阵, 返回 `0`。

注意: 给定的矩阵`grid` 的长度和宽度都不超过 50。
     
## 思路
1. 遍历矩阵，遇到陆地就向上下左右四个方向递归
2. 避免重复的递归，经过的点都设置为0

## 实现
```go
func maxAreaOfIsland(grid [][]int) int {
	var res int

	for row, rows := range grid {
		for col, v := range rows {
			if v != 1 {
				continue
			}

			tmp := dp(row, col, grid)
			if res < tmp {
				res = tmp
			}
		}
	}

	return res
}

func dp(row, col int, grid [][]int) int {
	grid[row][col] = 0
	var res = 1

	// 左
	if col > 0 && grid[row][col-1] == 1 {
		res += dp(row, col-1, grid)
	}

	// 上
	if row > 0 && grid[row-1][col] == 1 {
		res += dp(row-1, col, grid)
	}

	// 右
	if col < len(grid[0])-1 && grid[row][col+1] == 1 {
		res += dp(row, col+1, grid)
	}

	// 下
	if row < len(grid)-1 && grid[row+1][col] == 1 {
		res += dp(row+1, col, grid)
	}

	return res
}
```