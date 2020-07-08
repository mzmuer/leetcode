// 125. 验证回文串
// https://leetcode-cn.com/problems/valid-palindrome/

package question125

import (
	"strings"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	t.Log(isPalindrome("A man, a plan, a canal: Panama")) // true

	t.Log(isPalindrome("race a car")) // false

	t.Log(isPalindrome("0P")) // false
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	var (
		left, right = 0, len(s) - 1
		ans         = true
	)

	for left <= right {
		if !_isalnum(s[left]) {
			left++
			continue
		}

		if !_isalnum(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			ans = false
			break
		}

		left++
		right--
	}

	return ans
}

func _isalnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
