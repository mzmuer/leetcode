// 1025. 除数博弈
// https://leetcode-cn.com/problems/divisor-game/

package question1025

import (
	"testing"
)

func Test_divisorGame(t *testing.T) {
	t.Log(divisorGame(2))
}

func divisorGame(N int) bool {
	return N%2 == 0
}
