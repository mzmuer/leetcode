// 392. 判断子序列
// https://leetcode-cn.com/problems/is-subsequence/

package question392

import (
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	t.Log(isSubsequence("abc", "ahbgdc"))

	t.Log(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	idx := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[idx] {
			if idx++; idx == l {
				return true
			}
		}
	}

	return false
}
