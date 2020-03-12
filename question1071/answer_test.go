// 1071. 字符串的最大公因子
// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/

package question1013

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	str1 := "ABCABC"
	str2 := "ABC"
	t.Log(gcdOfStrings(str1, str2) == "ABC")

	str1 = "ABABAB"
	str2 = "ABAB"
	t.Log(gcdOfStrings(str1, str2) == "AB")

	str1 = "LEET"
	str2 = "CODE"
	t.Log(gcdOfStrings(str1, str2) == "")
}

func gcdOfStrings(str1 string, str2 string) string {
	T := str1[:_gcd(len(str1), len(str2))]
	if check(T, str1) && check(T, str2) {
		return T
	}

	return ""
}

func check(t, s string) bool {
	lenx := len(s) / len(t)
	var ans string
	for i := 1; i <= lenx; i++ {
		ans = ans + t
	}

	return ans == s
}

func _gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return _gcd(y, tmp)
	} else {
		return y
	}
}
