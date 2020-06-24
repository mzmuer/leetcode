// 169. 多数元素
// https://leetcode-cn.com/problems/majority-element/

package question169

import "testing"

func Test_majorityElement(t *testing.T) {
	a1 := []int{3, 2, 3}
	t.Log(majorityElement(a1) == 3)

	a2 := []int{2, 2, 1, 1, 1, 2, 2}
	t.Log(majorityElement(a2) == 2)

	a3 := []int{2}
	t.Log(majorityElement(a3) == 2)
}

func majorityElement(nums []int) int {
	var (
		count         = map[int]int{}
		majority, cnt int
		l             = len(nums)
	)

	for _, num := range nums {
		if count[num]++; count[num] > cnt {
			majority = num
			cnt = count[num]

			// 如果已经是多数了 提前跳出
			if count[num] > l/2 {
				break
			}
		}
	}

	return majority
}
