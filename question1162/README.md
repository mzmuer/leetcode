## 题目：[地图分析](https://leetcode-cn.com/problems/as-far-from-land-as-possible/)

你现在手里有一份大小为 N x N 的『地图』（网格） `grid`，上面的每个『区域』（单元格）都用 `0` 和 `1` 标记好了。其中 `0` 代表海洋，`1` 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。

我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：`(x0, y0)` 和 `(x1, y1)` 这两个区域之间的距离是 `|x0 - x1| + |y0 - y1|`。

如果我们的地图上只有陆地或者海洋，请返回 `-1`。

**示例1:**

![示例1](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/08/17/1336_ex1.jpeg "示例1")
>输入：[[1,0,1],[0,0,0],[1,0,1]]  
>输出：2  
>解释：   
>海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。

**示例2:**

![示例2](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/08/17/1336_ex2.jpeg "示例2")
>输入：[[1,0,0],[0,0,0],[0,0,0]]  
>输出：4  
>解释：   
>海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。

提示：
1. `1 <= grid.length == grid[0].length <= 100`
2. `grid[i][j] 不是 0 就是 1`

## 思路
1. 使用宽度优先搜索，当遇到海洋是从四个方线搜索，找到最近的陆地后返回

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1162/answer_test.go)
```go
func maxDistance(grid [][]int) int {
	var ans = -1

	for x, rows := range grid {
		for y, v := range rows {
			if v != 0 {
				continue
			}

			if d := findNearestLand(x, y, grid); ans < d {
				ans = d
			}
		}
	}

	return ans
}

type coordinate struct {
	x, y, step int
}

func findNearestLand(x, y int, grid [][]int) int {
	var (
		vis   = [100][100]int{}
		dt    = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		queue = make([]coordinate, 0)
	)

	queue = append(queue, coordinate{x, y, 0})
	for len(queue) != 0 {
		f := queue[0]
		queue = queue[1:]
		for _, d := range dt {
			nx := f.x + d[0]
			ny := f.y + d[1]
			if !(nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0])) {
				continue
			}

			if vis[nx][ny] == 0 {
				queue = append(queue, coordinate{nx, ny, f.step + 1})
				vis[nx][ny] = 1
				if grid[nx][ny] == 1 { // 找到最近的陆地
					return f.step + 1
				}
			}
		} // for
	} // for

	return -1
}
```