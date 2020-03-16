// 面试题 01.06. 字符串压缩
// https://leetcode-cn.com/problems/compress-string-lcci/

package question01_06

import (
	"strconv"
	"testing"
)

func Test_compressString(t *testing.T) {
	str1 := "aabcccccaaa"
	t.Log(compressString(str1) == "a2b1c5a3")

	str2 := "bb"
	t.Log(compressString(str2) == "bb")
}

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}

	var (
		ans   []byte
		count = 1
		cur   = S[0]
	)

	for i := 1; i < len(S); i++ {
		if S[i] == cur {
			count++
		} else {
			ans = append(ans, cur)
			ans = append(ans, strconv.Itoa(count)...)
			cur = S[i]
			count = 1
		}
	}

	ans = append(ans, cur)
	ans = append(ans, strconv.Itoa(count)...)

	if len(ans) >= len(S) {
		return S
	}

	return string(ans)
}
