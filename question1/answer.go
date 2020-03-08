// 1. 两数之和
// https://leetcode-cn.com/problems/two-sum/

package question1

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		diff := target - v
		if index, ok := m[diff]; ok {
			return []int{index, i}
		}

		m[v] = i
	}

	return []int{}
}
