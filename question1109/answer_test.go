// 1109. 航班预订统计
// https://leetcode-cn.com/problems/corporate-flight-bookings/

package question1109

import "testing"

func Test_corpFlightBookings(t *testing.T) {
	t.Log(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, v := range bookings {
		ans[v[0]-1] += v[2]
		if v[1] < n {
			ans[v[1]] -= v[2]
		}
	}

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}

	return ans
}
