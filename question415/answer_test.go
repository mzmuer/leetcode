// 415. 字符串相加
// https://leetcode-cn.com/problems/add-strings/

package question415

import (
	"testing"
)

func Test_integerBreak(t *testing.T) {
	t.Log(addStrings("1", "9"))

	t.Log(addStrings("2", "6572"))
}

func addStrings(num1 string, num2 string) string {
	var ans string
	var d int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x int
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		var y int
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		ans = string((x+y+d)%10+'0') + ans
		d = (x + y + d) / 10
	}

	if d != 0 {
		ans = string('0'+d) + ans
	}

	return ans
}
