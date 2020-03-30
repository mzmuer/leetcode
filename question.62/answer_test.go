// 面试题62. 圆圈中最后剩下的数字
// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

package question_62

import (
	"testing"
)

func Test_lastRemaining(t *testing.T) {
	t.Log(lastRemaining(5, 3) == 3)

	t.Log(lastRemaining(10, 17) == 2)
}

//1 <= deck.length <= 10000
//0 <= deck[i] < 10000
//func lastRemaining(n int, m int) int {
//	var f int
//	for i := 2; i != n+1; i++ {
//		f = (m + f) % i
//	}
//	return f
//}
func lastRemaining(n int, m int) int {
	return f(n, m)
}

func f(n, m int) int {
	if n == 1 {
		return 0
	}

	x := f(n-1, m)
	return (m + x) % n
}
