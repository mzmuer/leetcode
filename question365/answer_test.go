// 365. 水壶问题
// https://leetcode-cn.com/problems/water-and-jug-problem/

package question365

import (
	"testing"
)

func Test_canMeasureWater(t *testing.T) {
	x1 := 3
	y1 := 5
	z1 := 4
	t.Log(canMeasureWater(x1, y1, z1) == true)
}

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}

	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}

	return z%_gcd(x, y) == 0
}

func _gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return _gcd(y, tmp)
	} else {
		return y
	}
}
