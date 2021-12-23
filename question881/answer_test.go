// 881. 救生艇
// https://leetcode-cn.com/problems/boats-to-save-people/

package question881

import (
	"sort"
	"testing"
)

func Test_numRescueBoats(t *testing.T) {
	t.Log(numRescueBoats([]int{1, 2}, 3))
}

// 贪心
// 先选择最轻的人
// 如果我能和最重的一起走，那就一起走
// 如果不能，那么最重的人只能单独走
// 剩下的人再循环
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var ans int
	left, right := 0, len(people)-1

	for left <= right {
		if people[left]+people[right] > limit {
			right--
		} else {
			right--
			left++
		}
		ans++
	}

	return ans
}
