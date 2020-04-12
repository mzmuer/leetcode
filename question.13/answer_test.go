// 面试题13. 机器人的运动范围
// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

package question_13

import (
	"testing"
)

func Test_movingCount(t *testing.T) {
	t.Log(movingCount(2, 3, 1) == 3)

	t.Log(movingCount(3, 1, 0) == 1)

	t.Log(movingCount(1, 3, 11) == 3)

	t.Log(movingCount(11, 8, 16) == 88)

	t.Log(movingCount(16, 8, 4) == 15)
}

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
