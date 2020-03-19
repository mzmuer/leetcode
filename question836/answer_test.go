// 836. 矩形重叠
// https://leetcode-cn.com/problems/rectangle-overlap/

package question836

import (
	"testing"
)

func Test_isRectangleOverlap(t *testing.T) {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	t.Log(isRectangleOverlap(rec1, rec2) == true)

	//str1 = "ABABAB"
	//str2 = "ABAB"
	//t.Log(gcdOfStrings(str1, str2) == "AB")
	//
	//str1 = "LEET"
	//str2 = "CODE"
	//t.Log(gcdOfStrings(str1, str2) == "")
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || // left
		rec1[3] <= rec2[1] || // bottom
		rec1[0] >= rec2[2] || // right
		rec1[1] >= rec2[3]) // top
}
