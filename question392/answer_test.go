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

// func isSubsequence(s string, t string) bool {
// 	var (
// 		l   = len(s)
// 		idx = 0
// 	)
//
// 	for i := 0; i < len(t) && idx < l; i++ {
// 		if t[i] == s[idx] {
// 			idx++
// 		}
// 	}
//
// 	return idx == l
// }

// 动态规划
func isSubsequence(s string, t string) bool {
	tl := len(t)
	a := make([][26]int, tl+1)

	for i := 0; i < 26; i++ {
		a[tl][i] = tl
	}

	// f[i][j]代表，t中第i个元素的下一个字母"$j"所在的位置
	// 如果 t[i-1]!=$j 那么 f[i-1][j] 也就应该等于 f[i][j]
	// 如果 t[i-1]==$j 那么 f[i-1][j] 应该等于 i
	for i := tl - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				a[i][j] = i
			} else {
				a[i][j] = a[i+1][j]
			}
		}
	}

	add := 0
	for i := 0; i < len(s); i++ {
		if a[add][int(s[i]-'a')] == tl { // 到结尾都还没全部找到
			return false
		}
		add = a[add][int(s[i]-'a')] + 1
	}

	return true
}
