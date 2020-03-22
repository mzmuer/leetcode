// 945. 使数组唯一的最小增量
// https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/

package question945

import (
	"testing"
)

func Test_minIncrementForUnique(t *testing.T) {
	a1 := []int{1, 2, 2}
	t.Log(minIncrementForUnique(a1) == 1)

	a2 := []int{3, 2, 1, 2, 1, 7}
	t.Log(minIncrementForUnique(a2) == 6)
}

//0 <= A.length <= 40000
//0 <= A[i] < 40000
func minIncrementForUnique(A []int) int {
	arr := [80000]int{}
	for _, v := range A {
		arr[v]++
	}

	var (
		ans   int
		token int
	)
	for i := 0; i < 80000; i++ {
		if arr[i] > 1 {
			token += arr[i] - 1
			ans -= i * (arr[i] - 1)
		} else if token > 0 && arr[i] == 0 {
			token--
			ans += i
		}
	}

	return ans
}
