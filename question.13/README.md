## 题目：[机器人的运动范围](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

地上有一个m行n列的方格，从坐标 `[0,0]` 到坐标 `[m-1,n-1]` 。一个机器人从坐标 `[0, 0]` 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 `[35, 37]` ，因为`3+5+3+7=18`。但它不能进入方格 `[35, 38]`，因为`3+5+3+8=19`。请问该机器人能够到达多少个格子？

**示例1:**
>输入：m = 2, n = 3, k = 1  
>输出：3  

**示例2:**
>输入：m = 3, n = 1, k = 0  
>输出：1  

提示：
* `1 <= n,m <= 100`
* `0 <= k <= 20`

## 思路
1. 使用广度优先搜索，从0,0开始往下和右搜索。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question8/answer_test.go)
```go
type pair struct {
	x, y int
}

// 1 <= n,m <= 100
// 0 <= k <= 20
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	var ans = 1
	dt := [2][2]int{{0, 1}, {1, 0}} // 向右和向下的方向数组
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}
	queue := make([]pair, 0)
	queue = append(queue, pair{0, 0})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range dt {
			tx := d[0] + p.x
			ty := d[1] + p.y

			if tx < 0 || tx >= m || ty < 0 || ty >= n || vis[tx][ty] == 1 || _get(tx)+_get(ty) > k {
				continue
			}

			queue = append(queue, pair{tx, ty})
			vis[tx][ty] = 1
			ans++
		}
	}

	return ans
}

func _get(x int) int {
	res := 0
	for ; x > 0; x /= 10 {
		res += x % 10
	}
	return res
}
```