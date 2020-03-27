// 914. 卡牌分组
// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/

package question914

import (
	"testing"
)

func Test_hasGroupsSizeX(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 4, 3, 2, 1}
	t.Log(hasGroupsSizeX(a1) == true)

	a2 := []int{1, 1, 1, 2, 2, 2, 3, 3}
	t.Log(hasGroupsSizeX(a2) == false)
}

//1 <= deck.length <= 10000
//0 <= deck[i] < 10000
func hasGroupsSizeX(deck []int) bool {
	var count [10000]int
	for _, v := range deck {
		count[v]++
	}

	g := -1
	for i := 0; i < 10000; i++ {
		if count[i] != 0 {
			if g == -1 {
				g = count[i]
			} else {
				g = _gcd(g, count[i])
			}
		}
	}

	return g >= 2
}

func _gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return _gcd(y, tmp)
	} else {
		return y
	}
}
