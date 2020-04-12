// 887. 鸡蛋掉落
// https://leetcode-cn.com/problems/super-egg-drop/

package question887

import (
	"math"
	"testing"
)

func Test_trap(t *testing.T) {
	t.Log(superEggDrop(1, 2) == 2)
}

var memo = map[int]int{}

func superEggDrop(K int, N int) int {
	return dp(K, N)
}

func dp(K, N int) int {
	if _, ok := memo[N*100+K]; !ok {
		var ans int
		if N == 0 {
			ans = 0
		} else if K == 1 {
			ans = N
		} else {
			lo := 1
			hi := N
			for lo+1 < hi {
				x := (lo + hi) / 2
				t1 := dp(K-1, x-1)
				t2 := dp(K, N-x)

				if t1 < t2 {
					lo = x
				} else if t1 > t2 {
					hi = x
				} else {
					lo = x
					hi = x
				}
			}

			ans = 1 + _mix(_max(dp(K-1, lo-1), dp(K, N-lo)),
				_max(dp(K-1, hi-1), dp(K, N-hi)))
		}

		memo[N*100+K] = ans
	}

	return memo[N*100+K]
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}

func _mix(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}
