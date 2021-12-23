// 27. 移除元素
// https://leetcode-cn.com/problems/remove-element/

package question42

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	a1 := []int{3, 2, 2, 3}
	t.Log(removeElement(a1, 3) == 2, a1)

	a2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	t.Log(removeElement(a2, 2) == 5, a2)
}

func removeElement(nums []int, val int) int {

	var (
		ans int
		end = len(nums) - 1
	)
	for i := 0; i <= end; {
		if nums[i] == val {
			nums[i] = nums[end]
			end--
			continue
		}
		ans++
		i++
	}

	return ans
}
