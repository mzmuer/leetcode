// 447. 回旋镖的数量
// https://leetcode-cn.com/problems/number-of-boomerangs/

package question447

import (
	"testing"
)

func Test_numberOfBoomerangs(t *testing.T) {
	t.Log(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}

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
