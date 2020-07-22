## 题目：[螺旋矩阵](https://leetcode-cn.com/problems/spiral-matrix/)

给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

**示例1:**
>输入:  
>[  
> [ 1, 2, 3 ],  
> [ 4, 5, 6 ],  
> [ 7, 8, 9 ]  
>]   ，xaaZA
>输出: [1,2,3,6,9,8,7,4,5]

**示例2:**
>输入:  
>[  
>  [1, 2, 3, 4],  
>  [5, 6, 7, 8],  
>  [9,10,11,12] 
> 
>]  
>输出: [1,2,3,4,8,12,11,10,9,5,6,7]  
     
## 思路
1. 使用一个辅助矩阵，存储遍历过的元素。
2. 遍历矩阵，达到边界就转向，直到遍历完所有的元素

## [实现](https://github.com/mzmuer/leetcode/blob/master/question54/answer_test.go)
```go
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		ans = make([]int, 0)
		row = len(matrix)
		col = len(matrix[0])

		dt                = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		visited           = make([][]bool, row)
		total, turn, x, y = row * col, 0, 0, 0
	)

	// 初始化矩阵
	for i := range visited {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < total; i++ {
		visited[x][y] = true
		ans = append(ans, matrix[x][y])

		d := dt[turn]
		if x+d[0] >= row || x+d[0] < 0 || y+d[1] >= col || y+d[1] < 0 || visited[x+d[0]][y+d[1]] { // 转向
			turn = (turn + 1) % 4
		}

		d = dt[turn]
		x += d[0]
		y += d[1]
	}

	return ans
}
```