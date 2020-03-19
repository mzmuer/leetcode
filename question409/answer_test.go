// 409. 最长回文串
// https://leetcode-cn.com/problems/longest-palindrome/

package question409

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	str1 := "abccccdd"
	t.Log(longestPalindrome(str1) == 7)

	//str1 = "ABABAB"
	//str2 = "ABAB"
	//t.Log(gcdOfStrings(str1, str2) == "AB")
	//
	//str1 = "LEET"
	//str2 = "CODE"
	//t.Log(gcdOfStrings(str1, str2) == "")
}

func longestPalindrome(s string) int {
	chars := [52]int{}
	for _, c := range s {
		if c >= 'a' {
			chars[25+c-'a']++
		} else {
			chars[c-'A']++
		}
	}

	var ans int

	for _, v := range chars {
		if v%2 == 0 {
			ans += v
		} else {
			if ans%2 == 0 {
				ans += v
			} else {
				ans += v / 2 * 2
			}
		}
	}

	return ans
}
