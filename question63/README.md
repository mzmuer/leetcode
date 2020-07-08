## 题目：[不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/robot_maze.png)
网格中的障碍物和空位置分别用 `1` 和 `0` 来表示。

说明：`m` 和 `n` 的值均不超过 100。

**示例1:**
>输入:  
>[  
>  [0,0,0],  
>  [0,1,0],  
>  [0,0,0]  
>]  
>输出: 2  
>解释:  
>3x3 网格的正中间有一个障碍物。  
>从左上角到右下角一共有 2 条不同的路径：  
>1. 向右 -> 向右 -> 向下 -> 向下  
>2. 向下 -> 向下 -> 向右 -> 向右  
     
## 思路
1. 假设终点为`(i,j)`,达到终点的路径为`f(i,j)`。
2. 因为机器人只能向下或者向右运动，那么`f(i,j) = f(i-1,j)+f(i,j-1)`。那么可以使用动态规格来解决这个问题。
3. 如果当前位置有障碍，那么无法到达当前位置`f(i,j)=0`
4. 可以使用滚动数组思想优化的数组为,使得空间复杂度从`O(mn)`到`O(m)`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question63/answer_test.go)
```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
	}

	if obstacleGrid[0][0] == 0 {
		f[0][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] != 0 {
				f[i][j] = 0
				continue
			}

			if i-1 >= 0 && obstacleGrid[i-1][j] == 0 {
				f[i][j] += f[i-1][j]
			}

			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[i][j] += f[i][j-1]
			}
		}
	}

	return f[n-1][m-1]
}
```