// 136. 只出现一次的数字
// https://leetcode-cn.com/problems/single-number/submissions/

package question136

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	a1 := []int{4, 1, 2, 1, 2}
	t.Log(singleNumber(a1))
}

func singleNumber(nums []int) int {
	var ret int
	for i := range nums {
		ret ^= nums[i]
	}

	return ret
}
