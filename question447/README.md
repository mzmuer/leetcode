## 题目：[两数相加 II](https://leetcode-cn.com/problems/add-two-numbers-ii/)

给定平面上 `n` 对 **互不相同** 的点 `points` ，其中 `points[i] = [xi, yi]` 。**回旋镖** 是由点 `(i, j, k)` 表示的元组 ，其中 `i` 和 `j` 之间的距离和 `i` 和 `k` 之间的距离相等（**需要考虑元组的顺序**）。

返回平面上所有回旋镖的数量。

**示例1:**
>输入：points = [[0,0],[1,0],[2,0]]  
 输出：2  
 解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]

**示例2:**
>输入：points = [[1,1],[2,2],[3,3]]  
 输出：2

**示例3:**
>输入：points = [[1,1]]  
 输出：0

提示：
* `n == points.length`
* `1 <= n <= 500`
* `points[i].length == 2`
* `-104 <= xi, yi <= 104`
* 所有点都 **互不相同**


## 思路
1. 以点 `i` 为中心，遍历所有的点，获取到有点的距离
2. 用hash以距离为k保存下以 `i` 为中心，相同距离所有点的数量
3. 相同距离点的数量为`m`，那么组合有`m*(m-1)`种，即全排列
4. 遍历hash，累加得到当前 `i` 的组合数量
5. 遍历所有的点做为中心。得到最后结果

## [实现](https://github.com/mzmuer/leetcode/blob/master/question447/answer_test.go)
```go
func numberOfBoomerangs(points [][]int) int {
	var ans int
	for i, x := range points {
		distSet := map[int]int{}
		for j, y := range points {
			if i == j {
				continue
			}

			distSet[_virtualDist(x, y)]++
		}

		for _, v := range distSet {
			ans += v * (v - 1)
		}
	}

	return ans
}

func _virtualDist(x, y []int) int {
	return (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1])
}
```