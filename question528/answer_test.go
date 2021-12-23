// 528. 按权重随机选择
// https://leetcode-cn.com/problems/random-pick-with-weight/

package question528

import (
	"math/rand"
	"testing"
)

func Test_solution(t *testing.T) {
	obj := Constructor([]int{3, 14, 1, 7})
	for i := 0; i < 10; i++ {
		t.Log(obj.PickIndex())
	}
}

type Solution struct {
	pre     []int
	maximum int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}

	return Solution{pre: w, maximum: w[len(w)-1]}
}

// [3, 4, 6, 10]
func (this *Solution) PickIndex() int {
	n := rand.Intn(this.maximum) + 1

	// 二分
	i, j := 0, len(this.pre)
	for i < j {
		mid := (j + i) / 2
		if this.pre[mid] >= n {
			j = mid
		} else {
			i = mid + 1
		}
	}

	return i
}
