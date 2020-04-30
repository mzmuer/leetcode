// 202. 快乐数
// https://leetcode-cn.com/problems/happy-number/

package question202

import (
	"testing"
)

func Test_waysToChange(t *testing.T) {
	t.Log(isHappy(19))

}

func isHappy(n int) bool {
	var filter = make(map[int]struct{})

	for {
		if n == 1 {
			return true
		}

		if _, ok := filter[n]; ok {
			return false
		}

		filter[n] = struct{}{}
		tmp := 0
		for n != 0 {
			mod := n % 10
			tmp += mod * mod
			n /= 10
		}
		n = tmp
	}
}
